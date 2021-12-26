package main

import (
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"log"

	"github.com/tomarrell/snake/engine"
)

type vPayload struct {
	GameID string
	Width  int
	Height int
	Score  int
	Fruit  []engine.Fruit
	Snake  engine.Snake
}

func signState(state *vPayload) *string {
	if secret == "" {
		log.Println("WARNING: no env variable 'SECRET' provided, signatures will be insecure")
	}

	h := sha256.New()

	s, err := json.Marshal(state)
	if err != nil {
		panic("failed to sign state")
	}

	h.Write([]byte(s))
	h.Write([]byte(secret))
	sum := base64.URLEncoding.EncodeToString(h.Sum(nil))

	return &sum
}
