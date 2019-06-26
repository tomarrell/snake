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

func serveIndex(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "static/index.html")
}

func main() {
	e := engine.NewEngine()
	r := mux.NewRouter()

	r.HandleFunc("/", serveIndex)
	r.HandleFunc("/ws", websocketHandler(e))

	log.Println("Starting server of port", port)
	log.Fatal(http.ListenAndServe(port, r))
}

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

		// Writer routine
		go messageWriter(gameStateChan, writeChan, conn)
		go messageReader(e, gameID, conn, connID, writeChan, gameStateChan)
	}
}

func messageReader(
	e *engine.Engine,
	gameID *int,
	conn *websocket.Conn,
	connID string,
	wc chan (interface{}),
	gsc chan (engine.GameState),
) {
	for {
		_, payload, err := conn.ReadMessage()
		if err != nil {
			log.Println(connID, err)
			wc <- newAckError("unable to read message")
			return
		}

		val := gjson.ValidBytes(payload)
		if !val {
			log.Println(connID, "message is not valid json")
			wc <- newAckError("message is not valid json")
			return
		}

		mtype := gjson.GetBytes(payload, "type").String()

		switch mtype {
		case "new":
			if gameID == nil {
				log.Println(connID, "starting a new game")

				w := gjson.GetBytes(payload, "width").Int()
				h := gjson.GetBytes(payload, "height").Int()
				t := gjson.GetBytes(payload, "tick").Int()

				if w == 0 || h == 0 || t == 0 {
					wc <- newAckError("one of width, height, tick cannot be undefined")
					break
				}

				ng := e.NewGame(int(w), int(h), int(t))
				_, err = e.StartGame(ng, gsc)
				if err != nil {
					wc <- newAckError(err.Error())
					break
				}

				gameID = &ng
				wc <- newAckOk()
			} else {
				wc <- newAckError("game already exists")
			}
		case "destroy":
			if gameID != nil {
				e.EndGame(*gameID)
				e.DestroyGame(*gameID)
				gameID = nil
				log.Println(connID, "game destroyed")
				wc <- newAckOk()
			} else {
				wc <- newAckError("no game exists, create one first with type: new")
			}
		case "input":
			if gameID != nil {
				dir := gjson.GetBytes(payload, "direction").String()
				handleInput(wc, e, *gameID, dir)
			} else {
				wc <- newAckError("no game exists, create one first with type: new")
			}
		default:
			log.Println(connID, "received message type not valid")
			wc <- newAckError("invalid message type")
			return
		}
	}
}

func messageWriter(gsc chan (engine.GameState), wc chan (interface{}), conn *websocket.Conn) {
	for {
		select {
		case m := <-wc:
			conn.WriteJSON(m)
		case s := <-gsc:
			conn.WriteJSON(struct {
				MType string           `json:"type"`
				Data  engine.GameState `json:"data"`
			}{
				MType: "state",
				Data:  s,
			})
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
