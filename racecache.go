package main

import (
	"log"
	"strconv"
	"sync"

	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

// enumerated types
const (
	BTC Coin = "BTC"
	LTC Coin = "LTC"
	ETH Coin = "ETH"
)

// Coin blabla
type Coin string

// RaceCacheJSON blabla
var RaceCacheJSON SafeCache

// SafeCache is the current cache state
type SafeCache struct {
	RaceCacheJSON string
	mux           sync.Mutex
}

// Set Thread safe value set
func (cache *SafeCache) Set(value string) {
	cache.mux.Lock()
	cache.RaceCacheJSON = value
	cache.mux.Unlock()
}

// Get thread safe value get
func (cache *SafeCache) Get() (str string) {
	cache.mux.Lock()
	str = cache.RaceCacheJSON
	cache.mux.Unlock()
	return str
}

// RaceCache blabla
var RaceCache map[uint32]RaceData

func coinFromString(value string) Coin {
	if strings.Compare(value[0:3], "BTC") == 0 {
		return BTC
	}
	if strings.Compare(value[0:3], "LTC") == 0 {
		return LTC
	}
	if strings.Compare(value[0:3], "ETH") == 0 {
		return ETH
	}
	//log.Fatalf("Invalid coin value")
	log.Println("Error :" + value)
	return BTC
}

// Bet blabla
type Bet struct {
	Value float32 `json:"value"`
	Horse Coin    `json:"horse"`
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
	WinnerHorses    []Coin     `json:"winner_horses"`
	Bets            []Bet      `json:"bets"`
	Withdraws       []Withdraw `json:"withdraws"`
	Volume          float32    `json:"volume"`
	Refunded        bool       `json:"refunded"`
}

// Cache blabla
type Cache struct {
	List []RaceData `json:"list"`
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

	// Instantiate the contract and display its name
	contract, err := NewBetting(common.HexToAddress(race.ContractID), node.Conn)
	if err != nil {
		return data, err
	}

	btcWon, err := contract.WinnerHorse(nil, ToBytes32(string(BTC)))
	if err != nil {
		return data, err
	}
	ltcWon, err := contract.WinnerHorse(nil, ToBytes32(string(LTC)))
	if err != nil {
		return data, err
	}
	ethWon, err := contract.WinnerHorse(nil, ToBytes32(string(ETH)))
	if err != nil {
		return data, err
	}

	_, err = contract.Chronus(nil)
	if err != nil {
		return data, err
	}
	deposits, err := contract.BettingFilterer.FilterDeposit(&bind.FilterOpts{5000000, nil, nil})
	if err != nil {
		return data, err
	}
	for deposits.Next() {
		data.Bets = append(data.Bets, Bet{WeiToEth(deposits.Event.Value), coinFromString(FromBytes32(deposits.Event.Horse)), deposits.Event.From.Hex()})
	}
	deposits.Close()

	if btcWon {
		data.WinnerHorses = append(data.WinnerHorses, BTC)
	}
	if ltcWon {
		data.WinnerHorses = append(data.WinnerHorses, LTC)
	}
	if ethWon {
		data.WinnerHorses = append(data.WinnerHorses, ETH)
	}
	//data.Refunded = chronus.VoidedBet

	data.Volume = 0
	//if all people who played withdrew, its a refunded race
	for _, v := range data.Bets[:] {
		data.Volume += v.Value
	}

	updateRaceData(&data, node)

	return data, nil
}

func updateRaceData(race *RaceData, node *Node) error {
	var err error

	// Instantiate the contract and display its name
	contract, err := NewBetting(common.HexToAddress(race.ContractID), node.Conn)
	if err != nil {
		return err
	}

	withdraws, err := contract.BettingFilterer.FilterWithdraw(&bind.FilterOpts{6059602, nil, nil})
	if err != nil {
		return err
	}

	race.Withdraws = nil

	for withdraws.Next() {
		race.Withdraws = append(race.Withdraws, Withdraw{WeiToEth(withdraws.Event.Value), withdraws.Event.To.Hex()})
	}
	withdraws.Close()

	playersMap := make(map[string]bool)

	for _, v := range race.Withdraws[:] {
		playersMap[v.To] = true
	}

	refunded := true
	//if all people who played withdrew, its a refunded race
	for _, v := range race.Bets[:] {
		if !playersMap[v.From] {
			refunded = false
		}
	}

	race.Refunded = refunded

	return nil
}
