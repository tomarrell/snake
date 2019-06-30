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

	gameID := uuid.New()
	log.Println(gameID, "creating new game")

	width := ng.Width
	height := ng.Height
	score := 0
	fruit := []fruit{}

	// gameID
	// width
	// height
	// score

}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/new", newHandler)

	log.Println("Starting server on port:", "8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
