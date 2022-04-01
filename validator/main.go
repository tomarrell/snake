package main

import (
	"encoding/json"
	"flag"
	"log"
	"net/http"
	"os"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/tomarrell/snake/engine"
)

var (
	secret = os.Getenv("SECRET")
	port   = flag.String("port", "8080", "port to run the web server on")
)

func main() {
	flag.Parse()

	r := mux.NewRouter()

	r.HandleFunc("/snake/new", cors(newHandler)).Methods(http.MethodPost)
	r.HandleFunc("/snake/validate", cors(validatePath)).Methods(http.MethodPost)

	p := ":" + *port
	log.Println("Starting server of port", p)
	log.Fatal(http.ListenAndServe(p, r))
}

// Handle creating a new managed snake game
func newHandler(w http.ResponseWriter, r *http.Request) {
	var ng newGameRequest

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&ng)
	if err != nil {
		http.Error(w, "invalid json body", http.StatusBadRequest)
		return
	}

	if ng.Snake.BoundX != ng.Width || ng.Snake.BoundY != ng.Height {
		http.Error(w, "snake bounds don't match arena bounds", http.StatusBadRequest)
		return
	}

	gameID := uuid.New().String()
	log.Println(gameID[:11], "creating new game")
	s := signedStateResponse{
		gameID,
		ng.Width,
		ng.Height,
		0,
		[]engine.Fruit{
			engine.NewFruit(ng.Width, ng.Height),
			engine.NewFruit(ng.Width, ng.Height),
		},
		ng.Snake,
		nil,
	}

	s.Signature = signState(&vPayload{s.GameID, s.Width, s.Height, s.Score, s.Fruit, s.Snake})
	writeJSON(w, s)
}

// Validate a tick path reaching a piece of fruit
func validatePath(w http.ResponseWriter, r *http.Request) {
	var vr validateRequest
	e := engine.NewEngine()

	decoder := json.NewDecoder(r.Body)
	if decoder.Decode(&vr) != nil {
		http.Error(w, "invalid json body", http.StatusBadRequest)
		return
	}

	if vr.Signature == nil {
		http.Error(w, "empty signature", http.StatusBadRequest)
		return
	}

	checkSign := vPayload{vr.GameID, vr.Width, vr.Height, vr.Score, vr.Fruit, vr.Snake}
	if *signState(&checkSign) != *vr.Signature {
		log.Println(vr.GameID[:11], "invalid signature")
		http.Error(w, "invalid payload signature", http.StatusUnauthorized)
		return
	}

	mg := e.NewManagedGame(vr.Width, vr.Height, vr.Score, vr.Snake, vr.Fruit)
	defer e.DestroyManagedGame(mg)

	g, err := e.RunManagedGame(mg, vr.Ticks)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	log.Printf("%s validated game score: %d points with dimensions: %dx%d", vr.GameID[:11], g.Score, g.Width, g.Height)
	s := signedStateResponse{
		vr.GameID,
		g.Width,
		g.Height,
		g.Score,
		g.Fruit,
		g.Snake,
		nil,
	}

	s.Signature = signState(&vPayload{s.GameID, s.Width, s.Height, s.Score, s.Fruit, s.Snake})
	writeJSON(w, s)
}

func cors(cb func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		cb(w, r)
	}
}
