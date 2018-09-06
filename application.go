package main

import (
	"bytes"
	"encoding/json"
	"log"
	"os"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

var node *Node
var server *Server

var ops uint64
var fullRefresh uint32
var refreshRate int64 = 60

var wg sync.WaitGroup
var saveNeeded uint32

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "5000"
	}
	server = NewServer()
	var err error

	node, err = NewNode("https://mainnet.infura.io")
	if err != nil {
		log.Fatalf("Failed to init node: %v", err)
	}

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
		if atomic.LoadUint32(&fullRefresh) == 1 {
			log.Println("performing full blockchain data refresh")
			fetchNewData(true)
			persist()
			atomic.SwapUint32(&fullRefresh, 0)
		}
		if !fetchNewData(false) {
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

func fetchNewData(full bool) bool {
	atomic.StoreUint32(&saveNeeded, 0)

	var races []Race
	var err error
	if !full {
		// get finished races list from ethorse bridge
		log.Println("fetching ethorse bridge archive race list")
		races, err = fetchArchive()
		if err != nil {
			log.Fatal("Error :", err)
			return false
		}
		for _, v := range races {
			number, err := strconv.Atoi(v.RaceNumber)
			if err != nil {
				log.Fatal("Error :", err)
				return false
			}
			date, err := strconv.Atoi(v.Date)
			if err != nil {
				log.Fatal("Error :", err)
				return false
			}
			if int(uint32(number)) != number {
				log.Fatal("Invalid race number")
				return false
			}
			//we fully fetch a race only if unknown number of the race is live
			if (!server.data.contains(uint32(number))) || (strings.Compare(v.Active, "Closed") != 0) { //new value or live race
				log.Println("try to get it : #", number)
				wg.Add(1)
				atomic.AddUint64(&ops, 1)
				go asyncFetchRaceData(v, uint32(number), node)
			} else {
				elapsed := time.Now().Unix() - int64(date)
				if (elapsed < 48*60*60) || full {
					log.Println("Get BS data again : #", number)
					wg.Add(1)
					atomic.AddUint64(&ops, 1)
					go asyncUpdateRaceData(uint32(number), node)
				}
			}

		}
	} else {
		wg.Add(len(server.data.racesData))
		atomic.AddUint64(&ops, uint64(len(server.data.racesData)))
		//its a full refresh, we reload all race data from blockchain
		for _, value := range server.data.racesData {
			go asyncUpdateRaceData(value.RaceNumber, node)
		}
		atomic.StoreUint32(&saveNeeded, 1) //always save
	}

	wg.Wait()
	log.Println("DONE")

	return atomic.LoadUint32(&saveNeeded) == 1
}

func asyncFetchRaceData(race Race, raceNumber uint32, node *Node) {
	defer wg.Done()
	newRaceData, err := fetchRaceData(&race, node)
	if err != nil {
		log.Println("FAILED COMPLETELY: race #", race.RaceNumber)
	} else {
		server.data.mux.Lock()
		server.data.racesData[raceNumber] = newRaceData
		server.data.mux.Unlock()
		log.Println("Success: ", raceNumber)
	}

	atomic.AddUint64(&ops, ^uint64(0))

	atomic.StoreUint32(&saveNeeded, 1)
}

func asyncUpdateRaceData(raceNumber uint32, node *Node) {
	defer wg.Done()
	server.data.mux.Lock()
	race, _ := server.data.racesData[raceNumber]
	server.data.mux.Unlock()
	changed, err := updateRaceData(&race, node)
	if err != nil {
		log.Println("FAILED COMPLETELY: race #", race.RaceNumber)
	} else {
		if changed {
			server.data.mux.Lock()
			server.data.racesData[raceNumber] = race
			server.data.mux.Unlock()
			atomic.StoreUint32(&saveNeeded, 1)
		}

		log.Println("Success: ", raceNumber, " value:", atomic.LoadUint64(&ops))
	}

	atomic.AddUint64(&ops, ^uint64(0))
}

func fetchRaceData(race *Race, node *Node) (RaceData, error) {
	data := RaceData{}
	var err error

	data.ContractID = race.ContractID
	date, err := strconv.Atoi(race.Date)
	if err != nil {
		return data, err
	}
	raceNumber, err := strconv.Atoi(race.RaceNumber)
	if err != nil {
		return data, err
	}
	raceDuration, err := strconv.Atoi(race.RaceDuration)
	if err != nil {
		return data, err
	}
	bettingDuration, err := strconv.Atoi(race.BettingDuration)
	if err != nil {
		return data, err
	}
	endTime, err := strconv.Atoi(race.EndTime)
	if err != nil {
		return data, err
	}

	data.Date = uint64(date)
	data.RaceDuration = uint64(raceDuration)
	data.BettingDuration = uint64(bettingDuration)
	data.EndTime = uint64(endTime)
	data.RaceNumber = uint32(raceNumber)
	data.Version = uint64(race.V)
	data.Active = race.Active

	_, err = updateRaceData(&data, node)

	return data, err
}

func updateRaceData(race *RaceData, node *Node) (bool, error) {
	var err error

	original, err := json.Marshal(race)
	if err != nil {
		return false, err
	}

	// Instantiate the contract and display its name
	contract, err := NewBetting(common.HexToAddress(race.ContractID), node.Conn)
	if err != nil {
		return false, err
	}

	queries := make(map[string]bool)

	var btcWon, ltcWon, ethWon bool
	var deposits *BettingDepositIterator
	var withdraws *BettingWithdrawIterator
	var refunds *BettingRefundEnabledIterator
	changed := false

	for !queries["WinnerHorseBTC"] || !queries["WinnerHorseLTC"] || !queries["WinnerHorseETH"] || !queries["Bets"] || !queries["Withdraws"] || !queries["Refund"] {

		if !queries["WinnerHorseBTC"] {
			btcWon, err = contract.WinnerHorse(nil, ToBytes32("BTC"))
			queries["WinnerHorseBTC"] = (err == nil)
		}

		if !queries["WinnerHorseLTC"] {
			ltcWon, err = contract.WinnerHorse(nil, ToBytes32("LTC"))
			queries["WinnerHorseLTC"] = (err == nil)
		}

		if !queries["WinnerHorseETH"] {
			ethWon, err = contract.WinnerHorse(nil, ToBytes32("ETH"))
			queries["WinnerHorseETH"] = (err == nil)
		}

		if !queries["Bets"] {
			deposits, err = contract.BettingFilterer.FilterDeposit(&bind.FilterOpts{Start: 5000000, End: nil, Context: nil})
			queries["Bets"] = (err == nil)
		}

		if !queries["Withdraws"] {
			withdraws, err = contract.BettingFilterer.FilterWithdraw(&bind.FilterOpts{Start: 5000000, End: nil, Context: nil})
			queries["Withdraws"] = (err == nil)
		}

		if !queries["Refund"] {
			refunds, err = contract.BettingFilterer.FilterRefundEnabled(&bind.FilterOpts{Start: 5000000, End: nil, Context: nil})
			queries["Refund"] = (err == nil)
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

		if queries["Bets"] && (deposits != nil) {
			race.Bets = nil
			for deposits.Next() {
				race.Bets = append(race.Bets, Bet{WeiToEth(deposits.Event.Value), FromBytes32(deposits.Event.Horse)[0:3], deposits.Event.From.Hex()})
			}
			deposits.Close()
			deposits = nil

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
		}

		if queries["Withdraws"] && (withdraws != nil) {
			race.Withdraws = nil
			for withdraws.Next() {
				race.Withdraws = append(race.Withdraws, Withdraw{WeiToEth(withdraws.Event.Value), withdraws.Event.To.Hex()})
			}
			withdraws.Close()
			withdraws = nil
		}

		if queries["Refund"] && (refunds != nil) {
			race.Refunded = refunds.Next()
			refunds.Close()

			refunds = nil
		}

		if race.Bets == nil || race.WinnerHorses == nil && race.Active == "Closed" {
			race.Refunded = true
		}
	}

	now, err := json.Marshal(race)
	if err != nil {
		return false, err
	}

	changed = !bytes.Equal(now, original)

	return changed, nil
}
