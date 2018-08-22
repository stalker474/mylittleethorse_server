package main

// Coin blabla
type Coin string

// enumerated types
const (
	BTC Coin = "BTC"
	LTC Coin = "LTC"
	ETH Coin = "ETH"
)

// Bet blabla
type Bet struct {
	Value float32 `json:"value"`
	Horse Coin    `json:"horse"`
	From  string  `json:"from"`
}

// Withdraw blabla
type Withdraw struct {
	Value float32 `json:"value"`
	To    string  `json:"to"`
}

// RaceData blabla
type RaceData struct {
	Contractid      string     `json:"contractid"`
	Date            string     `json:"date"`
	RaceDuration    string     `json:"race_duration"`
	BettingDuration string     `json:"betting_duration"`
	EndTime         string     `json:"end_time"`
	RaceNumber      string     `json:"race_number"`
	WinnerHorses    []Coin     `json:"winner_horses"`
	Bets            []Bet      `json:"bets"`
	Withdraws       []Withdraw `json:"withdraws"`
	Volume          float32    `json:"volume"`
	Refunded        bool       `json:"refunded"`
}
