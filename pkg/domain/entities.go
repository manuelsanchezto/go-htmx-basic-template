package entities

type InterfaceResponseFull struct {
	Messages      []Message
	User          string
	FullResponses bool
}

type Message struct {
	Text       string
	IsFromUser bool
	Time       string
	IsLast     bool
}
