package handlers

import (
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

func CreateRoomHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusCreated)
	//Generate UUID for the room
	gameID := uuid.NewString()
	//Comunicate with the web socket to create a room with that ID
	fmt.Println(w, "Game Created here is the id: %s ", gameID)
}

func JoinRoomHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	fmt.Println(w, "Player %s joined the room %s", vars["username"], vars["gameId"])
}
