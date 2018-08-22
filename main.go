package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"
	"sync/atomic"
	"time"

	"google.golang.org/appengine"
)

var node *Node
var db Database

var ops uint64

func main() {
	var err error
	node, err = NewNode("https://mainnet.infura.io")
	if err != nil {
		log.Fatalf("Failed to init node: %v", err)
	}
	err = db.Load()
	if err != nil {
		log.Println("Failed to fetch database, creating one: %v", err)
		log.Println("fetching new data for the first time")
		fetchNewData()
	}

	log.Println("starting updating loop")
	go updateCache()
	log.Println("starting api")
	http.HandleFunc("/", handle)
	appengine.Main()
}

func updateCache() {
	for true {
		log.Println("fetching new data...")
		if !fetchNewData() {
			log.Println("checking latest withdraws...")
			checkForLateWithdraws(10) //last 10 races
		}
		persist()
		time.Sleep(1 * time.Minute)
	}
}

func persist() {
	db.Save()
	log.Println("Cache Updated")
}

var wg sync.WaitGroup

func fetchNewData() bool {
	var changed bool
	// get finished races list
	log.Println("fetching ethorse bridge archive race list")
	races, err := fetchArchive()
	if err != nil {
		log.Fatal("Error : %v", err)
		return false
	}

	races = races[0:3]

	var workersCount = len(races) - len(RaceCache.List)
	log.Println("Creating race data fetching workers (count) :", workersCount)
	// update new data
	log.Println("looping through new races to handle")
	wg.Add(workersCount)
	for i := len(RaceCache.List); i < len(races); i++ {
		index := i
		log.Println(index)
		var race RaceData
		RaceCache.List = append(RaceCache.List, race)
		go asyncFetchRaceData(races[index], &RaceCache.List[len(RaceCache.List)-1], node, messages)
		changed = true
	}

	wg.Wait()

	return changed
}

func asyncFetchRaceData(race Race, raceData *RaceData, node *Node, messages chan int) {
	defer wg.Done()
	newRaceData, err := fetchRaceData(&race, node)
	for err != nil {
		//log.Println("Error 2: ", err)
		log.Println("Failed: ", race.RaceNumber)
		newRaceData, err = fetchRaceData(&race, node)
	}
	*raceData = newRaceData
	atomic.AddUint64(&ops, 1)
	log.Println("Success: ", race.RaceNumber, " value:", atomic.LoadUint64(&ops))
	messages <- 3
}

func checkForLateWithdraws(lastN int) {
	length := len(RaceCache.List)
	for i, v := range RaceCache.List[:] {
		if i < length-lastN {
			updateRaceData(&v, node)
		}
	}
}

func handle(w http.ResponseWriter, r *http.Request) {
	resp, err := json.Marshal(RaceCache)
	if err != nil {
		log.Fatalf("Failed to marshal cache to json: %v", err)
	}

	fmt.Fprintln(w, string(resp))
}
