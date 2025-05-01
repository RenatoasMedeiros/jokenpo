package main

import (
	"fmt"
	"net/http"
	"os"
	"statistics/handlers"

	"github.com/gorilla/mux"
)

type api struct {
	config config
}

type config struct {
	addr string
	port string
}

func (api *api) run() error {
	r := mux.NewRouter()
	r.Use(withCORS)
	r.HandleFunc("/statistics/ranking", handlers.GetStatisticsFromDB).Methods("GET", "OPTIONS")

	srv := &http.Server{
		Addr:    api.config.addr + api.config.port,
		Handler: r,
	}
	fmt.Println("Server Statistics Started on Port: " + srv.Addr)
	return srv.ListenAndServe()
}

func withCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", os.Getenv("FRONTEND_URL"))
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}
		next.ServeHTTP(w, r)
	})
}
