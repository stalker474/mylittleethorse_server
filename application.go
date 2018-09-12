package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"log"
	"os"
	"strconv"
	"strings"
	"sync/atomic"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

var server *Server

var ops uint64
var fullRefresh uint32
var refreshRate int64 = 10

var saveNeeded uint32
var conn *ethclient.Client

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
		var err error
		conn, err = ethclient.Dial("wss://mainnet.infura.io/_ws")
		if err != nil {
			log.Fatalf("Failed to init node: %v", err)
		} else {
			defer conn.Close()
		}

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
			elapsed := time.Now().Unix() - int64(date)
			//we fully fetch only new races
			if (elapsed < 24*60*60) || full {
				log.Println("Fetching : #", number)
				atomic.AddUint64(&ops, 1)
				syncFetchRaceData(v, uint32(number))
			}
		}
	} else {
		atomic.AddUint64(&ops, uint64(len(server.data.racesData)))
		//its a full refresh, we reload all race data from blockchain
		for _, value := range server.data.racesData {
			syncUpdateRaceData(value.RaceNumber)
		}
		atomic.StoreUint32(&saveNeeded, 1) //always save
	}

	log.Println("DONE")

	return atomic.LoadUint32(&saveNeeded) == 1
}

func syncFetchRaceData(race Race, raceNumber uint32) {
	//defer wg.Done()
	newRaceData, err := fetchRaceData(&race)
	if err != nil {
		log.Println("FAILED COMPLETELY: race #", race.RaceNumber, " error:", err)
	} else {
		server.data.mux.Lock()
		server.data.racesData[raceNumber] = newRaceData
		server.data.mux.Unlock()
		log.Println("Success: ", raceNumber)
	}

	atomic.AddUint64(&ops, ^uint64(0))

	atomic.StoreUint32(&saveNeeded, 1)
}

func syncUpdateRaceData(raceNumber uint32) {
	server.data.mux.Lock()
	race, _ := server.data.racesData[raceNumber]
	server.data.mux.Unlock()
	changed, err := updateRaceData(&race)
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

func fetchRaceData(race *Race) (RaceData, error) {
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
	data.Active = race.Active

	_, err = updateRaceData(&data)

	return data, err
}

func updateRaceData(race *RaceData) (bool, error) {
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
			err = updateRaceData022(race)
		} else if strings.Compare(race.Version, "0.2.3") == 0 {
			err = updateRaceData023(race)
		} else {
			err = updateRaceData024(race)
		}
		if err != nil {
			log.Println("#", race.RaceNumber, " Error : ", err)
			conn, err = ethclient.Dial("wss://mainnet.infura.io/_ws")
			if err != nil {
				log.Println("#", race.RaceNumber, " Failed to reconnect : ", err)
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

func updateRaceData022(race *RaceData) error {
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
	deposits, err := contract.Betting022Filterer.FilterDeposit(&bind.FilterOpts{Start: 6000000, End: nil, Context: nil})
	if err != nil {
		return err
	}
	defer deposits.Close()
	withdraws, err := contract.Betting022Filterer.FilterWithdraw(&bind.FilterOpts{Start: 6000000, End: nil, Context: nil})
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

func updateRaceData023(race *RaceData) error {
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
	deposits, err := contract.Betting023Filterer.FilterDeposit(&bind.FilterOpts{Start: 6000000, End: nil, Context: nil})
	if err != nil {
		return err
	}
	defer deposits.Close()
	withdraws, err := contract.Betting023Filterer.FilterWithdraw(&bind.FilterOpts{Start: 6000000, End: nil, Context: nil})
	if err != nil {
		return err
	}
	defer withdraws.Close()
	refunds, err := contract.Betting023Filterer.FilterRefundEnabled(&bind.FilterOpts{Start: 6000000, End: nil, Context: nil})
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

func updateRaceData024(race *RaceData) error {
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
	deposits, err := contract.Betting024Filterer.FilterDeposit(&bind.FilterOpts{Start: 6000000, End: nil, Context: nil})
	if err != nil {
		return err
	}
	defer deposits.Close()
	withdraws, err := contract.Betting024Filterer.FilterWithdraw(&bind.FilterOpts{Start: 6000000, End: nil, Context: nil})
	if err != nil {
		return err
	}
	defer withdraws.Close()
	refunds, err := contract.Betting024Filterer.FilterRefundEnabled(&bind.FilterOpts{Start: 6000000, End: nil, Context: nil})
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
