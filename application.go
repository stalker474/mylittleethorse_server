package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"log"
	"os"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

var server *Server

var ops uint64
var refreshRate int64 = 10
var wg sync.WaitGroup

var saveNeeded uint32

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "5000"
	}
	server = NewServer()
	var err error

	if server.data.load() != nil {
		log.Println("Failed to fetch database, creating one : ", err)
		log.Println("fetching new data for the first time")
	}

	log.Println("starting updating loop")
	go updateCache()

	log.Println("starting api on port ", port)
	err = server.Serve(port)
	if err != nil {
		log.Println("Failed to start server : ", err)
	}
}

func updateCache() {
	for true {
		log.Println("fetching new data...")
		if !fetchNewData() {
			log.Println("No changes...")
		} else {
			persist()
		}
		time.Sleep(time.Duration(refreshRate) * time.Second)
	}
}

func persist() {
	server.data.save()
	server.resetCache()
	log.Println("Cache Updated")
}

func fetchNewData() bool {
	atomic.StoreUint32(&saveNeeded, 0)

	var races []Race
	var err error

	// get finished races list from ethorse bridge
	log.Println("fetching ethorse bridge archive race list")
	races, err = fetchArchive()
	if err != nil {
		log.Fatal("Error :", err)
		return false
	}
	//make a single list of bridge + ours data
	for _, v := range races {
		raceNumber, err := strconv.Atoi(v.RaceNumber)
		if err != nil {
			log.Fatal("Error :", err)
			return false
		}
		server.data.mux.Lock()
		_, contains := server.data.racesData[uint32(raceNumber)]
		//is this a race we dont have in our list yet?
		if !contains {
			//create it and append
			server.data.racesData[uint32(raceNumber)], err = v.toRaceData()
			if err != nil {
				log.Fatal("Error :", err)
				return false
			}
		}
		server.data.mux.Unlock()
	}
	atomic.AddUint64(&ops, uint64(len(server.data.racesData)))
	for raceNumber := range server.data.racesData {
		fetchRaceData(raceNumber)
	}

	log.Println("DONE")

	return atomic.LoadUint32(&saveNeeded) == 1
}

func fetchRaceData(raceNumber uint32) {
	conn, err := ethclient.Dial("wss://mainnet.infura.io/_ws")
	if err != nil {
		log.Fatalf("Failed to init node: %v", err)
	} else {
		defer conn.Close()
	}
	server.data.mux.Lock()
	race, _ := server.data.racesData[raceNumber]
	server.data.mux.Unlock()
	//Complete flag marks a race with all data up to date and impossible to change
	//Such as all winners withdrew their winnings
	if !race.Complete {
		changed, err := updateRaceData(&race, conn)
		if err != nil {
			log.Println("Failed: race #", race.RaceNumber)
		} else {
			if changed {
				server.data.mux.Lock()
				server.data.racesData[raceNumber] = race
				server.data.mux.Unlock()
				atomic.StoreUint32(&saveNeeded, 1)
			}

			log.Println("Success: race #", raceNumber)
		}
	}
	atomic.AddUint64(&ops, ^uint64(0))
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if strings.Compare(a, e) == 0 {
			return true
		}
	}
	return false
}

func contains2(s []Withdraw, e string) bool {
	for _, a := range s {
		if strings.Compare(a.To, e) == 0 {
			return true
		}
	}
	return false
}

func updateRaceData(race *RaceData, conn *ethclient.Client) (bool, error) {
	original, err := json.Marshal(race)
	if err != nil {
		return false, err
	}

	//add a version number if doesnt exist
	//a version number must be in this format X.X.X
	if len(race.Version) < 5 {
		// Instantiate the contract and display its name
		contract, err := NewBetting(common.HexToAddress(race.ContractID), conn)
		if err != nil {
			return false, err
		}
		race.Version, err = contract.Version(nil)
		if err != nil {
			return false, err
		}
	}

	err = errors.New("dummyError")
	for err != nil {
		if strings.Compare(race.Version, "0.2.2") == 0 {
			err = updateRaceData022(race, conn)
		} else if strings.Compare(race.Version, "0.2.3") == 0 {
			err = updateRaceData023(race, conn)
		} else {
			err = updateRaceData024(race, conn)
		}
		if err != nil {
			log.Println("#", race.RaceNumber, " Error : ", err)
			conn, err = ethclient.Dial("wss://mainnet.infura.io/_ws")
			if err != nil {
				log.Println("#", race.RaceNumber, " Failed to reconnect : ", err)
			}
		}
	}

	//the race data changed, check if its complete now
	race.Complete = true
	for _, bet := range race.Bets {
		if contains(race.WinnerHorses, bet.Horse) || race.Refunded {
			//this bet was won or was refunded
			if !contains2(race.Withdraws, bet.From) {
				race.Complete = false
				break
			}
		}
	}
	now, err := json.Marshal(race)
	if err != nil {
		return false, err
	}

	changed := !bytes.Equal(now, original)
	return changed, err
}

