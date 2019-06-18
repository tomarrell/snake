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

type GameState struct {
	Width  int
	Height int
	Snake  snake
	Fruit  []fruit
}

type game struct {
	id         int
	tickrate   int
	width      int
	height     int
	snake      snake
	fruit      []fruit
	inputChan  chan (KeyCode)
	outputChan chan (GameState)
	stopped    bool
	*sync.Mutex
}

func (g *game) stop() {
	g.Lock()
	defer g.Unlock()
	g.stopped = true
}

func (g *game) isStopped() bool {
	g.Lock()
	defer g.Unlock()
	return g.stopped
}

func (g *game) handleCollisions() {
	snakeHead := g.snake.head()

	for i, fruit := range g.fruit {
		if snakeHead.X == fruit.x && snakeHead.Y == fruit.y {
			g.snake.eatFruit(fruit.value)
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
		}

		time.Sleep(time.Duration(sleepTime))
	}
}
