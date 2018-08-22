package main

import (
	"encoding/json"
	"io/ioutil"
)

var tempDbFile = "/tmp/mle_db"

type Database struct {
}

func (Database) Save() error {
	resp, err := json.Marshal(RaceCache)
	if err != nil {
		return err
	}

	return ioutil.WriteFile(tempDbFile, resp, 0644)
}

func (Database) Load() error {
	jsonText, err := ioutil.ReadFile(tempDbFile)
	if err != nil {
		return err
	}
	return json.Unmarshal(jsonText, &RaceCache)
}
