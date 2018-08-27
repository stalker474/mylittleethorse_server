package main

import (
	"errors"
	"log"
	"os"
	"strconv"
	"sync"
	"sync/atomic"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

var node *Node
var db *PersistObject

var ops uint64
var fullRefresh uint32
var refreshRate int64 = 10

var wg sync.WaitGroup
var saveNeeded uint32

var listFailedMutex sync.Mutex
var listFailed []uint32

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "5000"
	}
	var server Server
	db := NewPersistObject()
	var err error

	node, err = NewNode("https://mainnet.infura.io")
	if err != nil {
		log.Fatalf("Failed to init node: %v", err)
	}

	if db.load() != nil {
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
	db.save()
	log.Println("Cache Updated")
}

func fetchNewData(full bool) bool {
	atomic.StoreUint32(&saveNeeded, 0)
	listFailedMutex.Lock()
	listFailed = nil
	listFailedMutex.Unlock()

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
			number, err := strconv.ParseUint(v.RaceNumber, 10, 32)
			if err != nil {
				log.Fatal("Error :", err)
				return false
			}
			if uint64(uint32(number)) != uint64(number) {
				log.Fatal("Invalid race number")
				return false
			}

			if !db.contains(uint32(number)) { //new value
				log.Println("I dont have this race in cache, try to get it : #", number)
				wg.Add(1)
				atomic.AddUint64(&ops, 1)
				go asyncFetchRaceData(v, uint32(number), node)
			} else {
				then, err := strconv.ParseInt(v.Date, 10, 64)
				if err != nil {
					log.Fatal("Error :", err)
					return false
				}
				elapsed := time.Now().Unix() - then
				if (elapsed < 48*60*60) || full {
					log.Println("Get BS data again : #", number)
					wg.Add(1)
					atomic.AddUint64(&ops, 1)
					go asyncUpdateRaceData(uint32(number), full, node)
				}
			}

		}
	} else {
		wg.Add(len(db.racesData))
		atomic.AddUint64(&ops, uint64(len(db.racesData)))
		//its a full refresh, we reload all race data from blockchain
		for _, value := range db.racesData {
			go asyncUpdateRaceData(value.RaceNumber, true, node)
		}
		atomic.StoreUint32(&saveNeeded, 1) //always save
	}

	wg.Wait()
	log.Println("DONE")

	return atomic.LoadUint32(&saveNeeded) == 1
}

func asyncFetchRaceData(race Race, raceNumber uint32, node *Node) {
	defer wg.Done()
	retry := 0
	err := errors.New("")
	var newRaceData RaceData
	for err != nil {
		if retry > 5 {
			log.Println("FAILED COMPLETELY: race #", race.RaceNumber)
			listFailedMutex.Lock()
			listFailed = append(listFailed, raceNumber)
			listFailedMutex.Unlock()
			atomic.AddUint64(&ops, ^uint64(0))
			return
		}
		log.Println("Failed: race #", race.RaceNumber, " retry :", retry)
		newRaceData, err = fetchRaceData(&race, node)
		retry++
	}

	db.racesData[raceNumber] = newRaceData
	atomic.AddUint64(&ops, ^uint64(0))
	log.Println("Success: ", raceNumber)

	atomic.StoreUint32(&saveNeeded, 1)
}

func asyncUpdateRaceData(raceNumber uint32, full bool, node *Node) {
	defer wg.Done()
	race, _ := db.racesData[raceNumber]
	retry := 0
	err := errors.New("")
	changed := false
	for err != nil {
		if retry > 5 {

			log.Println("FAILED COMPLETELY: race #", race.RaceNumber)
			listFailedMutex.Lock()
			listFailed = append(listFailed, raceNumber)
			listFailedMutex.Unlock()
			atomic.AddUint64(&ops, ^uint64(0))
			return
		}
		log.Println("Failed: race #", race.RaceNumber, " retry :", retry)
		changed, err = updateRaceData(&race, full, node)
		retry++
	}

	db.racesData[raceNumber] = race
	atomic.AddUint64(&ops, ^uint64(0))
	log.Println("Success: ", raceNumber, " value:", atomic.LoadUint64(&ops))
	if changed {
		atomic.StoreUint32(&saveNeeded, 1)
	}
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

	_, err = updateRaceData(&data, true, node)

	return data, err
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

	if full { //only fetch deposits during full update
		deposits, err := contract.BettingFilterer.FilterDeposit(&bind.FilterOpts{Start: 5000000, End: nil, Context: nil})
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

	withdraws, err := contract.BettingFilterer.FilterWithdraw(&bind.FilterOpts{Start: 5000000, End: nil, Context: nil})
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
	for _, v := range race.Bets[:] {
		race.Volume += v.Value
	}

	refunds, err := contract.BettingFilterer.FilterRefundEnabled(&bind.FilterOpts{Start: 5000000, End: nil, Context: nil})
	if err != nil {
		return false, err
	}

	race.Refunded = false
	//if refunds was triggered, this will contain a single RefundEnabled event
	for refunds.Next() {
		race.Refunded = true
	}

	return changed, nil
}
