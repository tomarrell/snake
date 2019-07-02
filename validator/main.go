package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/tomarrell/snake/engine"
)

func newHandler(w http.ResponseWriter, r *http.Request) {
	var ng newGameRequest

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&ng)
	if err != nil {
		http.Error(w, "invalid json body", http.StatusBadRequest)
		return
	}

	gameID := uuid.New().String()
	log.Println(gameID, "creating new game")
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

func validatePath(w http.ResponseWriter, r *http.Request) {
	e := engine.NewEngine()
	var vr validateRequest

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&vr)
	if err != nil {
		http.Error(w, "invalid json body", http.StatusBadRequest)
	}

	toSign := vPayload{
		vr.GameID,
		vr.Width,
		vr.Height,
		vr.Score,
		vr.Fruit,
		vr.Snake,
	}

	if *signState(&toSign) != *vr.Signature {
		log.Println(vr.GameID, "invalid signature")
		http.Error(w, "invalid payload signature", http.StatusUnauthorized)
		return
	}

	g := e.NewManagedGame(vr.Width, vr.Height)

	// setup game
	// play each tick against the snake position
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/new", newHandler)
	r.HandleFunc("/validate", validatePath)

	log.Println("Starting server on port:", "8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
