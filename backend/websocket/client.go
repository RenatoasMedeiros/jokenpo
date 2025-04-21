package websocket

import (
	"fmt"
	"jokenpo/backend/models"
	"jokenpo/database"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

type Client struct {
	ID uuid.UUID
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
		// When the player sends the move
		fmt.Println("msg.type", msg.Type)
		// if msg.Type == "move" || msg.Type == "join" || msg.Type == "result" {
		// 	c.Room.Broadcast <- msg
		// }
		switch msg.Type {
		case "move":
			// Save the move for this player
			c.Room.Moves[c.ID] = msg.Body

			var player1ID, player2ID uuid.UUID
			i := 0
			// Check if both players moved
			if len(c.Room.Moves) == 2 {
				playerChoices := make([]string, 0, 2)
				for id, choice := range c.Room.Moves {
					playerChoices = append(playerChoices, choice)
					if i == 0 {
						player1ID = id
					} else {
						player2ID = id
					}
					i++
				}

				// Calculate winner
				winner := determineWinner(playerChoices[0], playerChoices[1])

				// Broadcast winner
				resultMsg := Message{
					Type: "result",
					Body: winner, // "player1", "player2" or "draw"
				}
				c.Room.Broadcast <- resultMsg

				// Save into DB (optional)
				saveGameResultToDB(player1ID, player2ID, *c.Room)
			}
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

func determineWinner(player1Move string, player2Move string) string {
	switch player1Move {
	case "rock":
		if player2Move == "scissors" {
			return "player1"
		} else {
			return "player2"
		}
	case "paper":
		if player2Move == "rock" {
			return "player1"
		} else {
			return "player2"
		}
	case "scissors":
		if player2Move == "paper" {
			return "player1"
		} else {
			return "player2"
		}
	}
	return "draw"
}

func saveGameResultToDB(player1 uuid.UUID, player2 uuid.UUID, room Room) {
	game := models.Games{
		Player1:       player1,
		Player2:       player2,
		Player1Choice: room.Moves[player1],
		Player2Choice: room.Moves[player2],
	}

	db := database.GetDB()

	query := `
		INSERT INTO games (id, player_1, player_2, player_1_choice, player_2_choice, winner) VALUES ($1, $2, $3, $4, $5, $6)
	`
	_, err := db.Exec(query, room.ID, game.Player1, game.Player2, game.Player1Choice, game.Player2Choice, game.Winner)
	if err != nil {
		// http.Error(w, "Error on the database, error: "+err.Error(), http.StatusInternalServerError)
		fmt.Errorf("Error on the database, error: " + err.Error())
		return
	}
}
