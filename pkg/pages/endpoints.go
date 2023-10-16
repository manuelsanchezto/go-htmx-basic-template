package pages

import (
	"log"
	"time"

	"github.com/labstack/echo/v4"
	dbconn "htmx.try/m/v2/pkg/dbconn"
	entities "htmx.try/m/v2/pkg/domain"
)

var conn = *dbconn.NewInMemoryDB()

func Index(c echo.Context) error {
	user := c.QueryParam("user")
	Messages := getMessages(user)

	return c.Render(200, "index.html", entities.ConversationByUsers{
		Messages: Messages,
		User:     user,
	})
}

func AddMessage(c echo.Context) error {
	user := c.FormValue("user")
	log.Println("usuario: ", user)
	text := c.FormValue("messagetext")
	var Messages = appendnews(user, entities.Message{text, true, time.Now().Format("15:04:05")})
	return c.Render(200, "index.html", entities.ConversationByUsers{
		Messages: Messages,
		User:     user,
	})
}

func getMessages(user string) []entities.Message {
	//get the messages from the database
	log.Println(conn)
	val, ok := conn.Get(user)
	if !ok {
		return []entities.Message{}
	}
	return val.Messages
}

func getFullConn(user string) entities.ConversationByUsers {
	//get the messages from the database
	val, ok := conn.Get(user)
	if !ok {
		return entities.ConversationByUsers{}
	}
	return val
}

func appendnews(user string, newMessage entities.Message) []entities.Message {

	newMessages := append(getMessages(user), newMessage)
	var conversation = getFullConn(user)
	conversation.Messages = newMessages
	conn.Set(user, conversation)

	return newMessages

}
