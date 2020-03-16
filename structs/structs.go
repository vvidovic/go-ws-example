package structs

type ID struct {
	ID int `json:"id"`
}

type Message struct {
	ID      int    `json:"id,omitempty"`
	Sender  string `json:"sender"`
	Message string `json:"message"`
}

type MessageList []Message
