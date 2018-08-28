package main

import (
	"log"
	"os"
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
	server.Serve(port)
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
		for _, v := range races[:] {
			number := v.RaceNumber
			if err != nil {
				log.Fatal("Error :", err)
				return false
			}
			if uint64(uint32(number)) != uint64(number) {
				log.Fatal("Invalid race number")
				return false
			}

			if !server.data.contains(uint32(number)) { //new value
				log.Println("I dont have this race in cache, try to get it : #", number)
				wg.Add(1)
				atomic.AddUint64(&ops, 1)
				go asyncFetchRaceData(v, uint32(number), node)
			} else {
				elapsed := time.Now().Unix() - int64(v.Date)
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
		server.data.racesData[raceNumber] = newRaceData
		log.Println("Success: ", raceNumber)
	}

	atomic.AddUint64(&ops, ^uint64(0))

	atomic.StoreUint32(&saveNeeded, 1)
}

func asyncUpdateRaceData(raceNumber uint32, node *Node) {
	defer wg.Done()
	race, _ := server.data.racesData[raceNumber]
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
	data.Date = race.Date
	data.RaceDuration = race.RaceDuration
	data.BettingDuration = race.BettingDuration
	data.EndTime = race.EndTime
	data.RaceNumber = race.RaceNumber
	data.Version = race.V

	_, err = updateRaceData(&data, node)

	return data, err
}

func updateRaceData(race *RaceData, node *Node) (bool, error) {
	var err error

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

		prevLenBets := len(race.Bets)
		prevLenWithdraws := len(race.Withdraws)
		prevLenWinners := len(race.WinnerHorses)

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
			for _, v := range race.Bets {
				race.Volume += v.Value
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

		if queries["Refunded"] && (refunds != nil) {
			race.Refunded = refunds.Next()
		}

		changed = (prevLenBets != len(race.Bets)) || (prevLenWithdraws != len(race.Withdraws)) || (prevLenWinners != len(race.WinnerHorses))
	}

	return changed, nil
}
