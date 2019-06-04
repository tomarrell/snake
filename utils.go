package main

import (
	"encoding/json"
	"net/http"
)

// Return success with JSON body
func respondJSON(w http.ResponseWriter, v interface{}) {
	j, err := json.Marshal(v)
	if err != nil {
		http.Error(w, "Failed to marshall JSON response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}
