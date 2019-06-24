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

	r.HandleFunc("/ws", websocketHandler(e))

	log.Println("Starting server of port", port)
	log.Fatal(http.ListenAndServe(port, r))
}

// Create a new game and return the UID of the game
// to the client.
func websocketHandler(e *engine.Engine) httpHandler {
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
				conn.WriteJSON(newAckError("unable to read message"))
				return
			}

			val := gjson.ValidBytes(payload)
			if !val {
				log.Println(connID, "message is not valid json")
				conn.WriteJSON(newAckError("message is not valid json"))
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
				ifNotNil(gameID, e.DestroyGame)
				conn.WriteJSON(newAckOk())
			case "input":
				log.Println("got input")
			default:
				log.Println(connID, "received message type not valid")
				conn.WriteJSON(newAckError("invalid message type"))
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
