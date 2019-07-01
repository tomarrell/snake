package main

import (
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
)

func signState(state *signedStateResponse) *string {
	h := sha256.New()

	s, err := json.Marshal(state)
	if err != nil {
		panic("failed to sign state")
	}

	h.Write([]byte(s))
	sum := base64.URLEncoding.EncodeToString(h.Sum(nil))

	return &sum
}
