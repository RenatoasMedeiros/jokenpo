package websocket

import (
	"fmt"
	"net/http"
	"os"

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

		if origin == os.Getenv("FRONTEND_URL") ||
			origin == "http://"+os.Getenv("BACKEND_ADDR")+os.Getenv("BACKEND_PORT") { //TODO CHANGE TO HTTPS FOR PRODUCTION
			return true
		}
		return false
	},
}
