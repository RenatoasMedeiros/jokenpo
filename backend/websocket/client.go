package websocket

import "github.com/gorilla/websocket"

type Client struct {
	/*We use the pointer here because we dont want to copy the connection, just use it*/
	Conn *websocket.Conn
	/*Here is the same logic that the webscoket.Conn room is a structure, and in this scenario we don't want to */
	Room *Room
	//Channel is already kind of reference..
	Send chan Message // Channel to send messages to this client
}
