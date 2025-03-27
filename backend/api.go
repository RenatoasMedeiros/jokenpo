package main

import (
	"fmt"
	handlers "jokenpo/backend/handlers"
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
	r.HandleFunc("/room", handlers.CreateRoomHandler)
	r.HandleFunc("/join/{id}", handlers.JoinRoomHandler)
	r.HandleFunc("/auth/login", handlers.LoginPlayer)
	r.HandleFunc("/auth/register", handlers.CreatePlayerAccount)
	// r.HandleFunc("/finishGame", EndGameHandler)

	srv := &http.Server{
		Addr:    api.config.addr,
		Handler: r,
	}
	fmt.Println("Server Started on Port: " + srv.Addr)
	return srv.ListenAndServe()
}
