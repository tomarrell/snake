package engine

import (
	"sync"
	"time"
)

// KeyCode represents an input key event code
type KeyCode int

const (
	// KeyLeft shifts the snake to the West
	KeyLeft KeyCode = iota
	// KeyRight shifts the snake to the East
	KeyRight
	// KeyUp shifts the snake to the North
	KeyUp
	// KeyDown shifts the snake to the South
	KeyDown
)

// Game is the stru
type game struct {
	id        int
	tickrate  int
	width     int
	height    int
	snake     snake
	inputChan chan (KeyCode)
	stopped   bool
	*sync.Mutex
}

// Stop prevents further execution of the game
func (g *game) stop() {
	g.Lock()
	defer g.Unlock()
	g.stopped = true
}

// IsStopped returns wether the game is stopped
func (g *game) isStopped() bool {
	g.Lock()
	defer g.Unlock()
	return g.stopped
}

func (g *game) update() {
	g.snake.update()
}

func (g *game) handleInput() {
	select {
	case input := <-g.inputChan:
		switch input {
		case KeyRight:
			g.snake.right()
		case KeyLeft:
			g.snake.left()
		case KeyUp:
			g.snake.up()
		case KeyDown:
			g.snake.down()
		}
	default:
	}
}

// Run begins the main loop execution of the game
func (g *game) run(wg *sync.WaitGroup) {
	defer wg.Done()

	sleepTime := float32(1*time.Second) / float32(g.tickrate)

	for {
		if g.isStopped() {
			break
		}

		// Handle input
		g.handleInput()
		// Update the position of snake based on velocity
		g.update()

		time.Sleep(time.Duration(sleepTime))
	}
}
