package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
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

		var gameID *int
		defer ifNotNil(gameID, e.DestroyGame)

		for {
			messageType, p, err := conn.ReadMessage()
			if err != nil {
				log.Println(err)
				http.Error(w, "something went wrong processing the WS message", http.StatusInternalServerError)
				return
			}
			if messageType != 1 {
				http.Error(w, "unsupported message type, please send text", http.StatusBadRequest)
				return
			}

			switch string(p) {
			case "new":
				if gameID != nil {
					id := e.NewGame(80, 80, 10)
					gameID = &id
				}
			case "destroy":
			}

			if err := conn.WriteMessage(messageType, p); err != nil {
				log.Println(err)
				http.Error(w, "something went wrong sending the WS message", http.StatusInternalServerError)
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
