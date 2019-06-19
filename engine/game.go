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

// GameState represents all the information
// returned to the client for rendering
type GameState struct {
	Width  int
	Height int
	Snake  Snake
	Fruit  []Fruit
	Score  int
}

type game struct {
	id         int
	tickrate   int
	width      int
	height     int
	snake      Snake
	fruit      []Fruit
	score      int
	inputChan  chan (KeyCode)
	outputChan chan (GameState)
	stopped    bool
	*sync.RWMutex
}

func (g *game) stop() {
	g.Lock()
	defer g.Unlock()
	g.stopped = true
}

func (g *game) isStopped() bool {
	g.RLock()
	defer g.RUnlock()
	return g.stopped
}

func (g *game) handleCollisions() {
	g.Lock()
	defer g.Unlock()

	snakeHead := g.snake.head()

	for i, fruit := range g.fruit {
		if snakeHead.X == fruit.X && snakeHead.Y == fruit.Y {
			g.score += int(fruit.Value)
			g.snake.eatFruit(fruit.Value)
			g.fruit[i] = newFruit(g.width, g.height)
		}
	}
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

func (g *game) update() {
	g.Lock()
	defer g.Unlock()
	g.snake.update()
}

func (g *game) run(wg *sync.WaitGroup) {
	defer wg.Done()

	sleepTime := float32(1*time.Second) / float32(g.tickrate)

	for {
		if g.isStopped() {
			break
		}

		g.handleCollisions()
		g.handleInput()
		g.update()

		g.outputChan <- GameState{
			g.width,
			g.height,
			g.snake,
			g.fruit,
			g.score,
		}

		time.Sleep(time.Duration(sleepTime))
	}
}
