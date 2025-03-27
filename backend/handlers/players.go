package handlers

import (
	"encoding/json"
	"fmt"
	"jokenpo/backend/models"
	"jokenpo/database"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

var JWT_KET = []byte(os.Getenv("JWT_KEY"))

func CreatePlayerAccount(w http.ResponseWriter, r *http.Request) {
	var player models.Player
	if err := json.NewDecoder(r.Body).Decode(&player); err != nil {
		http.Error(w, "Invalid request body, error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Println("Player collected", player)

	db := database.GetDB()

	// create the user:
	if player.ID == uuid.Nil {
		player.ID = uuid.New()
	}
	fmt.Println("Final player before go to the database", player)
	query := `
		INSERT INTO players (id, username, password) VALUES ($1, $2, $3) RETURNING username
	`
	fmt.Println("Arrived here")

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(player.Password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, "Error on database, unable to proper validate the password"+err.Error(), http.StatusInternalServerError)
	}

	err = db.QueryRowContext(r.Context(), query, player.ID, player.Username, hashedPassword).Scan(&player.Username)
	if err != nil {
		http.Error(w, "Error on the database, error: "+err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Println("And here Arrived here")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(player.Username)
}

func LoginPlayer(w http.ResponseWriter, r *http.Request) {
	var player models.Player
	if err := json.NewDecoder(r.Body).Decode(&player); err != nil {
		http.Error(w, "Invalid request body, error: "+err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Println("Player received", player)
	db := database.GetDB()

	query := `SELECT id, username, password FROM players WHERE username = $1`

	var storedPlayer models.Player
	err := db.QueryRow(query, player.Username).Scan(&storedPlayer.ID, &storedPlayer.Username, &storedPlayer.Password)
	fmt.Println("StoredPlayer received", storedPlayer)

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
