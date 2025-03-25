package handlers

import (
	"database/sql"
	"encoding/json"
	"jokenpo/backend/models"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type PlayerDB struct {
	db *sql.DB
}

var JWT_KET = []byte(os.Getenv("JWT_KEY"))

func (p PlayerDB) CreatePlayerAccount(w http.ResponseWriter, r *http.Request) {
	var player models.Player

	if err := json.NewDecoder(r.Body).Decode(&player); err != nil {
		http.Error(w, "Invalid request body, error: "+err.Error(), http.StatusInternalServerError)
		return
	}
	// create the user:
	if player.ID == uuid.Nil {
		player.ID = uuid.New()
	}
	query := `
		INSERT INTO players (id, username, password) VALUES ($1, $2, $3, $4)
		RETURNING username
	`

	err := p.db.QueryRowContext(r.Context(), query, player.ID, player.Username, player.Password).Scan(&player.Username)
	if err != nil {
		http.Error(w, "Error on the database, error: "+err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(player.Username)
}

func (p PlayerDB) LoginUser(w http.ResponseWriter, r *http.Request) {
	var player models.Player
	if err := json.NewDecoder(r.Body).Decode(&player); err != nil {
		http.Error(w, "Invalid request body, error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	query := `SELECT id, username, password FROM players WHERE id = $1`

	var storedPlayer models.Player
	err := p.db.QueryRow(query, player.ID).Scan(&storedPlayer.ID, &storedPlayer.Username, &storedPlayer.Password)

	if err != nil || bcrypt.CompareHashAndPassword([]byte(storedPlayer.Password), []byte(player.Password)) != nil {
		http.Error(w, "Invalid Credentials", http.StatusUnauthorized)
		return
	}

	// Create a JWT token with 24h of duration!
	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &jwt.StandardClaims{
		ExpiresAt: expirationTime.Unix(),
		Subject:   storedPlayer.ID.String(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign the token with the secret key
	tokenString, err := token.SignedString(JWT_KET)
	if err != nil {
		http.Error(w, "Could not create the token", http.StatusInternalServerError)
		return
	}

	response := struct {
		Token string `json:"token"`
	}{
		Token: tokenString,
	}

	// Marshal the struct into JSON
	jsonData, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}
