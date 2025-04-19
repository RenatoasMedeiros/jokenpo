package handlers

import (
	"encoding/json"
	"fmt"
	"jokenpo/backend/models"
	"jokenpo/backend/websocket"
	"jokenpo/database"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

func CreateRoomHandler(w http.ResponseWriter, r *http.Request) {
	//Generate UUID for the room
	gameID := uuid.NewString()

	room := &websocket.Room{
		ID: gameID,
		/*
			Maps need to be initialized before used, thats why we need the make
			The key, is a pointer to a Client (single player connection, the bool represents if they are (or not) in the room
		*/
		Clients: make(map[*websocket.Client]bool),
		//We use channels to protect from race conditions!
		Broadcast:  make(chan websocket.Message),
		Register:   make(chan *websocket.Client),
		Unregister: make(chan *websocket.Client),
	}
	//Go routine
	go room.Run()

	//this Rooms is a global map to keep track of all active rooms
	websocket.Rooms[gameID] = room

	//Comunicate with the web socket to create a room with that ID
	fmt.Println(w, "Game Created here is the id: %s ", gameID)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(gameID)
}

func JoinRoomHandler(w http.ResponseWriter, r *http.Request) {
	gameID := mux.Vars(r)["room_id"]

	//Now is needed to upgrade from HTTP -> WS
	conn, err := websocket.Upgrader.Upgrade(w, r, nil)
	if err != nil {
		http.Error(w, "Could not open websocket connection", http.StatusBadRequest)
		return
	}
	fmt.Println("Arrived here1")
	room, ok := websocket.Rooms[gameID]
	if !ok {
		http.Error(w, "Room not found", http.StatusNotFound)
		return
	}

	//Create the client:
	client := &websocket.Client{
		Conn: conn,
		Room: room,
		Send: make(chan websocket.Message),
	}

	//Actually register the client on the room
	room.Register <- client
	fmt.Println("Arrived here2")
	go client.ReadPump()
	go client.WritePump()

	fmt.Printf("Client connected to room %s\n", gameID)
}

func EndGameHandler(w http.ResponseWriter, r *http.Request) {
	//mux.Vars goes to the variables on the request url not on the x
	vars := mux.Vars(r) //this vars will only have the id

	var input models.GameInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid request body: "+err.Error(), http.StatusBadRequest)
		return
	}

	player1UUID, err := uuid.Parse(input.Player1)
	if err != nil {
		http.Error(w, "Invalid player_1 UUID: "+err.Error(), http.StatusBadRequest)
		return
	}

	player2UUID, err := uuid.Parse(input.Player2)
	if err != nil {
		http.Error(w, "Invalid player_2 UUID: "+err.Error(), http.StatusBadRequest)
		return
	}

	game := models.Games{
		Player1:       player1UUID,
		Player2:       player2UUID,
		Player1Choice: input.Player1Choice,
		Player2Choice: input.Player2Choice,
	}

	db := database.GetDB()

	idGame, err := uuid.Parse(vars["id"])
	if err != nil {
		http.Error(w, "Invalid id, error: "+err.Error(), http.StatusBadRequest)
	}

	switch game.Player1Choice {
	case "rock":
		if game.Player2Choice == "scissors" {
			game.Winner = game.Player1
		} else {
			game.Winner = game.Player2
		}
	case "paper":
		if game.Player2Choice == "rock" {
			game.Winner = game.Player1
		} else {
			game.Winner = game.Player2
		}
	case "scissors":
		if game.Player2Choice == "paper" {
			game.Winner = game.Player1
		} else {
			game.Winner = game.Player2
		}
	}

	query := `
		INSERT INTO games (id, player_1, player_2, player_1_choice, player_2_choice, winner) VALUES ($1, $2, $3, $4, $5, $6) RETURNING winner
	`
	err = db.QueryRowContext(r.Context(), query, idGame, game.Player1, game.Player2, game.Player1Choice, game.Player2Choice, game.Winner).Scan(&game.Winner)
	if err != nil {
		http.Error(w, "Error on the database, error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(game.Winner)
}
