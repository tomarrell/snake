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

	s.Signature = signState(&s)
	log.Printf("signed new game %v", s.Signature)

	writeJSON(w, s)
}

func valdiatePath(w http.ResponseWriter, r *http.Request) {

}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/new", newHandler)
	r.HandleFunc("/validate", newHandler)

	log.Println("Starting server on port:", "8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
