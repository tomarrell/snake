package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func newHandler(w http.ResponseWriter, r *http.Request) {
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", newHandler)

	log.Fatal(http.ListenAndServe(":8080", r))
}
