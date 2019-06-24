package main

import (
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"github.com/tidwall/gjson"
	"github.com/tomarrell/snake/engine"
)

type httpHandler = func(w http.ResponseWriter, r *http.Request)

const port = ":8080"

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func main() {
	e := engine.NewEngine()
	r := mux.NewRouter()

	r.HandleFunc("/ws", newGameHandler(e))

	log.Println("Starting server of port", port)
	log.Fatal(http.ListenAndServe(port, r))
}

// Create a new game and return the UID of the game
// to the client.
func newGameHandler(e *engine.Engine) httpHandler {
	return func(w http.ResponseWriter, r *http.Request) {
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Println(err)
			return
		}

		connID := uuid.New().String()[:11]
		log.Println(connID, "new connection established")

		var gameID *int
		defer ifNotNil(gameID, e.DestroyGame)

		for {
			_, payload, err := conn.ReadMessage()
			if err != nil {
				log.Println(connID, err)
				http.Error(w, "something went wrong processing the WS message", http.StatusInternalServerError)
				return
			}

			val := gjson.ValidBytes(payload)
			if !val {
				log.Println(connID, "message is not valid json")
				http.Error(w, "message is not valid json", http.StatusBadRequest)
				return
			}

			mtype := gjson.GetBytes(payload, "type").String()

			switch mtype {
			case "new":
				if gameID == nil {
					ng := e.NewGame(80, 80, 10)
					gameID = &ng
					conn.WriteJSON(newAckOk())
				} else {
					conn.WriteJSON(newAckError("game already exists"))
				}
			case "destroy":
				log.Println("got destroy")
				if gameID != nil {

				} else {

				}
			case "input":
				log.Println("got input")
			default:
				log.Println(connID, "received message type not valid")
				http.Error(w, "received message type not valid", http.StatusBadRequest)
				return
			}

			if err := conn.WriteMessage(1, payload); err != nil {
				log.Println(err)
				http.Error(w, "something went wrong sending message", http.StatusInternalServerError)
				return
			}
		}
	}
}

func ifNotNil(val *int, fn func(int)) {
	if val != nil {
		fn(*val)
	}
}