func updateRaceData022(race *RaceData, conn *ethclient.Client) error {
	contract, err := NewBetting022(common.HexToAddress(race.ContractID), conn)
	if err != nil {
		return err
	}

	btcWon, err := contract.WinnerHorse(nil, ToBytes32("BTC"))
	if err != nil {
		return err
	}
	ltcWon, err := contract.WinnerHorse(nil, ToBytes32("LTC"))
	if err != nil {
		return err
	}
	ethWon, err := contract.WinnerHorse(nil, ToBytes32("ETH"))
	if err != nil {
		return err
	}
	deposits, err := contract.Betting022Filterer.FilterDeposit(nil)
	if err != nil {
		return err
	}
	defer deposits.Close()
	withdraws, err := contract.Betting022Filterer.FilterWithdraw(nil)
	if err != nil {
		return err
	}
	defer withdraws.Close()

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

	race.Bets = nil
	for deposits.Next() {
		race.Bets = append(race.Bets, Bet{WeiToEth(deposits.Event.Value), FromBytes32(deposits.Event.Horse)[0:3], deposits.Event.From.Hex()})
	}

	race.Volume = 0
	race.Odds = []Odd{{Value: 0.0, Horse: "BTC"}, {Value: 0.0, Horse: "ETH"}, {Value: 0.0, Horse: "LTC"}}
	poolsMap := make(map[string]float32)
	for _, v := range race.Bets {
		race.Volume += v.Value
		poolsMap[v.Horse] += v.Value
	}

	for key, value := range poolsMap {
		odd := race.findOdds(key)
		if odd != nil {
			odd.Value = race.Volume / value * 0.925
		}
	}

	race.Withdraws = nil
	for withdraws.Next() {
		race.Withdraws = append(race.Withdraws, Withdraw{WeiToEth(withdraws.Event.Value), withdraws.Event.To.Hex()})
	}

	if race.Bets == nil || race.WinnerHorses == nil && race.Active == "Closed" {
		race.Refunded = true
	}
	return nil
}

func updateRaceData023(race *RaceData, conn *ethclient.Client) error {
	contract, err := NewBetting023(common.HexToAddress(race.ContractID), conn)
	if err != nil {
		return err
	}

	btcWon, err := contract.WinnerHorse(nil, ToBytes32("BTC"))
	if err != nil {
		return err
	}
	ltcWon, err := contract.WinnerHorse(nil, ToBytes32("LTC"))
	if err != nil {
		return err
	}
	ethWon, err := contract.WinnerHorse(nil, ToBytes32("ETH"))
	if err != nil {
		return err
	}
	deposits, err := contract.Betting023Filterer.FilterDeposit(nil)
	if err != nil {
		return err
	}
	defer deposits.Close()
	withdraws, err := contract.Betting023Filterer.FilterWithdraw(nil)
	if err != nil {
		return err
	}
	defer withdraws.Close()
	refunds, err := contract.Betting023Filterer.FilterRefundEnabled(nil)
	if err != nil {
		return err
	}
	defer refunds.Close()

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

	race.Bets = nil
	for deposits.Next() {
		race.Bets = append(race.Bets, Bet{WeiToEth(deposits.Event.Value), FromBytes32(deposits.Event.Horse)[0:3], deposits.Event.From.Hex()})
	}

	race.Volume = 0
	race.Odds = []Odd{{Value: 0.0, Horse: "BTC"}, {Value: 0.0, Horse: "ETH"}, {Value: 0.0, Horse: "LTC"}}
	poolsMap := make(map[string]float32)
	for _, v := range race.Bets {
		race.Volume += v.Value
		poolsMap[v.Horse] += v.Value
	}

	for key, value := range poolsMap {
		odd := race.findOdds(key)
		if odd != nil {
			odd.Value = race.Volume / value * 0.925
		}
	}

	race.Withdraws = nil
	for withdraws.Next() {
		race.Withdraws = append(race.Withdraws, Withdraw{WeiToEth(withdraws.Event.Value), withdraws.Event.To.Hex()})
	}

	race.Refunded = refunds.Next()

	if race.Bets == nil || race.WinnerHorses == nil && race.Active == "Closed" {
		race.Refunded = true
	}
	return nil
}

func updateRaceData024(race *RaceData, conn *ethclient.Client) error {
	contract, err := NewBetting024(common.HexToAddress(race.ContractID), conn)
	if err != nil {
		return err
	}

	btcWon, err := contract.WinnerHorse(nil, ToBytes32("BTC"))
	if err != nil {
		return err
	}
	ltcWon, err := contract.WinnerHorse(nil, ToBytes32("LTC"))
	if err != nil {
		return err
	}
	ethWon, err := contract.WinnerHorse(nil, ToBytes32("ETH"))
	if err != nil {
		return err
	}
	deposits, err := contract.Betting024Filterer.FilterDeposit(nil)
	if err != nil {
		return err
	}
	defer deposits.Close()
	withdraws, err := contract.Betting024Filterer.FilterWithdraw(nil)
	if err != nil {
		return err
	}
	defer withdraws.Close()
	refunds, err := contract.Betting024Filterer.FilterRefundEnabled(nil)
	if err != nil {
		return err
	}
	defer refunds.Close()

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

	race.Bets = nil
	for deposits.Next() {
		race.Bets = append(race.Bets, Bet{WeiToEth(deposits.Event.Value), FromBytes32(deposits.Event.Horse)[0:3], deposits.Event.From.Hex()})
	}

	race.Volume = 0
	race.Odds = []Odd{{Value: 0.0, Horse: "BTC"}, {Value: 0.0, Horse: "ETH"}, {Value: 0.0, Horse: "LTC"}}
	poolsMap := make(map[string]float32)
	for _, v := range race.Bets {
		race.Volume += v.Value
		poolsMap[v.Horse] += v.Value
	}

	for key, value := range poolsMap {
		odd := race.findOdds(key)
		if odd != nil {
			odd.Value = race.Volume / value * 0.925
		}
	}

	race.Withdraws = nil
	for withdraws.Next() {
		race.Withdraws = append(race.Withdraws, Withdraw{WeiToEth(withdraws.Event.Value), withdraws.Event.To.Hex()})
	}

	race.Refunded = refunds.Next()

	if race.Bets == nil || race.WinnerHorses == nil && race.Active == "Closed" {
		race.Refunded = true
	}
	return nil
}
