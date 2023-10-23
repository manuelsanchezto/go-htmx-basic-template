package pages

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"htmx.try/m/v2/pkg/dbconn"
	"htmx.try/m/v2/pkg/domain"
	"htmx.try/m/v2/pkg/domain/dto"
)

var conn = *dbconn.NewInMemoryDB()
var template = "index.html"
var url = "http://manuelsanchez.sisnet360.com:8082/"

func Index(c echo.Context) error {
	user := c.QueryParam("user")
	conversaciones := getConversations(user)

	return c.Render(200, template, domain.InterfaceResponseFull{
		User:          user,
		Conversations: conversaciones,
	})
}

func StartNewConversation(c echo.Context) error {
	user := c.QueryParam("user")
	conn.DeleteData(user)
	conn.DeleteResponses(user)
	conversaciones := getConversations(user)
	return c.Render(200, template, domain.InterfaceResponseFull{
		User:          user,
		Conversations: conversaciones,
	})
}

func AddMessage(c echo.Context) error {
	user := c.FormValue("user")
	question := c.FormValue("question")
	module := c.FormValue("module")
	time := time.Now().Format("15:04:05")

	quest := domain.NewMessage(question, time)
	answ := domain.NewMessage("", "")

	conversacion := domain.NewConversation(quest, answ, false, "invisible", true)
	var conversaciones = getConversations(user)
	if len(conversaciones) > 0 {
		indiceUltimo := len(conversaciones) - 1
		conversaciones[indiceUltimo].IsLast = false
	}
	conversaciones = append(conversaciones, conversacion)
	var cosas = getFullConn(user)

	cosas.Conversations = conversaciones
	conn.SetData(user, cosas)

	go generateMessage(user, module)

	return c.Render(http.StatusOK, template, domain.InterfaceResponseFull{
		User:          user,
		Conversations: conversaciones,
	})
}

func CloseActions(c echo.Context) error {
	user := c.FormValue("user")
	conversaciones := getConversations(user)
	indice := len(conversaciones) - 1
	conversaciones[indice].Actions = "invisible"

	return c.Render(http.StatusOK, template, domain.InterfaceResponseFull{
		User:          user,
		Conversations: conversaciones,
	})
}

func GetBussinessLine(c echo.Context) error {
	user := c.FormValue("user")
	log.Println(user)
	respuesta := getLastResponse(user)
	if respuesta == nil {
		log.Println("No hay respuesta")
		return nil
	}
	mensajeServidor := getLastConversation(user)
	if mensajeServidor == nil {
		log.Println("No hay mensaje del servidor")
		return nil
	}
	loadBussinessLine(*respuesta, mensajeServidor.Question.Text)

	return nil

}

func NewMongoDB() *mongo.Client {
	// Set client options
	clientOptions := options.Client().ApplyURI("mongodb://root:example@20.56.93.5:27017")

	// Connect to MongoDB
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")

	return client
}
func loadBussinessLine(respuesta domain.Response, textoServidor string) string {
	var introducir = DataIn{_id: primitive.NewObjectID(), texto: textoServidor}
	err := SaveJSONData(NewMongoDB(), "copilot", "responses", introducir)
	if err != nil {
		return introducir._id.Hex()
	}
	return ""
}

type DataIn struct {
	_id   primitive.ObjectID
	texto string
}

func SaveJSONData(client *mongo.Client, databaseName string, collectionName string, data interface{}) error {
	// Get a handle for your collection
	collection := client.Database(databaseName).Collection(collectionName)

	// Insert a single document
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, err := collection.InsertOne(ctx, data)
	if err != nil {
		return err
	}
	return nil
}

func getLastResponse(user string) *domain.Response {
	vals, ok := conn.GetResponses(user)
	if !ok {
		return nil
	}
	val := vals[len(vals)-1]
	return &val
}

func getLastConversation(user string) *domain.Conversation {
	vals := getConversations(user)
	for _, val := range vals {
		if val.IsLast {
			return &val
		}
	}
	return nil
}

func getFullConn(user string) domain.InterfaceResponseFull {
	//get the user conversations from the database
	val, ok := conn.GetData(user)
	if !ok {
		return domain.InterfaceResponseFull{}
	}
	return val
}

