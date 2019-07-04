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
	Width    int     `json:"width"`
	Height   int     `json:"height"`
	Snake    Snake   `json:"snake"`
	Fruit    []Fruit `json:"fruit"`
	Score    int     `json:"score"`
	Finished bool    `json:"finished"`
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
	finished   bool
	*sync.RWMutex
}

func newGame(id, tickrate, width, height int) *game {
	return &game{
		id,
		tickrate,
		width,
		height,
		newSnake(width, height),
		[]Fruit{NewFruit(width, height), NewFruit(width, height)},
		0,
		nil,
		nil,
		false,
		new(sync.RWMutex),
	}
}

func (g *game) stop() {
	g.Lock()
	g.finished = true
	g.Unlock()
}

func (g *game) isFinished() bool {
	g.RLock()
	finished := g.finished
	g.RUnlock()

	return finished
}

func (g *game) handleFruitCollisions() {
	g.Lock()
	defer g.Unlock()

	head := g.snake.head()

	for i, fruit := range g.fruit {
		if head.X == fruit.X && head.Y == fruit.Y {
			g.score += int(fruit.Value)
			g.snake.eatFruit(fruit.Value)
			g.fruit[i] = NewFruit(g.width, g.height)
		}
	}
}

func (g *game) checkSelfCollision() {
	g.Lock()
	defer g.Unlock()

	head := g.snake.head()

	for i := 1; i < len(g.snake.Parts); i++ {
		part := g.snake.Parts[i]
		if head.X == part.X && head.Y == part.Y {
			g.finished = true
			return
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
		if g.isFinished() {
			break
		}

		g.handleFruitCollisions()
		g.handleInput()
		g.update()
		g.checkSelfCollision()

		g.outputChan <- GameState{
			g.width,
			g.height,
			g.snake,
			append([]Fruit{}, g.fruit...),
			g.score,
			g.finished,
		}

		time.Sleep(time.Duration(sleepTime))
	}
}
