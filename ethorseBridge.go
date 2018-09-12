package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
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

func (r *Race) toRaceData() (RaceData, error) {
	data := RaceData{}
	var err error

	data.ContractID = r.ContractID
	date, err := strconv.Atoi(r.Date)
	if err != nil {
		return data, err
	}
	raceNumber, err := strconv.Atoi(r.RaceNumber)
	if err != nil {
		return data, err
	}
	raceDuration, err := strconv.Atoi(r.RaceDuration)
	if err != nil {
		return data, err
	}
	bettingDuration, err := strconv.Atoi(r.BettingDuration)
	if err != nil {
		return data, err
	}
	endTime, err := strconv.Atoi(r.EndTime)
	if err != nil {
		return data, err
	}

	data.Date = uint64(date)
	data.RaceDuration = uint64(raceDuration)
	data.BettingDuration = uint64(bettingDuration)
	data.EndTime = uint64(endTime)
	data.RaceNumber = uint32(raceNumber)
	data.Active = r.Active
	data.Complete = false

	return data, nil
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
