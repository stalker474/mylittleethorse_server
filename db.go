package main

import (
	"encoding/json"
	"io/ioutil"
	"sort"
	"time"
)

var tempDbFile = "mle_db.json"

// Database blabla
type Database struct {
}

// ByRaceNumber sorter
type ByRaceNumber []RaceData

func (a ByRaceNumber) Len() int           { return len(a) }
func (a ByRaceNumber) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByRaceNumber) Less(i, j int) bool { return a[i].RaceNumber > a[j].RaceNumber }

// Save blabla
func (Database) Save() error {
	var cache Cache
	for _, v := range RaceCache {
		cache.List = append(cache.List, v)
	}
	cache.LastUpdate = time.Now().Unix()
	sort.Sort(ByRaceNumber(cache.List))
	resp, err := json.Marshal(cache)
	if err != nil {
		return err
	}

	RaceCacheString.Set(&cache)
	return ioutil.WriteFile(tempDbFile, resp, 0644)
}

// Load blabla
func (Database) Load() error {
	jsonText, err := ioutil.ReadFile(tempDbFile)
	if err != nil {
		return err
	}
	var cache Cache
	err = json.Unmarshal(jsonText, &cache)
	RaceCacheString.Set(&cache)
	for _, v := range cache.List[:] {
		RaceCache[v.RaceNumber] = v
	}
	return nil
}
