package main

import (
	"encoding/json"
	"net/http"
)

func writeJSON(w http.ResponseWriter, payload interface{}) {
	j, err := json.Marshal(payload)
	if err != nil {
		http.Error(w, "failed to marshal response payload", http.StatusInternalServerError)
	}

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Write(j)
}
