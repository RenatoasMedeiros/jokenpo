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
	r.Use(withCORS)
	r.HandleFunc("/auth/register", handlers.CreatePlayerAccount).Methods("POST", "OPTIONS")
	r.HandleFunc("/auth/login", handlers.LoginPlayer).Methods("POST", "OPTIONS")
	// r.HandleFunc("/endgame/{id}", handlers.EndGameHandler).Methods("POST")
	r.Use(middlewares.Authenticate)
	r.HandleFunc("/room", handlers.CreateRoomHandler).Methods("POST", "OPTIONS")
	r.HandleFunc("/join/{room_id}", handlers.JoinRoomHandler).Methods("GET", "OPTIONS")
	// r.HandleFunc("/finishGame", EndGameHandler)

	srv := &http.Server{
		Addr:    api.config.addr,
		Handler: r,
	}
	fmt.Println("Server Started on Port: " + srv.Addr)
	return srv.ListenAndServe()
}

func withCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}
		next.ServeHTTP(w, r)
	})
}
