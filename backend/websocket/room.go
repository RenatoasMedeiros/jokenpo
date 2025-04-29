package websocket

import (
	"fmt"

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
