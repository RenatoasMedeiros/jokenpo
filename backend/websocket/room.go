package websocket

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
		case Client := <-r.Register:
			if numClients < 2 {
				r.Clients[Client] = true
				numClients++
			}

		case Client := <-r.Unregister:
			if _, ok := r.Clients[Client]; ok {
				delete(r.Clients, Client)
				close(Client.Send)
			}
			//Still need to think up the logic for the message!
			// case Message := <-r.Broadcast:
			// 	for Client := range r.Clients {
			// 		select {
			// 		case Client.send <- Message:
			// 		default:
			// 			close(Client.send)
			// 			delete(r.Client, Client)
			// 		}
			// 	}
			// }
		}
	}
}
