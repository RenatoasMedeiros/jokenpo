package backend

import (
	handlers "jokenpo/backend/handlers"
	"net/http"

	"github.com/gorilla/mux"
)

func routes() {
	r := mux.NewRouter()
	r.HandleFunc("/room", handlers.CreateRoomHandler)
	r.HandleFunc("/join/{id}", handlers.JoinRoomHandler)
	// r.HandleFunc("/finishGame", EndGameHandler)
	http.Handle("/", r)
}
