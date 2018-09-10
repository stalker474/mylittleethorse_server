package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

var endpointArchive = "https://bet.ethorse.com/bridge/getHistoricRaces"
var endpointActive = "https://bet.ethorse.com/bridge/getActiveRaces"

// Race blabla
type Race struct {
	RaceID          string `json:"_id"`
	ContractID      string `json:"contractid"`
	Date            string `json:"date"`
	RaceDuration    string `json:"race_duration"`
	BettingDuration string `json:"betting_duration"`
	EndTime         string `json:"end_time"`
	RaceNumber      string `json:"race_number"`
	V               int32  `json:"__v"`
	Active          string `json:"active"`
}

func fetchArchive() ([]Race, error) {
	var races []Race
	var racesActive []Race
	respArchive, err := http.Get(endpointArchive)
	if err != nil {
		return races, err
	}

	respActive, err := http.Get(endpointActive)
	if err != nil {
		return races, err
	}

	defer respArchive.Body.Close()
	defer respActive.Body.Close()
	bodyArchive, err := ioutil.ReadAll(respArchive.Body)
	if err != nil {
		return races, err
	}
	bodyActive, err := ioutil.ReadAll(respActive.Body)
	if err != nil {
		return races, err
	}

	err = json.Unmarshal(bodyArchive, &races)
	if err != nil {
		return races, err
	}
	err = json.Unmarshal(bodyActive, &racesActive)
	if err != nil {
		return races, err
	}
	//copy active races into archive races
	races = append(racesActive, races...)
	return races, nil
}
