package pages

import (
	"time"

	"github.com/labstack/echo/v4"
	dbconn "htmx.try/m/v2/pkg/dbconn"
	entities "htmx.try/m/v2/pkg/domain"
)

var conn = *dbconn.NewInMemoryDB()

func Index(c echo.Context) error {
	user := c.QueryParam("user")
	Messages := getMessages(user)
	isFullResponses := getIfFullResponses(user)

	return c.Render(200, "index.html", entities.InterfaceResponseFull{
		Messages:      Messages,
		User:          user,
		FullResponses: isFullResponses,
	})
}

func AddMessage(c echo.Context) error {
	user := c.FormValue("user")
	text := c.FormValue("messagetext")
	var Messages = appendnews(user, entities.Message{text, true, time.Now().Format("15:04:05"), true})
	Messages = appendnews(user, entities.Message{"", false, time.Now().Format("15:04:05"), true})
	var cosas = getFullConn(user)
	cosas.FullResponses = false
	conn.Set(user, cosas)
	//TODO: Desarrollar esta rutina
	go generateMessage(user)

	return c.Render(200, "index.html", entities.InterfaceResponseFull{
		Messages:      Messages,
		User:          user,
		FullResponses: false,
	})
}
func RecoverMessages(c echo.Context) error {
	user := c.FormValue("user")
	var Messages = getMessages(user)
	for _, val := range Messages {
		if val.IsFromUser == false && val.Text == "" {
			return c.Render(200, "index.html", entities.InterfaceResponseFull{
				Messages:      Messages,
				User:          user,
				FullResponses: false,
			})
		}
	}
	c.Response().Before(func() { c.Response().Header().Add("HX-Trigger", "done") })
	return c.Render(200, "index.html", entities.InterfaceResponseFull{
		Messages:      Messages,
		User:          user,
		FullResponses: true,
	})
}

func getMessages(user string) []entities.Message {
	//get the messages from the database
	val, ok := conn.Get(user)
	if !ok {
		return []entities.Message{}
	}
	return val.Messages
}

func getFullConn(user string) entities.InterfaceResponseFull {
	//get the messages from the database
	val, ok := conn.Get(user)
	if !ok {
		return entities.InterfaceResponseFull{}
	}
	return val
}

func appendnews(user string, newMessage entities.Message) []entities.Message {
	var getMessages = getMessages(user)
	var lastElement = len(getMessages)
	if len(getMessages)%2 != 0 {
		lastElement = len(getMessages) - 1
	}
	for i := 0; i < lastElement; i++ {
		getMessages[i].IsLast = false
	}
	newMessages := append(getMessages, newMessage)
	var conversation = getFullConn(user)
	conversation.Messages = newMessages
	conn.Set(user, conversation)

	return newMessages

}

func generateMessage(user string) {
	time.Sleep(1 * time.Second)
	var Messages = getMessages(user)
	for pos, val := range Messages {
		if val.IsFromUser == false && val.Text == "" {
			Messages[pos] = entities.Message{"Hello, how can I help you?", false, time.Now().Format("15:04:05"), true}
			var updatedConv = getFullConn(user)
			updatedConv.Messages = Messages
			updatedConv.FullResponses = true
			conn.Set(user, updatedConv)
			return
		}
	}
}

func RecoverResponses(c echo.Context) error {
	user := c.QueryParam("user")
	Messages := getMessages(user)
	for _, val := range Messages {
		if val.IsFromUser == false && val.Text == "" {
			return c.Render(200, "index.html", entities.InterfaceResponseFull{
				Messages:      Messages,
				User:          user,
				FullResponses: false,
			})
		}
	}
	c.Response().Before(func() { c.Response().Header().Add("HX-Trigger", "Done") })
	return c.Render(200, "index.html", entities.InterfaceResponseFull{
		Messages:      Messages,
		User:          user,
		FullResponses: true,
	})
}

func getIfFullResponses(user string) bool {
	//get the messages from the database
	val, ok := conn.Get(user)
	if !ok {
		return true
	}
	return val.FullResponses
}
