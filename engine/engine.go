package engine

import (
	"errors"
	"sync"
)

var (
	e    *Engine
	once sync.Once
	wg   sync.WaitGroup
)

// Engine controls the entire set
// of games being played within it
type Engine struct {
	inputChan chan (KeyCode)
	games     []*game
}

// Start blocks the current thread the Engine
// runs the games
func (e *Engine) Start() {
	wg.Wait()
}

// NewEngine constructs a new singleton instance
// of Engine, or returns the existing one
func NewEngine() *Engine {
	once.Do(func() {
		e = &Engine{
			games: []*game{},
		}
	})

	e.Purge()

	return e
}

// NewGame creates a new game of snake to be run by the engine
func (e *Engine) NewGame(width, height, tickrate int) (ID int) {
	ID = len(e.games)
	newGame := game{
		ID,
		tickrate,
		width,
		height,
		newSnake(width, height),
		[]Fruit{newFruit(width, height), newFruit(width, height)},
		0,
		nil,
		nil,
		false,
		new(sync.RWMutex),
	}

	e.games = append(e.games, &newGame)
	return
}

// StartGame takes a game ID, starts the game and returns
// a channel to handle input events to the game
func (e *Engine) StartGame(ID int) (chan (GameState), error) {
	var game *game

	game, exists := e.getGame(ID)
	if exists == false {
		return nil, errors.New("no game with given ID")
	}

	game.inputChan = make(chan (KeyCode), 1)
	game.outputChan = make(chan (GameState))

	wg.Add(1)
	go game.run(&wg)

	return game.outputChan, nil
}

// SendInput forwards the given KeyCode
// on to the game routine
func (e *Engine) SendInput(ID int, key KeyCode) error {
	g, ok := e.getGame(ID)
	if !ok {
		return errors.New("No game found with ID")
	}

	g.RLock()
	velX := g.snake.velX
	g.RUnlock()

	if velX != 0 {
		switch key {
		case KeyUp, KeyDown:
			g.inputChan <- key
		}
	} else {
		switch key {
		case KeyRight, KeyLeft:
			g.inputChan <- key
		}
	}

	return nil
}

// EndGame stops a game with the given ID from
// running. It however does not remove it from
// the Engine.
func (e *Engine) EndGame(ID int) {
	game, _ := e.getGame(ID)
	if game != nil {
		game.stop()
	}
}

// Purge destroys all the currently running games.
// This is a completely lossy action.
func (e *Engine) Purge() {
	e.games = nil
}

func (e *Engine) getGame(ID int) (*game, bool) {
	for _, g := range e.games {
		if g.id == ID {
			return g, true
		}
	}

	return nil, false
}
