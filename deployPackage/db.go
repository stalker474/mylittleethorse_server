package main

import (
	"encoding/json"
	"io/ioutil"
)

var tempDbFile = "mle_db.json"

// Database blabla
type Database struct {
}

// Save blabla
func (Database) Save() error {
	var cache Cache
	for _, v := range RaceCache {
		cache.List = append(cache.List, v)
	}
	resp, err := json.Marshal(cache)
	if err != nil {
		return err
	}

	return ioutil.WriteFile(tempDbFile, resp, 0644)
	RaceCacheJSON = string(resp)

	return nil
}

// Load blabla
func (Database) Load() error {
	jsonText, err := ioutil.ReadFile(tempDbFile)
	if err != nil {
		return err
	}
	var cache Cache
	err = json.Unmarshal(jsonText, &cache)
	RaceCacheJSON = string(jsonText)
	for _, v := range cache.List[:] {
		RaceCache[v.RaceNumber] = v
	}
	return nil
}
