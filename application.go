package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"sync"
	"sync/atomic"
	"time"
)

var node *Node
var db Database

var ops uint64
var fullRefresh uint32

var wg sync.WaitGroup
var saveNeeded uint32

var listFailedMutex sync.Mutex
var listFailed []uint32

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "5000"
	}

	var err error
	node, err = NewNode("https://mainnet.infura.io")
	if err != nil {
		log.Fatalf("Failed to init node: %v", err)
	}
	RaceCache = make(map[uint32]RaceData)
	err = db.Load()
	if err != nil {
		log.Println("Failed to fetch database, creating one : ", err)
		log.Println("fetching new data for the first time")
	}

	log.Println("starting updating loop")
	go updateCache()
	log.Println("starting api on port ", port)

	http.HandleFunc("/api/json", func(w http.ResponseWriter, r *http.Request) {
		enableCors(&w)
		_, zipped := r.URL.Query()["gz"]
		if zipped {
			fmt.Fprintln(w, RaceCacheString.GetZIP())
		} else {
			fmt.Fprintln(w, RaceCacheString.GetJSON())
		}

	})

	http.HandleFunc("/api/csv", func(w http.ResponseWriter, r *http.Request) {
		enableCors(&w)
		fmt.Fprintln(w, RaceCacheString.GetCSV())
	})

	http.HandleFunc("/api/admin", func(w http.ResponseWriter, r *http.Request) {
		enableCors(&w)
		keys, ok := r.URL.Query()["method"]

		if !ok || len(keys[0]) < 1 {
			fmt.Fprintln(w, "Missing method")
			fmt.Fprintln(w, "updatedb : refreshes json cache data from current state")
			fmt.Fprintln(w, "recache : refresh ALL data keeping only old ethorse bridge info")
			fmt.Fprintln(w, "report : gives info about current update status")
			return
		}

		switch keys[0] {
		case "updatedb":
			db.Save()
			fmt.Fprintln(w, "Recached")
			return
		case "recache":
			atomic.StoreUint32(&fullRefresh, 1)
			fmt.Fprintln(w, "Full recache ordered")
			return
		case "report":
			listFailedMutex.Lock()
			res, _ := json.Marshal(listFailed)
			listFailedMutex.Unlock()
			fmt.Fprintln(w, "Failed races:"+string(res))

			fmt.Fprintln(w, "Left to process: ", atomic.LoadUint64(&ops))
			return
		default:
			fmt.Fprintln(w, "Unknown method")
		}
	})

	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
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

		}
		atomic.StoreUint32(&fullRefresh, 0)
		time.Sleep(1 * time.Minute)
	}
}

func persist() {
	db.Save()
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
			_, exists := RaceCache[uint32(number)]
			if !exists { //new value
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
		wg.Add(len(RaceCache))
		atomic.AddUint64(&ops, uint64(len(RaceCache)))
		//its a full refresh, we reload all race data from blockchain
		for _, value := range RaceCache {
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
	newRaceData, err := fetchRaceData(&race, node)
	retry := 0
	for err != nil {
		if retry > 5 {
			log.Println("FAILED COMPLETELY: race #", race.RaceNumber)
			listFailedMutex.Lock()
			listFailed = append(listFailed, raceNumber)
			listFailedMutex.Unlock()
			return
		}
		log.Println("Failed: race #", race.RaceNumber, " retry :", retry)
		newRaceData, err = fetchRaceData(&race, node)
		retry++
	}

	RaceCache[raceNumber] = newRaceData
	atomic.AddUint64(&ops, ^uint64(0))
	log.Println("Success: ", raceNumber)

	atomic.StoreUint32(&saveNeeded, 1)
}

func asyncUpdateRaceData(raceNumber uint32, full bool, node *Node) {
	defer wg.Done()
	race, _ := RaceCache[raceNumber]
	changed, err := updateRaceData(&race, full, node)
	retry := 0
	for err != nil {
		if retry > 3 {
			log.Println("FAILED COMPLETELY: race #", race.RaceNumber)
			return
		}
		log.Println("Failed: race #", race.RaceNumber, " retry :", retry)
		changed, err = updateRaceData(&race, full, node)
		retry++
	}

	RaceCache[raceNumber] = race
	atomic.AddUint64(&ops, ^uint64(0))
	log.Println("Success: ", raceNumber, " value:", atomic.LoadUint64(&ops))
	if changed {
		atomic.StoreUint32(&saveNeeded, 1)
	}
}
