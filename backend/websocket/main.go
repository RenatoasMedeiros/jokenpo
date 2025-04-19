package websocket

import (
	"net/http"

	"github.com/gorilla/websocket"
)

// this rooms are global because it must be visible by both backend handlers and websocket code
var Rooms = make(map[string]*Room)

var Upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		// origin := r.Header.Get("Origin")
		// return origin == "http://localhost:3000" || origin == "http://localhost:8080"
		return true
	},
}
