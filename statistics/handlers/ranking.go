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
	PlayerID    string `json:"player_id"`
	Username    string `json:"username"`
	Wins        int    `json:"wins"`
	GamesPlayed int    `json:"games_played"`
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
        WITH PlayerGameCounts AS (
            SELECT
                p.id,
                p.username,
                COUNT(g.id) AS games_played -- Count total games participated in
            FROM players p
            LEFT JOIN games g ON p.id = g.player_1 OR p.id = g.player_2
            GROUP BY p.id, p.username
        ), PlayerWinCounts AS (
            SELECT
                p.id,
                COUNT(g.winner) as wins -- Count games won
            FROM players p
            LEFT JOIN games g ON g.winner = p.id
            GROUP BY p.id
        )
        SELECT
            pgc.id,
            pgc.username,
            COALESCE(pwc.wins, 0) as wins, -- Use COALESCE for players with 0 wins
            pgc.games_played
        FROM PlayerGameCounts pgc
        LEFT JOIN PlayerWinCounts pwc ON pgc.id = pwc.id
        ORDER BY wins DESC, games_played DESC -- Define order (e.g., wins first, then games played)
        LIMIT 20;
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
		if err := rows.Scan(&ps.PlayerID, &ps.Username, &ps.Wins, &ps.GamesPlayed); err != nil {
			log.Println("Scan error (top players):", err)
			continue
		}
		topPlayers = append(topPlayers, ps)
	}
	if err := rows.Err(); err != nil {
		log.Println("Error iterating top players rows:", err)
	}
	fmt.Println("statistics - topPlayers: ", topPlayers)

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
	if totalMoves > 0 { // Prevent division by zero!
		for move, count := range moveCounts {
			movePercent[move] = (float64(count) / float64(totalMoves)) * 100
		}
	}
	fmt.Println("Move Percent: ", movePercent)
	// Combined Statistics
	stats := StatisticsResponse{
		TopPlayers:     topPlayers,
		MoveUsageStats: movePercent,
	}
	fmt.Println("statistics - stats: ", stats)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(stats)
}
