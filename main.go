package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"google.golang.org/appengine"
)

var node *Node
var db Database

func main() {
	var err error
	node, err = NewNode(ipc)
	if err != nil {
		log.Fatalf("Failed to init node: %v", err)
	}
	err = db.Load()
	if err != nil {
		log.Println("Failed to fetch database, creating one: %v", err)
	}
	log.Println("fetching new data for the first time")
	fetchNewData()
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
		log.Println("Cache Updated")
		db.Save()
		time.Sleep(1 * time.Minute)
	}
}

func fetchNewData() bool {
	var changed bool
	// get finished races list
	log.Println("fetching ethorse bridge archive race list")
	races, err := fetchArchive()
	if err != nil {
		log.Fatal("Error : %v", err)
		return false
	}

	// update new data
	log.Println("looping through new races to handle")
	for i := len(RaceCache.List); i < len(races); i++ {
		if i > 100 {
			continue
		}
		index := i
		log.Println(index)
		race, err := fetchRaceData(&races[index], node)
		if err != nil {
			log.Fatal("Error : %v", err)
		}
		RaceCache.List = append(RaceCache.List, race)
		log.Println("New race added")
		changed = true
	}

	return changed
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
