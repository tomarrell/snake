package main

import (
	"github.com/tomarrell/snake/engine"
)

// New game
type newGameRequest struct {
	Width  int          `json:"width"`
	Height int          `json:"height"`
	Snake  engine.Snake `json:"snake"`
}

// Validate game
type validateRequest struct {
	GameID    string         `json:"gameId"`
	Width     int            `json:"width"`
	Height    int            `json:"height"`
	Score     int            `json:"score"`
	Snake     engine.Snake   `json:"snake"`
	Fruit     []engine.Fruit `json:"fruit"`
	Signature *string        `json:"signature"`
	Ticks     []engine.Tick  `json:"ticks"`
}

// Response
type signedStateResponse struct {
	GameID    string         `json:"gameId"`
	Width     int            `json:"width"`
	Height    int            `json:"height"`
	Score     int            `json:"score"`
	Fruit     []engine.Fruit `json:"fruit"`
	Snake     engine.Snake   `json:"snake"`
	Signature *string        `json:"signature"`
}
