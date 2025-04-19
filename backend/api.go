package main

import (
	"fmt"
	handlers "jokenpo/backend/handlers"
	middlewares "jokenpo/backend/middlewares"
	"net/http"

	"github.com/gorilla/mux"
)

type api struct {
	config config
}

type config struct {
	addr string
}

func (api *api) run() error {
	r := mux.NewRouter()
	r.HandleFunc("/room", handlers.CreateRoomHandler).Methods("POST")
	r.HandleFunc("/join/{room_id}", handlers.JoinRoomHandler).Methods("GET")
	r.HandleFunc("/endgame/{id}", handlers.EndGameHandler).Methods("POST")
	r.HandleFunc("/auth/login", handlers.LoginPlayer).Methods("POST")
	r.HandleFunc("/auth/register", handlers.CreatePlayerAccount).Methods("POST")
	r.Use(middlewares.Authenticate)
	// r.HandleFunc("/finishGame", EndGameHandler)

	srv := &http.Server{
		Addr:    api.config.addr,
		Handler: r,
	}
	fmt.Println("Server Started on Port: " + srv.Addr)
	return srv.ListenAndServe()
}
