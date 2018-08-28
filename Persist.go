package main

import (
	"bytes"
	"compress/zlib"
	"encoding/csv"
	"encoding/json"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"
)

var tempDbFile = "mle_db.json"

// ByRaceNumber sorter
type ByRaceNumber []RaceData

func (a ByRaceNumber) Len() int           { return len(a) }
func (a ByRaceNumber) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByRaceNumber) Less(i, j int) bool { return a[i].RaceNumber > a[j].RaceNumber }

// PersistObject blabla
type PersistObject struct {
	racesData map[uint32]RaceData
	mux       sync.Mutex
}

// Bet blabla
type Bet struct {
	Value float32 `json:"value"`
	Horse string  `json:"horse"`
	From  string  `json:"from"`
}

// Withdraw blabla
type Withdraw struct {
	Value float32 `json:"value"`
	To    string  `json:"to"`
}

// RaceData blabla
type RaceData struct {
	ContractID      string     `json:"contractid"`
	Date            uint64     `json:"date"`
	RaceDuration    uint64     `json:"race_duration"`
	BettingDuration uint64     `json:"betting_duration"`
	EndTime         uint64     `json:"end_time"`
	RaceNumber      uint32     `json:"race_number"`
	Version         uint32     `json:"version"`
	WinnerHorses    []string   `json:"winner_horses"`
	Bets            []Bet      `json:"bets"`
	Withdraws       []Withdraw `json:"withdraws"`
	Volume          float32    `json:"volume"`
	Refunded        bool       `json:"refunded"`
	Active          string     `json:"active"`
}

// Cache blabla
type Cache struct {
	List       []RaceData `json:"list"`
	LastUpdate int64      `json:"last_update"`
}

// NewPersistObject Create new database
func NewPersistObject() (p *PersistObject) {
	p = new(PersistObject)
	p.racesData = make(map[uint32]RaceData)
	return p
}

// Save blabla
func (p *PersistObject) save() error {
	cache := NewCache(p.racesData, 0, 99999)
	resp, err := json.Marshal(cache)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(tempDbFile, resp, 0644)
}

// Load blabla
func (p *PersistObject) load() error {
	jsonText, err := ioutil.ReadFile(tempDbFile)
	if err != nil {
		return err
	}
	var cache Cache
	err = json.Unmarshal(jsonText, &cache)
	p.racesData = cache.toMap()
	return nil
}

func (p *PersistObject) contains(raceNumber uint32) bool {
	_, exists := p.racesData[raceNumber]
	return exists
}

func (p *PersistObject) toJSON(from uint32, to uint32) (s string, err error) {
	p.mux.Lock()
	data, err := json.Marshal(NewCache(p.racesData, from, to))
	p.mux.Unlock()

	s = string(data)
	return s, err
}

func (p *PersistObject) toLightJSON() (s string, err error) {
	p.mux.Lock()

	var r []Race
	for _, value := range p.racesData {
		r = append(r, Race{
			ContractID:      value.ContractID,
			Date:            value.Date,
			RaceDuration:    value.RaceDuration,
			BettingDuration: value.BettingDuration,
			EndTime:         value.EndTime,
			RaceNumber:      value.RaceNumber,
			V:               value.Version,
			Active:          "Closed"})
	}

	data, err := json.Marshal(r)
	p.mux.Unlock()

	s = string(data)
	return s, err
}

func (p *PersistObject) toZJSON(from uint32, to uint32) (s string, err error) {
	var b bytes.Buffer
	w := zlib.NewWriter(&b)
	data, err := p.toJSON(from, to)
	w.Write([]byte(data))
	w.Close()

	s = string(data)
	return s, err
}

func (p *PersistObject) toCSV(from uint32, to uint32) (s string, err error) {
	records := [][]string{
		{"race_number", "date", "race_duration", "betting_duration", "end_time", "winner_horses", "volume", "refunded"},
	}
	p.mux.Lock()
	cache := NewCache(p.racesData, from, to)
	p.mux.Unlock()
	for _, v := range cache.List {
		var strs []string
		strs = append(strs,
			strconv.FormatInt(int64(v.RaceNumber), 10),
			strconv.FormatUint(v.Date, 10),
			strconv.FormatUint(v.RaceDuration, 10),
			strconv.FormatUint(v.BettingDuration, 10),
			strconv.FormatUint(v.EndTime, 10),
			strings.Join(v.WinnerHorses[:], "&"),
			strconv.FormatFloat(float64(v.Volume), 'f', 2, 32),
			strconv.FormatBool(v.Refunded),
		)
		records = append(records, strs)
	}
	var buffer bytes.Buffer
	wr := csv.NewWriter(&buffer)
	err = wr.WriteAll(records)

	return buffer.String(), err
}

// NewCache blabla
func NewCache(m map[uint32]RaceData, from uint32, to uint32) (cache *Cache) {
	cache = new(Cache)
	for _, v := range m {
		if (v.RaceNumber >= from) && (v.RaceNumber <= to) {
			cache.List = append(cache.List, v)
		}
	}
	cache.LastUpdate = time.Now().Unix()
	sort.Sort(ByRaceNumber(cache.List))

	return cache
}

// ToMap blabla
func (c *Cache) toMap() map[uint32]RaceData {
	m := make(map[uint32]RaceData)
	for _, v := range c.List[:] {
		m[v.RaceNumber] = v
	}
	return m
}
