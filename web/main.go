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
		var gameStateChan = make(chan (engine.GameState))
		var writeChan = make(chan (interface{}))
		defer ifNotNil(gameID, e.DestroyGame)

		go func(gsc chan (engine.GameState), wc chan (interface{})) {
			for {
				select {
				case m := <-writeChan:
					conn.WriteJSON(m)
				case s := <-gsc:
					conn.WriteJSON(s)
				}
			}
		}(gameStateChan, writeChan)

		for {
			_, payload, err := conn.ReadMessage()
			if err != nil {
				log.Println(connID, err)
				writeChan <- newAckError("unable to read message")
				return
			}

			val := gjson.ValidBytes(payload)
			if !val {
				log.Println(connID, "message is not valid json")
				writeChan <- newAckError("message is not valid json")
				return
			}

			mtype := gjson.GetBytes(payload, "type").String()

			switch mtype {
			case "new":
				if gameID == nil {
					log.Println(connID, "starting a new game")
					ng := e.NewGame(80, 80, 10)
					_, err = e.StartGame(ng, gameStateChan)
					if err != nil {
						writeChan <- newAckError(err.Error())
						break
					}

					gameID = &ng
					writeChan <- newAckOk()
				} else {
					writeChan <- newAckError("game already exists")
				}
			case "destroy":
				if gameID != nil {
					log.Println(connID, "game destroyed")
					e.DestroyGame(*gameID)
					writeChan <- newAckOk()
				} else {
					writeChan <- newAckError("no game exists, create one first with type: new")
				}
			case "input":
				if gameID != nil {
					dir := gjson.GetBytes(payload, "direction").String()
					handleInput(writeChan, e, *gameID, dir)
				} else {
					writeChan <- newAckError("no game exists, create one first with type: new")
				}
			default:
				log.Println(connID, "received message type not valid")
				writeChan <- newAckError("invalid message type")
				return
			}
		}

	}
}

func handleInput(writeChan chan (interface{}), e *engine.Engine, id int, input string) {
	switch input {
	case "left":
		e.SendInput(id, engine.KeyLeft)
	case "right":
		e.SendInput(id, engine.KeyRight)
	case "up":
		e.SendInput(id, engine.KeyUp)
	case "down":
		e.SendInput(id, engine.KeyDown)
	default:
		writeChan <- newAckError("invalid direction")
	}
}

func ifNotNil(val *int, fn func(int)) {
	if val != nil {
		fn(*val)
	}
}
