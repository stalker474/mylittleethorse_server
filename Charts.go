package main

import (
	"fmt"
	"strings"
	"time"
)

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

// GetCharts blabla
func GetCharts(server *Server, periodBegin uint64, periodEnd uint64) (Stats, error) {
	server.data.mux.Lock()
	stats := Stats{}
	stats.CoinInfo = make(map[string]CoinInfo)
	stats.CoinInfo["BTC"] = CoinInfo{WinsCount: 0}
	stats.CoinInfo["ETH"] = CoinInfo{WinsCount: 0}
	stats.CoinInfo["LTC"] = CoinInfo{WinsCount: 0}

	playerCount := make(map[string]uint32)
	dayPlayerCount := make(map[string]map[string]uint32)
	days := make(map[string]Day)
	stats.PeriodBegin = periodBegin
	stats.PeriodEnd = periodEnd
	seedingAddr := "0x1F92771237Bd5eae04e91B4B6F1d1a78D41565a2"
	for _, race := range server.data.racesData {
		//count only non refunded closed races
		if race.Date >= periodBegin && race.Date <= periodEnd {
			if !race.Refunded && (strings.Compare(race.Active, "Closed") == 0) {
				//local to race data
				coinVolume := make(map[string]float32)
				coinBetsCount := make(map[string]uint32)

				tm := time.Unix(int64(race.Date), 0)
				formattedDay := fmt.Sprintf("%d-%02d-%02d", tm.Year(), tm.Month(), tm.Day())
				day, exists := days[formattedDay]
				if !exists {
					day = Day{}
					day.Coins = make(map[string]Coin)
				}
				day.Label = formattedDay
				stats.TotalRaces++
				stats.TotalVolume += race.Volume

				for _, horse := range race.WinnerHorses {
					stats.CoinInfo[horse] = CoinInfo{WinsCount: stats.CoinInfo[horse].WinsCount}
				}
				userVolume := float32(0)
				for _, bet := range race.Bets {
					playerCount[bet.From]++
					dayPlayerCount[formattedDay][bet.From]++
					coinVolume[bet.Horse] += bet.Value
					coinBetsCount[bet.Horse]++
					if strings.Compare(bet.From, seedingAddr) != 0 {
						userVolume += bet.Value
					}
				}

				day.Volume += race.Volume
				day.UserVolume += userVolume

				for coin, value := range coinVolume {
					c, _ := day.Coins[coin]
					c.Volume += value
					c.BetsCount += coinBetsCount[coin]
					day.Coins[coin] = c
				}
			}
		}
	}
	for _, day := range days {
		stats.Days = append(stats.Days, day)
	}
	stats.TotalPlayersCount = uint32(len(playerCount))
	for formattedDay, coins := range dayPlayerCount {
		day, _ := days[formattedDay]
		day.PlayersCount = uint32(len(coins))
		days[formattedDay] = day
	}
	server.data.mux.Unlock()

	return stats, nil
}
