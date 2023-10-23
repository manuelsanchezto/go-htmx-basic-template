package domain

type InterfaceResponseFull struct {
	User          string
	Conversations []Conversation
}

type Conversation struct {
	Question   Message
	Answer     Message
	IsAnswered bool
	Actions    string
	IsLast     bool
}

type Message struct {
	Text string
	Time string
}

func NewMessage(text string, time string) Message {
	return Message{
		Text: text,
		Time: time,
	}
}

func NewConversation(question Message, answer Message, answered bool, actions string, last bool) Conversation {
	return Conversation{
		Question:   question,
		Answer:     answer,
		IsAnswered: answered,
		Actions:    actions,
		IsLast:     last,
	}
}
