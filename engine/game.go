package engine

import (
	"fmt"
	"time"
)

type KeyCode int

const (
	KeyLeft KeyCode = iota
	KeyRight
	KeyUp
	KeyDown
)

type Game struct {
	ID        int
	Tickrate  int
	Width     int
	Height    int
	inputChan chan (Command)
}

func (g *Game) Run() {
	sleepTime := float32(1*time.Second) / float32(g.Tickrate)

	for {
		fmt.Println("Hello World: Game is running!")
		time.Sleep(time.Duration(sleepTime))
	}
}
