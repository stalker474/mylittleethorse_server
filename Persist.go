package main

import (
	"bytes"
	"compress/gzip"
	"encoding/csv"
	"encoding/json"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"
)

var tempDbFile = "mle_db.json"

// ByRaceNumber sorter
type ByRaceNumber []RaceData

func (a ByRaceNumber) Len() int           { return len(a) }
func (a ByRaceNumber) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByRaceNumber) Less(i, j int) bool { return a[i].RaceNumber > a[j].RaceNumber }

// PersistObject blabla
type PersistObject struct {
	racesData map[uint32]RaceData
	mux       sync.Mutex
}

// NewPersistObject Create new database
func NewPersistObject() (p *PersistObject) {
	p = new(PersistObject)
	p.racesData = make(map[uint32]RaceData)
	return p
}

// Save blabla
func (p *PersistObject) save() error {
	p.mux.Lock()
	cache := NewCache(p.racesData, 0, 99999)
	p.mux.Unlock()
	resp, err := json.Marshal(cache)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(tempDbFile, resp, 0644)
}

// Load blabla
func (p *PersistObject) load() error {
	jsonText, err := ioutil.ReadFile(tempDbFile)
	if err != nil {
		return err
	}
	var cache Cache
	err = json.Unmarshal(jsonText, &cache)
	p.mux.Lock()
	p.racesData = cache.toMap()
	p.mux.Unlock()
	return nil
}

func (p *PersistObject) contains(raceNumber uint32) bool {
	_, exists := p.racesData[raceNumber]
	return exists
}

func (p *PersistObject) toJSON(from uint32, to uint32) (s string, err error) {
	p.mux.Lock()
	data, err := json.Marshal(NewCache(p.racesData, from, to))
	p.mux.Unlock()

	s = string(data[:])
	return s, err
}

func (p *PersistObject) toZJSON(from uint32, to uint32) (s string, err error) {
	var buf bytes.Buffer
	zw := gzip.NewWriter(&buf)
	data, err := p.toJSON(from, to)

	_, err = zw.Write([]byte(data))
	if err != nil {
		return s, err
	}

	if err := zw.Close(); err != nil {
		return s, err
	}

	return buf.String(), err
}

func (p *PersistObject) toCharts(from uint64, to uint64) (s string, err error) {
	var buf bytes.Buffer
	zw := gzip.NewWriter(&buf)

	stats := Stats{}
	stats.CoinInfo = make(map[string]CoinInfo)
	stats.CoinInfo["BTC"] = CoinInfo{WinsCount: 0}
	stats.CoinInfo["ETH"] = CoinInfo{WinsCount: 0}
	stats.CoinInfo["LTC"] = CoinInfo{WinsCount: 0}

	playerCount := make(map[string]uint32)
	dayPlayerCount := make(map[string]map[string]uint32)
	days := make(map[string]Day)
	stats.PeriodBegin = from
	stats.PeriodEnd = to
	seedingAddr := "0x1F92771237Bd5eae04e91B4B6F1d1a78D41565a2"

	p.mux.Lock()
	for _, race := range server.data.racesData {
		//count only non refunded closed races
		if race.Date >= from && race.Date <= to {
			if !race.Refunded && (strings.Compare(race.Active, "Closed") == 0) {
				//local to race data
				coinVolume := make(map[string]float32)
				coinBetsCount := make(map[string]uint32)

				tm := time.Unix(int64(race.Date), 0)
				const format = "Jan 02 2006"
				formattedDay := tm.Format(format)
				day, exists := days[formattedDay]
				if !exists {
					day = Day{}
					day.Coins = make(map[string]Coin)
				}
				day.Label = formattedDay
				dayTime, err := time.Parse(format, formattedDay)
				if err != nil {
					return "", err
				}
				day.Date = uint64(dayTime.Unix())
				stats.TotalRaces++
				stats.TotalVolume += race.Volume

				for _, horse := range race.WinnerHorses {
					stats.CoinInfo[horse] = CoinInfo{WinsCount: stats.CoinInfo[horse].WinsCount + 1}
				}
				userVolume := float32(0)
				for _, bet := range race.Bets {
					playerCount[bet.From]++
					_, exists := dayPlayerCount[formattedDay]
					if !exists {
						dayPlayerCount[formattedDay] = make(map[string]uint32)
					}
					dayPlayerCount[formattedDay][bet.From]++
					coinVolume[bet.Horse] += bet.Value
					coinBetsCount[bet.Horse]++
					if strings.Compare(bet.From, seedingAddr) != 0 {
						userVolume += bet.Value
					}
				}

				day.Volume += race.Volume
				day.UserVolume += userVolume

				for coin, value := range coinVolume {
					c, _ := day.Coins[coin]
					c.Volume += value
					c.BetsCount += coinBetsCount[coin]
					day.Coins[coin] = c
				}

				days[formattedDay] = day
			}
		}
	}
	p.mux.Unlock()

	stats.TotalPlayersCount = uint32(len(playerCount))
	for formattedDay, coins := range dayPlayerCount {
		day, _ := days[formattedDay]
		day.PlayersCount = uint32(len(coins))
		days[formattedDay] = day
	}

	for _, day := range days {
		stats.Days = append(stats.Days, day)
	}
	sort.Slice(stats.Days, func(i, j int) bool {
		return stats.Days[i].Date < stats.Days[j].Date
	})

	data, err := json.Marshal(stats)

	_, err = zw.Write([]byte(data))
	if err != nil {
		return s, err
	}

	if err := zw.Close(); err != nil {
		return s, err
	}

	return buf.String(), err
}

