package entities

type ConversationByUsers struct {
	Messages []Message
	User     string
}

type Message struct {
	Text       string
	IsFromUser bool
	Time       string
}
