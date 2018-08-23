package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

var endpointArchive = "https://bet.ethorse.com/bridge/getHistoricRaces"

// Race blabla
type Race struct {
	ID              string `json:"_id"`
	ContractID      string `json:"contractid"`
	Date            string `json:"date"`
	RaceDuration    string `json:"race_duration"`
	BettingDuration string `json:"betting_duration"`
	EndTime         string `json:"end_time"`
	RaceNumber      string `json:"race_number"`
	V               int    `json:"__v"`
	Active          string `json:"active"`
}

func fetchArchive() ([]Race, error) {
	var races []Race
	resp, err := http.Get(endpointArchive)
	if err != nil {
		return races, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return races, err
	}

	json.Unmarshal(body, &races)
	return races, nil
}
