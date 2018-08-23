package main

import (
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

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
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
		fetchNewData()
	}

	log.Println("starting updating loop")
	go updateCache()
	log.Println("starting api on port ", port)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, RaceCacheJSON.Get())
	})

	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func updateCache() {
	for true {
		log.Println("fetching new data...")
		if !fetchNewData() {
			log.Println("No changes...")
		} else {
			persist()
		}

		time.Sleep(1 * time.Minute)
	}
}

func persist() {
	db.Save()
	log.Println("Cache Updated")
}

var wg sync.WaitGroup

func fetchNewData() bool {
	changed := false
	// get finished races list
	log.Println("fetching ethorse bridge archive race list")
	races, err := fetchArchive()
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
			go asyncFetchRaceData(v, uint32(number), node)
			changed = true
		}

	}

	wg.Wait()
	log.Println("DONE")

	return changed
}

func asyncFetchRaceData(race Race, raceNumber uint32, node *Node) {
	defer wg.Done()
	newRaceData, err := fetchRaceData(&race, node)
	retry := 0
	for err != nil {
		if retry > 3 {
			log.Println("FAILED COMPLETELY: race #", race.RaceNumber)
			return
		}
		log.Println("Failed: race #", race.RaceNumber, " retry :", retry)
		newRaceData, err = fetchRaceData(&race, node)
		retry++
	}

	RaceCache[raceNumber] = newRaceData
	atomic.AddUint64(&ops, 1)
	log.Println("Success: ", raceNumber, " value:", atomic.LoadUint64(&ops))
}
