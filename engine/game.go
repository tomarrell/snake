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
type game struct {
	id        int
	tickrate  int
	width     int
	height    int
	snake     snake
	inputChan chan (Command)
	stopped   bool
}

// Stop prevents further execution of the game
func (g *game) stop() {
	g.stopped = true
}

// IsStopped returns wether the game is stopped
func (g *game) isStopped() bool {
	return g.stopped
}

func (g *game) update() {
	g.snake.update()
}

// Run begins the main loop execution of the game
func (g *game) run(wg *sync.WaitGroup) {
	defer wg.Done()

	sleepTime := float32(1*time.Second) / float32(g.tickrate)

	for {
		if g.isStopped() {
			break
		}

		fmt.Println("New tick in game:", g.id, "End", g.isStopped())
		time.Sleep(time.Duration(sleepTime))
	}
}