func (p *PersistObject) getUserData(from uint64, to uint64, address string) (s string, err error) {
	var buf bytes.Buffer
	zw := gzip.NewWriter(&buf)

	user := User{}
	user.Address = strings.ToLower(address)
	betAmount := float32(0.0)
	earnedAmount := float32(0.0)
	wins := make(map[string]uint32)
	losses := make(map[string]uint32)
	ranks := make(map[string]float32)

	//used for achievements detection
	winStreak := 0
	longestWinStreak := 0
	lossStreak := 0
	longestLossStreak := 0
	isDoubleBettor := false
	isTripleBettor := false

	p.mux.Lock()
	for _, race := range server.data.racesData {
		//count only non refunded closed races
		if race.Date >= from && race.Date <= to {
			if !race.Refunded && (strings.Compare(race.Active, "Closed") == 0) {
				participated := false
				won := false

				betsCount := make(map[string]bool)

				for _, bet := range race.Bets {
					betFrom := strings.ToLower(bet.From)
					won = Contains(race.WinnerHorses, bet.Horse)
					if won {
						wins[betFrom]++
					} else {
						losses[betFrom]++
					}
					//make sure this user exists in the ranks map for later
					ranks[betFrom] = 0

					if strings.Compare(betFrom, user.Address) == 0 {
						participated = true
						betAmount += bet.Value

						if won {
							if longestLossStreak < lossStreak {
								longestLossStreak = lossStreak
							}
							lossStreak = 0
							winStreak++
							user.Horseys = append(user.Horseys, Horsey{RaceNumber: race.RaceNumber, RaceAddress: race.ContractID, Symbol: bet.Horse})
						} else {
							if longestWinStreak < winStreak {
								longestWinStreak = winStreak
							}
							winStreak = 0
							lossStreak++
						}

						betsCount[bet.Horse] = true
						isTripleBettor = len(betsCount) > 2
						isDoubleBettor = len(betsCount) > 1
					}
				}
				for _, with := range race.Withdraws {
					if strings.Compare(strings.ToLower(with.To), user.Address) == 0 {
						earnedAmount += with.Value
					}
				}

				if participated {
					user.GamesCount++
					if won {
						user.WinsCount++
					} else {
						user.LossesCount++
					}
				}
			}
		}
	}
	p.mux.Unlock()

	user.Benefit = earnedAmount - betAmount

	type usr struct {
		address string
		ratio   float32
	}

	//compute ranks of everyone based on win/loss ratio
	var ranksArray []usr
	for user := range ranks {
		ranksArray = append(ranksArray, usr{address: user, ratio: float32(wins[user]) / float32(losses[user])})
	}

	//sort

	sort.Slice(ranksArray, func(i, j int) bool {
		return ranksArray[i].ratio < ranksArray[j].ratio
	})

	//find user rank
	for i := 0; i < len(ranksArray); i++ {
		if strings.Compare(ranksArray[i].address, user.Address) == 0 {
			user.Rank = uint32(i + 1)
		}
	}

	if user.GamesCount > 0 {
		//handle achievements

		//achievement by games count
		if user.GamesCount >= 100 {
			user.Achievements = append(user.Achievements, Achievement{Label: "Whale"})
		} else if user.GamesCount >= 50 {
			user.Achievements = append(user.Achievements, Achievement{Label: "Addict"})
		} else if user.GamesCount >= 10 {
			user.Achievements = append(user.Achievements, Achievement{Label: "Gambler"})
		}
		//achievement by rank
		if user.Rank == 1 {
			user.Achievements = append(user.Achievements, Achievement{Label: "Alpha and Omega"})
		} else if user.Rank <= 3 {
			user.Achievements = append(user.Achievements, Achievement{Label: "Challenger"})
		} else if user.Rank <= 10 {
			user.Achievements = append(user.Achievements, Achievement{Label: "I know what I'm doing"})
		}
		//achievement by winning streak
		if longestWinStreak >= 10 {
			user.Achievements = append(user.Achievements, Achievement{Label: "This isn't gambling!!!"})
		} else if longestWinStreak >= 5 {
			user.Achievements = append(user.Achievements, Achievement{Label: "How do you do it?!!"})
		} else if longestWinStreak >= 3 {
			user.Achievements = append(user.Achievements, Achievement{Label: "I start to get this game!"})
		}
		//achievement by losing streak
		if longestLossStreak >= 10 {
			user.Achievements = append(user.Achievements, Achievement{Label: "Just stop!!!"})
		} else if longestLossStreak >= 5 {
			user.Achievements = append(user.Achievements, Achievement{Label: "Stop playing!"})
		} else if longestLossStreak >= 3 {
			user.Achievements = append(user.Achievements, Achievement{Label: "Bad luck"})
		}
		//achievement based on bet type
		if isTripleBettor {
			user.Achievements = append(user.Achievements, Achievement{Label: "I take no risks"})
		}
		if isDoubleBettor {
			user.Achievements = append(user.Achievements, Achievement{Label: "Hedging is key"})
		}
		if !isTripleBettor && !isDoubleBettor {
			user.Achievements = append(user.Achievements, Achievement{Label: "One shot is enough"})
		}
	}

	data, err := json.Marshal(user)

	_, err = zw.Write([]byte(data))
	if err != nil {
		return s, err
	}

	if err := zw.Close(); err != nil {
		return s, err
	}

	return buf.String(), err
}

