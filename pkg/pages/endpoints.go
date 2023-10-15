package pages

import (
	"github.com/labstack/echo/v4"
)

type IndexPage struct {
	Messages []Message
}

type Message struct {
	Text       string
	IsFromUser bool
	Time       string
}

func Index(c echo.Context) error {
	Messages := getMessages()
	return c.Render(200, "index.html", IndexPage{
		Messages: Messages,
	})
}

func AddMessage(c echo.Context) error {
	var Messages = appendnews(getMessages(), Message{"Hello", true, "12:00"})
	return c.Render(200, "index.html", IndexPage{
		Messages: Messages,
	})
}

func getMessages() []Message {
	var messages []Message
	return messages
}

func appendnews(messages []Message, newMessage Message) []Message {
	//generate a new Message array that copys Message and appends the newMessage appended
	newMessages := append(messages, newMessage)
	return newMessages

}
