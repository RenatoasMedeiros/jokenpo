package websocket

import (
	"encoding/json"
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

type MovePayload struct {
	GameID string `json:"game_id"`
	Move   string `json:"move"`
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
		fmt.Println("WHOLE MESSAGE:", msg)
		// if msg.Type == "move" || msg.Type == "join" || msg.Type == "result" {
		// 	c.Room.Broadcast <- msg
		// }
		switch msg.Type {
		case "move":
			// Save the move for this player
			fmt.Println("MOVE DETECTED - Client", c)
			fmt.Println("msg.Body", msg.Body)

			var movePayload MovePayload
			err := json.Unmarshal([]byte(msg.Body), &movePayload)
			if err != nil {
				fmt.Println("Error parsing move payload:", err)
				return
			}

			// Now you can safely use:
			fmt.Println("Move:", movePayload.Move)
			fmt.Println("Game ID:", movePayload.GameID)
			fmt.Println("Sender:", msg.Sender)
			c.Room.Moves[uuid.MustParse(msg.Sender)] = movePayload.Move
			fmt.Println("MOVE DETECTED - c.Room.Moves", c.Room.Moves)

			// var player1ID, player2ID uuid.UUID
			// i := 0
			// // Check if both players moved
			// fmt.Println("len c.Room.Moves", len(c.Room.Moves))
			// if len(c.Room.Moves) == 2 {
			// 	playerChoices := make([]string, 0, 2)
			// 	for id, choice := range c.Room.Moves {
			// 		playerChoices = append(playerChoices, choice)
			// 		fmt.Println("Choices", choice)
			// 		if i == 0 {
			// 			player1ID = id
			// 			fmt.Println("Player1ID", player1ID)
			// 		} else {
			// 			fmt.Println("player2ID", player2ID)
			// 			player2ID = id
			// 		}
			// 		fmt.Println("DENTRO NO LOOP I", i)
			// 		i++
			// 	}

			// 	// Calculate winner
			// 	winner, winnerId := DetermineWinner(playerChoices[0], player1ID, playerChoices[1], player2ID)
			// 	fmt.Println("WINNER", winner)

			// 	if winner != "draw" {
			// 		// Broadcast winner
			// 		resultMsg := Message{
			// 			Type: "result",
			// 			Body: winnerId.String(),
			// 		}
			// 		c.Room.Broadcast <- resultMsg
			// 	} else {
			// 		// Broadcast winner
			// 		resultMsg := Message{
			// 			Type: "result",
			// 			Body: winner,
			// 		}
			// 		c.Room.Broadcast <- resultMsg
			// 	}

			// 	saveGameResultToDB(player1ID, player2ID, *c.Room)
			// }
		}
	}
}

// constantly listens if the server wants to send something to the client.
func (c *Client) WritePump() {
	fmt.Println("WritePump started for client", c.ID)
	defer func() {
		fmt.Println("Closing WritePump for client", c.ID)
		c.Conn.Close()
	}()
	for {
		select {
		case msg, ok := <-c.Send:
			if !ok {
				fmt.Println("Channel closed for client", c.ID)
				c.Conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}
			fmt.Println("Sending message to client", c.ID, ":", msg)
			err := c.Conn.WriteJSON(msg)
			if err != nil {
				fmt.Println("Write error:", err)
				return
			}
		}
	}
}

func DetermineWinner(player1Move string, player1ID uuid.UUID, player2Move string, player2ID uuid.UUID) (string, uuid.UUID) {
	if player1Move == player2Move {
		return "draw", uuid.MustParse("ffffffff-ffff-ffff-ffff-ffffffffffff")
	}

	switch player1Move {
	case "rock":
		if player2Move == "scissors" {
			return "player1", player1ID
		}
	case "paper":
		if player2Move == "rock" {
			return "player1", player1ID
		}
	case "scissors":
		if player2Move == "paper" {
			return "player1", player1ID
		}
	}
	return "player2", player2ID
}

func SaveGameResultToDB(player1, player2 uuid.UUID, move1, move2 string, winner uuid.UUID) {
	game := models.Games{
		Player1:       player1,
		Player2:       player2,
		Player1Choice: move1,
		Player2Choice: move2,
		Winner:        winner,
	}

	db := database.GetDB()
	query := `
		INSERT INTO games (id, player_1, player_2, player_1_choice, player_2_choice, winner)
		VALUES ($1, $2, $3, $4, $5, $6)
	`
	_, err := db.Exec(query, uuid.New(), game.Player1, game.Player2, game.Player1Choice, game.Player2Choice, game.Winner)
	if err != nil {
		fmt.Println("Error saving game result:", err)
	}
}