func (p *PersistObject) getRanks(from uint64, to uint64) (s string, err error) {
	var buf bytes.Buffer
	zw := gzip.NewWriter(&buf)

	wins := make(map[string]uint32)
	losses := make(map[string]uint32)
	games := make(map[string]uint32)
	ranks := make(map[string]float32)
	wedgered := make(map[string]float32)
	withdrawn := make(map[string]float32)

	p.mux.Lock()
	for _, race := range server.data.racesData {
		//count only non refunded closed races
		if race.Date >= from && race.Date <= to {
			if !race.Refunded && (strings.Compare(race.Active, "Closed") == 0) {
				won := false
				players := make(map[string]bool)
				for _, bet := range race.Bets {
					betFrom := strings.ToLower(bet.From)
					players[betFrom] = true
					wedgered[betFrom] += bet.Value
					won = Contains(race.WinnerHorses, bet.Horse)
					if won {
						wins[betFrom]++
					} else {
						losses[betFrom]++
					}
					//make sure this user exists in the ranks map for later
					ranks[betFrom] = 0
				}
				for _, with := range race.Withdraws {
					withdrawn[strings.ToLower(with.To)] += with.Value
				}
				for player := range players {
					games[player]++
				}
			}

		}
	}
	p.mux.Unlock()

	type usr struct {
		address string
		ratio   float32
	}

	//compute ranks of everyone based on win/loss ratio
	var ranksArray []usr
	for user := range ranks {
		ranksArray = append(ranksArray, usr{address: user, ratio: float32(wins[user]) / float32(losses[user])})
	}

	//sort

	sort.Slice(ranksArray, func(i, j int) bool {
		return ranksArray[i].ratio < ranksArray[j].ratio
	})

	var ranksList []Rank
	for rank, user := range ranksArray {
		ranksList = append(ranksList, Rank{
			Address:     user.address,
			Benefit:     wedgered[user.address] - withdrawn[user.address],
			Rank:        uint32(rank + 1),
			WinsCount:   wins[user.address],
			LossesCount: losses[user.address],
			GamesCount:  games[user.address]})
	}

	data, err := json.Marshal(ranksList)

	_, err = zw.Write([]byte(data))
	if err != nil {
		return s, err
	}

	if err := zw.Close(); err != nil {
		return s, err
	}

	return buf.String(), err
}

