package main

// Charts blablabla
type Charts struct {
}

// Coin blablabla
type Coin struct {
	Volume    float32 `json:"volume"`
	BetsCount uint32  `json:"bets"`
}

// CoinInfo blablabla
type CoinInfo struct {
	WinsCount uint32 `json:"wins_count"`
}

// Day blablabla
type Day struct {
	Label        string          `json:"label"`
	Date         uint64          `json:"date"`
	Volume       float32         `json:"volume"`
	UserVolume   float32         `json:"user_volume"`
	PlayersCount uint32          `json:"players_count"`
	Coins        map[string]Coin `json:"coins"`
}

// Stats blabla
type Stats struct {
	PeriodBegin       uint64              `json:"period_begin"`
	PeriodEnd         uint64              `json:"period_end"`
	Days              []Day               `json:"days"`
	TotalPlayersCount uint32              `json:"total_players"`
	TotalVolume       float32             `json:"total_volume"`
	TotalRaces        uint32              `json:"total_races"`
	CoinInfo          map[string]CoinInfo `json:"coin_info"`
}
