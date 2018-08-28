package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

var endpointArchive = "https://bet.ethorse.com/bridge/getHistoricRaces"

// Race blabla
type Race struct {
	ContractID      string `json:"contractid"`
	Date            uint64 `json:"date"`
	RaceDuration    uint64 `json:"race_duration"`
	BettingDuration uint64 `json:"betting_duration"`
	EndTime         uint64 `json:"end_time"`
	RaceNumber      uint32 `json:"race_number"`
	V               uint32 `json:"__v"`
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