func (p *PersistObject) toCSV(from uint32, to uint32) (s string, err error) {
	records := [][]string{
		{"race_number", "date", "race_duration", "betting_duration", "end_time", "winner_horses", "volume", "refunded"},
	}
	p.mux.Lock()
	cache := NewCache(p.racesData, from, to)
	p.mux.Unlock()
	for _, v := range cache.List {
		var strs []string
		strs = append(strs,
			strconv.FormatInt(int64(v.RaceNumber), 10),
			strconv.FormatUint(v.Date, 10),
			strconv.Itoa(int(v.RaceDuration)),
			strconv.Itoa(int(v.BettingDuration)),
			strconv.Itoa(int(v.EndTime)),
			strings.Join(v.WinnerHorses[:], "&"),
			strconv.FormatFloat(float64(v.Volume), 'f', 2, 32),
			strconv.FormatBool(v.Refunded),
		)
		records = append(records, strs)
	}
	var buffer bytes.Buffer
	wr := csv.NewWriter(&buffer)
	err = wr.WriteAll(records)

	return buffer.String(), err
}

func (r *RaceData) findOdds(horse string) *Odd {
	for i := 0; i < len(r.Odds); i++ {
		if strings.Compare(horse, r.Odds[i].Horse) == 0 {
			return &r.Odds[i]
		}
	}
	return nil
}

func (r *RaceData) toRace() Race {
	return Race{
		ContractID:      r.ContractID,
		Date:            strconv.Itoa(int(r.Date)),
		RaceDuration:    strconv.Itoa(int(r.RaceDuration)),
		BettingDuration: strconv.Itoa(int(r.BettingDuration)),
		EndTime:         strconv.Itoa(int(r.EndTime)),
		RaceNumber:      strconv.Itoa(int(r.RaceNumber)),
		V:               0,
		Active:          r.Active}
}

// NewCache blabla
func NewCache(m map[uint32]RaceData, from uint32, to uint32) (cache *Cache) {
	cache = new(Cache)
	for _, v := range m {
		if (v.RaceNumber >= from) && (v.RaceNumber <= to) {
			cache.List = append(cache.List, v)
		}
	}
	cache.LastUpdate = time.Now().Unix()
	sort.Sort(ByRaceNumber(cache.List))

	return cache
}

// ToMap blabla
func (c *Cache) toMap() map[uint32]RaceData {
	m := make(map[uint32]RaceData)
	for _, v := range c.List[:] {
		m[v.RaceNumber] = v
	}
	return m
}

func Contains(s []string, e string) bool {
	for _, a := range s {
		if strings.Compare(a, e) == 0 {
			return true
		}
	}
	return false
}

func Contains2(s []Withdraw, e string) bool {
	for _, a := range s {
		if strings.Compare(a.To, e) == 0 {
			return true
		}
	}
	return false
}
