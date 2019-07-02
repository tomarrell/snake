package main

import (
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"

	"github.com/tomarrell/snake/engine"
)

const secret = "secret_key_read_from_env"

type vPayload struct {
	GameID string
	Width  int
	Height int
	Score  int
	Fruit  []engine.Fruit
	Snake  snake
}

func signState(state *vPayload) *string {
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
