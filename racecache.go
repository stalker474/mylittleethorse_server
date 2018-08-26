package main

import (
	"bytes"
	"compress/zlib"
	"encoding/csv"
	"encoding/json"
	"log"
	"strconv"
	"strings"
	"sync"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

// RaceCacheString blabla
var RaceCacheString SafeCache

// SafeCache is the current cache state
type SafeCache struct {
	RaceCacheJSON    string
	RaceCacheCSV     string
	RaceCacheJSONZIP string
	mux              sync.Mutex
}

// Set Thread safe value set
func (c *SafeCache) Set(cache *Cache) error {
	data, err := json.Marshal(cache)
	if err != nil {
		return err
	}
	c.mux.Lock()
	c.RaceCacheJSON = string(data)

	var b bytes.Buffer
	w := zlib.NewWriter(&b)
	w.Write(data)
	w.Close()

	c.RaceCacheJSONZIP = b.String()

	records := [][]string{
		{"race_number", "date", "race_duration", "betting_duration", "end_time", "winner_horses", "volume", "refunded"},
	}
	for _, v := range cache.List[:] {
		var strs []string

		strs = append(strs,
			strconv.FormatInt(int64(v.RaceNumber), 10),
			strconv.FormatUint(v.Date, 10),
			strconv.FormatUint(v.RaceDuration, 10),
			strconv.FormatUint(v.BettingDuration, 10),
			strconv.FormatUint(v.EndTime, 10),
			strings.Join(v.WinnerHorses[:], "&"),
			strconv.FormatFloat(float64(v.Volume), 'f', 2, 32),
			strconv.FormatBool(v.Refunded),
		)
		records = append(records, strs)
	}
	var buffer bytes.Buffer
	wr := csv.NewWriter(&buffer)
	wr.WriteAll(records)
	c.RaceCacheCSV = buffer.String()
	c.mux.Unlock()
	return nil
}

// GetJSON thread safe value get
func (c *SafeCache) GetJSON() (str string) {
	c.mux.Lock()
	str = c.RaceCacheJSON
	c.mux.Unlock()
	return str
}

// GetZIP thread safe value get
func (c *SafeCache) GetZIP() (str string) {
	c.mux.Lock()
	str = c.RaceCacheJSONZIP
	c.mux.Unlock()
	return str
}

// GetCSV thread safe value get
func (c *SafeCache) GetCSV() (str string) {
	c.mux.Lock()
	str = c.RaceCacheCSV
	c.mux.Unlock()
	return str
}

// RaceCache blabla
var RaceCache map[uint32]RaceData

// Bet blabla
type Bet struct {
	Value float32 `json:"value"`
	Horse string  `json:"horse"`
	From  string  `json:"from"`
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
	WinnerHorses    []string   `json:"winner_horses"`
	Bets            []Bet      `json:"bets"`
	Withdraws       []Withdraw `json:"withdraws"`
	Volume          float32    `json:"volume"`
	Refunded        bool       `json:"refunded"`
}

// Cache blabla
type Cache struct {
	List       []RaceData `json:"list"`
	LastUpdate int64      `json:"last_update"`
}

func fetchRaceData(race *Race, node *Node) (RaceData, error) {
	data := RaceData{}
	var err error

	data.ContractID = race.ContractID
	data.Date, err = strconv.ParseUint(race.Date, 10, 64)
	if err != nil {
		return data, err
	}
	data.RaceDuration, err = strconv.ParseUint(race.RaceDuration, 10, 64)
	if err != nil {
		return data, err
	}
	data.BettingDuration, err = strconv.ParseUint(race.BettingDuration, 10, 64)
	if err != nil {
		return data, err
	}
	data.EndTime, err = strconv.ParseUint(race.EndTime, 10, 64)
	if err != nil {
		return data, err
	}
	raceNumber, err := strconv.ParseUint(race.RaceNumber, 10, 32)
	if err != nil {
		return data, err
	}
	data.RaceNumber = uint32(raceNumber)

	updateRaceData(&data, true, node)

	return data, nil
}

func updateRaceData(race *RaceData, full bool, node *Node) (bool, error) {
	var err error

	changed := false

	// Instantiate the contract and display its name
	contract, err := NewBetting(common.HexToAddress(race.ContractID), node.Conn)
	if err != nil {
		return false, err
	}

	btcWon, err := contract.WinnerHorse(nil, ToBytes32("BTC"))
	if err != nil {
		return false, err
	}
	ltcWon, err := contract.WinnerHorse(nil, ToBytes32("LTC"))
	if err != nil {
		return false, err
	}
	ethWon, err := contract.WinnerHorse(nil, ToBytes32("ETH"))
	if err != nil {
		return false, err
	}

	c, err := contract.Chronus(nil)
	if err != nil {
		return false, err
	}
	if c.BettingDuration > 0 {
		log.Println("DATAAAAAAAA ", c.BettingDuration)
	}

	if full { //only fetch deposits during full update
		deposits, err := contract.BettingFilterer.FilterDeposit(&bind.FilterOpts{5000000, nil, nil})
		if err != nil {
			return false, err
		}
		var newBets []Bet
		for deposits.Next() {
			newBets = append(newBets, Bet{WeiToEth(deposits.Event.Value), FromBytes32(deposits.Event.Horse)[0:3], deposits.Event.From.Hex()})
		}
		race.Bets = newBets
		changed = true
		deposits.Close()
	}

	if btcWon || ltcWon || ethWon {
		race.WinnerHorses = nil
	}

	if btcWon {
		race.WinnerHorses = append(race.WinnerHorses, "BTC")
	}
	if ltcWon {
		race.WinnerHorses = append(race.WinnerHorses, "LTC")
	}
	if ethWon {
		race.WinnerHorses = append(race.WinnerHorses, "ETH")
	}

	withdraws, err := contract.BettingFilterer.FilterWithdraw(&bind.FilterOpts{5000000, nil, nil})
	if err != nil {
		return false, err
	}

	var newWithdraws []Withdraw

	for withdraws.Next() {
		newWithdraws = append(newWithdraws, Withdraw{WeiToEth(withdraws.Event.Value), withdraws.Event.To.Hex()})
	}
	withdraws.Close()

	if len(newWithdraws) > 0 && (len(newWithdraws) != len(race.Withdraws)) {
		race.Withdraws = newWithdraws
		changed = true
	}

	playersMap := make(map[string]bool)

	for _, v := range race.Withdraws[:] {
		playersMap[v.To] = true
	}

	race.Volume = 0
	somebodyWon := false
	winners := strings.Join(race.WinnerHorses, "")
	//if all people who played withdrew, its a refunded race
	for _, v := range race.Bets[:] {
		race.Volume += v.Value
		if strings.Contains(winners, v.Horse) {
			somebodyWon = true
		}
	}

	race.Refunded = (race.Volume == 0) || (race.WinnerHorses == nil) || (len(race.Bets) == 1) || (!somebodyWon)

	return changed, nil
}