func getConversations(user string) []domain.Conversation {
	val, ok := conn.GetData(user)
	if !ok {
		return []domain.Conversation{}
	}
	return val.Conversations
}

func generateMessage(user string, module string) {
	time.Sleep(1 * time.Second)
	var conversaciones = getConversations(user)

	for pos, val := range conversaciones {
		if !val.IsAnswered {
			resp := requestAnswer(conversaciones[pos].Question, user, module)
			var response string
			//response = "Respuesta del servidor"
			//resp := &response

			if resp == nil {
				response = "Ha ocurrido un error"
			} else {
				response = *resp
				conversaciones[pos].Actions = "visible"
			}
			conversaciones[pos].Answer = domain.Message{Text: response, Time: time.Now().Format("15:04:05")}
			conversaciones[pos].IsAnswered = true
			var updatedConv = getFullConn(user)
			updatedConv.Conversations = conversaciones
			conn.SetData(user, updatedConv)
			return
		}
	}
}

func recoverExample() *dto.Base {
	var respuesta dto.Base
	raw, err := os.ReadFile("/home/usuario/Escritorio/ejemplo.json")
	if err != nil {
		fmt.Println(err.Error())
	}
	json.Unmarshal(raw, &respuesta)
	return &respuesta
}

func requestAnswer(message domain.Message, user string, module string) *string {

	if !checkStatus() {
		err := "Server disconnected"
		log.Println(err)
		return nil
	}

	messageNoSpaces := strings.Replace(message.Text, " ", "%20", -1)
	base, errBase := getBase(messageNoSpaces)

	sections, errSections := getSections(messageNoSpaces, module)

	if errBase != nil || errSections != nil {
		fmt.Println(errBase)
		fmt.Println(errSections)
		return nil
	}

	producto := base.Result.Business_line_data.Business_line.Producto
	var props []string
	props = append(props, sections.Result.AdditionalProp1, sections.Result.AdditionalProp2, sections.Result.AdditionalProp3)
	mensaje := fmt.Sprintf("Si te he entendido correctamente, quieres que realice cambios sobre la linea de negocio %s, sobre las siguientes secciones:\n -%v", producto, props)
	//Guardamos respuesta en base de datos
	response := domain.NewResponse(props, base.Result.Business_line_data)
	conn.SetResponse(user, response)
	return &mensaje
}

func checkStatus() bool {
	res, err := http.Get(url + "health_check")
	if err != nil {
		log.Println("Impossible to build request: " + err.Error())
		return false
	}
	if res.StatusCode == 200 {
		return true
	}
	return false
}

func getBase(message string) (*dto.Base, error) {
	res, err := http.Get(url + "base?query=" + message)
	if err != nil {
		log.Println("Impossible to build request: " + err.Error())
		return recoverExample(), nil
		//return nil, err
	}
	if res.StatusCode == 200 {
		resBody, err := io.ReadAll(res.Body)
		if err != nil {
			log.Println("Impossible to read all body of response " + err.Error())
			return nil, err
		}

		var response dto.Base
		err = json.Unmarshal(resBody, &response)
		if err != nil {
			log.Println("Impossible to parse the response " + err.Error())
			return nil, err
		}
		return &response, nil
	} else {
		return recoverExample(), nil
	}

	/*error := errors.New("Error: response received with status code " + res.Status)
	log.Println(error.Error())
	return nil, error*/

}

func getSections(message string, module string) (*dto.SectionsToEdit, error) {
	res, err := http.Get(url + "sections_to_edit?query=" + message + "&module=" + module)
	if err != nil {
		log.Println("Impossible to build request: " + err.Error())
		return nil, err
	}

	if res.StatusCode == 200 {
		resBody, err := io.ReadAll(res.Body)
		if err != nil {
			log.Println("Impossible to read all body of response " + err.Error())
			return nil, err
		}

		var response dto.SectionsToEdit
		err = json.Unmarshal(resBody, &response)
		if err != nil {
			log.Println("Impossible to parse the response " + err.Error())
			return nil, err
		}
		log.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>")
		log.Println(response)
		log.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>")

		return &response, nil
	}

	error := errors.New("Error: response received with status code " + res.Status)
	log.Println(error.Error())
	return nil, error
}
