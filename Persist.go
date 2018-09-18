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

// Bet blabla
type Bet struct {
	Value float32 `json:"value"`
	Horse string  `json:"horse"`
	From  string  `json:"from"`
}

// odd blabla
type Odd struct {
	Value float32 `json:"value"`
	Horse string  `json:"horse"`
}

// Withdraw blabla
type Withdraw struct {
	Value float32 `json:"value"`
	To    string  `json:"to"`
}

// RaceData blabla
type RaceData struct {
	ContractID      string     `json:"contractid"`
	Date            uint64     `json:"date"`
	RaceDuration    uint64     `json:"race_duration"`
	BettingDuration uint64     `json:"betting_duration"`
	EndTime         uint64     `json:"end_time"`
	RaceNumber      uint32     `json:"race_number"`
	Version         string     `json:"version"`
	WinnerHorses    []string   `json:"winner_horses"`
	Odds            []Odd      `json:"odds"`
	Bets            []Bet      `json:"bets"`
	Withdraws       []Withdraw `json:"withdraws"`
	Volume          float32    `json:"volume"`
	Refunded        bool       `json:"refunded"`
	Active          string     `json:"active"`
	Complete        bool       `json:"complete"`
}

// Cache blabla
type Cache struct {
	List       []RaceData `json:"list"`
	LastUpdate int64      `json:"last_update"`
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
	//var buf bytes.Buffer
	//zw := gzip.NewWriter(&buf)

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

	/*_, err = zw.Write([]byte(data))
	if err != nil {
		return s, err
	}

	if err := zw.Close(); err != nil {
		return s, err
	}*/

	//return buf.String(), err
	return string(data), err
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
