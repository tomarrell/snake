package main

// New game
type part struct {
	X int `json:"x"`
	Y int `json:"y"`
}

type snake struct {
	VelX  int    `json:"velX"`
	VelY  int    `json:"velY"`
	Parts []part `json:"parts"`
}

type newGameRequest struct {
	Width  int   `json:"width"`
	Height int   `json:"height"`
	Snake  snake `json:"snake"`
}

// Validate game
type tick struct {
	VelX int `json:"velX"`
	VelY int `json:"velY"`
}

type validateRequest struct {
	GameID    string  `json:"gameId"`
	Width     int     `json:"width"`
	Height    int     `json:"height"`
	Score     int     `json:"score"`
	Snake     snake   `json:"snake"`
	Fruit     []fruit `json:"fruit"`
	Signature string  `json:"signature"`
	Ticks     []tick  `json:"ticks"`
}

// Response
type fruit struct {
	X int `json:"x"`
	Y int `json:"y"`
}

type signedStateResponse struct {
	GameID    string  `json:"gameId"`
	Width     int     `json:"width"`
	Height    int     `json:"height"`
	Score     int     `json:"score"`
	Fruit     []fruit `json:"fruit"`
	Snake     snake   `json:"snake"`
	Signature string  `json:"signature"`
}
