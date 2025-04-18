package websocket

type Message struct {
	Type   string `json:"type"`   // e.g. "move", "join", "result"
	Body   string `json:"body"`   // JSON string (can parse separately)
	Sender string `json:"sender"` // (optional) username or player id
}
