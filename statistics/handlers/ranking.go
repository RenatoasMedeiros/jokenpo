package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"jokenpo/database"
	"log"
	"net/http"
)

type PlayerStats struct {
	PlayerID string `json:"player_id"`
	Username string `json:"username"`
	Wins     int    `json:"wins"`
}

type StatisticsResponse struct {
	TopPlayers     []PlayerStats      `json:"top_players"`
	MoveUsageStats map[string]float64 `json:"move_usage_percent"`
}

func GetStatisticsFromDB(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	db := database.GetDB()

	// --- 1. Top 20 Players by Wins ---
	topPlayersQuery := `
		SELECT p.id, p.username, COUNT(g.winner) as wins
		FROM games g
		JOIN players p ON g.winner = p.id
		GROUP BY p.id, p.username
		ORDER BY wins DESC
		LIMIT 20
	`
	rows, err := db.QueryContext(ctx, topPlayersQuery)
	if err != nil {
		log.Println("DB error (top players):", err)
		http.Error(w, "Failed to fetch top players", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var topPlayers []PlayerStats
	for rows.Next() {
		var ps PlayerStats
		if err := rows.Scan(&ps.PlayerID, &ps.Username, &ps.Wins); err != nil {
			log.Println("Scan error (top players):", err)
			continue
		}
		topPlayers = append(topPlayers, ps)
		fmt.Println("statistics - PlayerStats: ", ps)
	}

	// --- 2. Move Usage Stats ---
	moveUsageQuery := `
		SELECT move, COUNT(*) FROM (
			SELECT player_1_choice AS move FROM games
			UNION ALL
			SELECT player_2_choice AS move FROM games
		) AS all_moves
		GROUP BY move
	`
	moveRows, err := db.QueryContext(ctx, moveUsageQuery)
	if err != nil {
		log.Println("DB error (move usage):", err)
		http.Error(w, "Failed to fetch move usage stats", http.StatusInternalServerError)
		return
	}
	defer moveRows.Close()

	moveCounts := make(map[string]int)
	var totalMoves int
	for moveRows.Next() {
		var move string
		var count int
		if err := moveRows.Scan(&move, &count); err != nil {
			log.Println("Scan error (move usage):", err)
			continue
		}
		moveCounts[move] = count
		totalMoves += count
	}

	movePercent := make(map[string]float64)
	for move, count := range moveCounts {
		movePercent[move] = (float64(count) / float64(totalMoves)) * 100
	}

	// Combined Statistics
	stats := StatisticsResponse{
		TopPlayers:     topPlayers,
		MoveUsageStats: movePercent,
	}
	fmt.Println("statistics - stats: ", stats)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(stats)
}
