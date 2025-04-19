package websocket

import (
	"fmt"

	"github.com/gorilla/websocket"
)

type Client struct {
	/*We use the pointer here because we dont want to copy the connection, just use it*/
	Conn *websocket.Conn
	/*Here is the same logic that the webscoket.Conn room is a structure, and in this scenario we don't want to */
	Room *Room
	//Channel is already kind of reference..
	Send chan Message // Channel to send messages to this client
}

// constantly listens if the client sends something.
func (c *Client) ReadPump() {
	defer func() {
		c.Room.Unregister <- c
		c.Conn.Close()
	}()

	for {
		var msg Message
		err := c.Conn.ReadJSON(&msg)
		if err != nil {
			break
		}
		fmt.Println("Arrived here3")
		// When the player sends the move
		fmt.Println("msg.type", msg.Type)
		if msg.Type == "move" || msg.Type == "join" || msg.Type == "result" {
			fmt.Println("Arrived here4")
			c.Room.Broadcast <- msg
		}
	}
}

// constantly listens if the server wants to send something to the client.
func (c *Client) WritePump() {
	defer c.Conn.Close()
	for {
		select {
		case msg, ok := <-c.Send:
			if !ok {
				// Channel closed
				c.Conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}
			c.Conn.WriteJSON(msg)
		}
	}
}
