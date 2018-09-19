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

// Bet blabla
type Bet struct {
	Value float32 `json:"value"`
	Horse string  `json:"horse"`
	From  string  `json:"from"`
}

// odd blabla
type Odd struct {
	Value float32 `json:"value"`
	Horse string  `json:"horse"`
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
	Version         string     `json:"version"`
	WinnerHorses    []string   `json:"winner_horses"`
	Odds            []Odd      `json:"odds"`
	Bets            []Bet      `json:"bets"`
	Withdraws       []Withdraw `json:"withdraws"`
	Volume          float32    `json:"volume"`
	Refunded        bool       `json:"refunded"`
	Active          string     `json:"active"`
	Complete        bool       `json:"complete"`
}

// RaceDataLight blabla
type RaceDataLight struct {
	ContractID      string   `json:"contractid"`
	Date            uint64   `json:"date"`
	RaceDuration    uint64   `json:"race_duration"`
	BettingDuration uint64   `json:"betting_duration"`
	EndTime         uint64   `json:"end_time"`
	RaceNumber      uint32   `json:"race_number"`
	Version         string   `json:"version"`
	WinnerHorses    []string `json:"winner_horses"`
	Odds            []Odd    `json:"odds"`
	Volume          float32  `json:"volume"`
	Refunded        bool     `json:"refunded"`
	Active          string   `json:"active"`
	Complete        bool     `json:"complete"`
}

// Cache blabla
type Cache struct {
	List       []RaceData `json:"list"`
	LastUpdate int64      `json:"last_update"`
}

// CacheLight blabla
type CacheLight struct {
	List       []RaceDataLight `json:"list"`
	LastUpdate int64           `json:"last_update"`
}

// Achievement blablabla
type Achievement struct {
	Label string `json:"label"`
}

// Horsey blablabla
type Horsey struct {
	RaceNumber  uint32 `json:"race_number"`
	RaceAddress string `json:"race_address"`
	Symbol      string `json:"symbol"`
}

// User blablabla
type User struct {
	Address      string        `json:"address"`
	GamesCount   uint32        `json:"games_count"`
	RankCash     uint32        `json:"rank_cash"`
	RankWinLoss  uint32        `json:"rank_winloss"`
	RankWinner   uint32        `json:"rank_winner"`
	WinsCount    uint32        `json:"wins_count"`
	LossesCount  uint32        `json:"losses_count"`
	Benefit      float32       `json:"benefit"`
	Achievements []Achievement `json:"achievements"`
	Horseys      []Horsey      `json:"horseys"`
}

// Rank blablabla
type Rank struct {
	Address     string  `json:"address"`
	Benefit     float32 `json:"benefit"`
	RankCash    uint32  `json:"rank_cash"`
	RankWinLoss uint32  `json:"rank_winloss"`
	RankWinner  uint32  `json:"rank_winner"`
	WinsCount   uint32  `json:"wins_count"`
	LossesCount uint32  `json:"losses_count"`
	GamesCount  uint32  `json:"games_count"`
}
