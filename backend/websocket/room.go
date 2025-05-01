package websocket

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

// i have copied this structure and the run() from https://github.com/gorilla/websocket/blob/main/examples/chat/hub.go
type Room struct {
	ID         uuid.UUID
	Clients    map[*Client]bool
	Broadcast  chan Message
	Register   chan *Client
	Unregister chan *Client
	Moves      map[uuid.UUID]string //map[clientid]move
}

func (r *Room) Run() {
	numClients := 0
	fmt.Println("Inside the Run Function, numClients ", numClients)
	for {
		fmt.Println("Inside the FOR Run Function")
		//select here since we have multiple options to interact with the Room
		select {
		case Client := <-r.Register:
			if numClients < 2 {
				fmt.Println("Inside the RUN Function - Select r.Register, client: ", Client)
				r.Clients[Client] = true
				numClients++
				fmt.Println("Register Client:", Client)
				fmt.Println("Total of Clients", numClients)

				if len(r.Clients) == 2 {
					fmt.Println("Room has 2 players, broadcasting start")

					startMsg := Message{
						Type: "start",
						Body: "",
					}
					r.Broadcast <- startMsg

					//On the start of the game set the countdown
					go func(r *Room) {
						// Send countdown
						for i := 10; i > 0; i-- {
							r.Broadcast <- Message{
								Type: "countdown",
								Body: fmt.Sprintf("%ds", i),
							}
							time.Sleep(1 * time.Second)
						}

						r.Broadcast <- Message{
							Type: "countdown",
							Body: "Time's up!",
						}

						// Decide outcome based on how many moves were made
						moveCount := len(r.Moves)

						// Default values
						var p1ID, p2ID uuid.UUID
						var p1Move, p2Move string
						var winner uuid.UUID

						i := 0
						for id, move := range r.Moves {
							if i == 0 {
								p1ID = id
								p1Move = move
							} else {
								p2ID = id
								p2Move = move
							}
							i++
						}

						var resultBody string

						if moveCount == 0 {
							resultBody = "draw (no one played)"
						} else if moveCount == 1 {
							resultBody = fmt.Sprintf("%s wins (opponent didn't play)", p1ID)
							winner = p1ID
						} else if moveCount == 2 {
							winnerLabel, winID := DetermineWinner(p1Move, p1ID, p2Move, p2ID)
							if winnerLabel == "draw" {
								resultBody = "draw"
							} else {
								resultBody = fmt.Sprintf("winner: %s", winID)
								winner = winID
							}
						}

						// Broadcast gameover message
						r.Broadcast <- Message{
							Type: "gameover",
							Body: resultBody,
						}

						// Save game to DB even in edge cases
						SaveGameResultToDB(p1ID, p2ID, p1Move, p2Move, winner)

						// Close room
						for client := range r.Clients {
							r.Unregister <- client
							client.Conn.Close()
						}
						delete(Rooms, r.ID)
						fmt.Println("Room closed and deleted:", r.ID)
					}(r)

				}

			}

		case Client := <-r.Unregister:
			fmt.Println("Inside the RUN Function - Select r.UnRegister, Client: ", Client)
			if _, ok := r.Clients[Client]; ok {
				delete(r.Clients, Client)
				close(Client.Send)
				fmt.Println("UnRegister Client:", Client)
			}
		case message := <-r.Broadcast:
			fmt.Printf("Broadcast Message %+v\n:", message)
			for Client := range r.Clients {
				select {
				case Client.Send <- message:
				default:
					close(Client.Send)
					delete(r.Clients, Client)
					fmt.Println("Client removed due to unresponsiveness.")

				}
			}
		}
	}
}
