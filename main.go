package main

import (
	"database/sql"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
	uuid "github.com/satori/go.uuid"
	log "github.com/sirupsen/logrus"
)

const (
	port = ":8080"
)

/// Serve a token to the client to check future requests
func generateToken(w http.ResponseWriter, r *http.Request) {
	u1 := uuid.NewV4()
	respondJSON(w, token{UUID: u1.String()})
}

func main() {
	db, err := sql.Open("sqlite3", "./db.db")
	if err != nil {
		log.Fatal("Failed to open sqlite DB")
	}
	defer db.Close()

	r := mux.NewRouter()

	api := r.PathPrefix("/api").Subrouter()
	api.HandleFunc("/token", generateToken)

	log.Info("Starting server of port ", port)
	log.Fatal(http.ListenAndServe(port, r))
}
