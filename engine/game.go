package engine

import (
	"fmt"
	"sync"
	"time"
)

// KeyCode represents an input key event code
type KeyCode int

const (
	KeyLeft KeyCode = iota
	KeyRight
	KeyUp
	KeyDown
)

// Game is the stru
type Game struct {
	ID        int
	Tickrate  int
	Width     int
	Height    int
	inputChan chan (Command)
	End       bool
}

func (g *Game) stop() {
	g.End = true
}

func (g *Game) run(wg *sync.WaitGroup) {
	defer wg.Done()

	sleepTime := float32(1*time.Second) / float32(g.Tickrate)

	for {
		if g.End {
			break
		}

		fmt.Println("New tick in game:", g.ID, "End", g.End)
		time.Sleep(time.Duration(sleepTime))
	}
}
