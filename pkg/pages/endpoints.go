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

func getMessages() []Message {
	var messages []Message
	return messages
}
