package websocket

import "github.com/gorilla/websocket"

type Client struct {
	Conn *websocket.Conn
	Room *Room
	Send chan Message // Channel to send messages to this client
}
