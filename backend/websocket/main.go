package websocket

import (
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

// this rooms are global because it must be visible by both backend handlers and websocket code
var Rooms = make(map[uuid.UUID]*Room)

var Upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		origin := r.Header.Get("Origin")
		fmt.Println("WebSocket handshake Origin:", origin) // DEBUG

		if origin == "http://localhost:5173" ||
			origin == "http://localhost:3000" ||
			origin == "http://localhost:8080" {
			return true
		}
		return false
	},
}
