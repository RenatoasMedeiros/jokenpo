package websocket

import "fmt"

// i have copied this structure and the run() from https://github.com/gorilla/websocket/blob/main/examples/chat/hub.go
type Room struct {
	ID         string
	Clients    map[*Client]bool
	Broadcast  chan Message
	Register   chan *Client
	Unregister chan *Client
}

func (r *Room) Run() {
	numClients := 0
	for {
		//select here since we have multiple options to interact with the Room
		select {
		case client := <-r.Register:
			if numClients < 2 {
				r.Clients[client] = true
				numClients++
				fmt.Println("Register Client:", client)
				fmt.Println("Total of Clients", numClients)
			}

		case client := <-r.Unregister:
			if _, ok := r.Clients[client]; ok {
				delete(r.Clients, client)
				close(client.Send)
				fmt.Println("UnRegister Client:", client)
			}
		case message := <-r.Broadcast:
			fmt.Printf("Broadcast Message %+v\n:", message)
			for client := range r.Clients {
				select {
				case client.Send <- message:
				default:
					close(client.Send)
					delete(r.Clients, client)
					fmt.Println("Client removed due to unresponsiveness.")

				}
			}
		}
	}
}
