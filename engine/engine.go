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
	inputChan    chan (KeyCode)
	games        []*game
	managedGames []*managedGame
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
		[]Fruit{NewFruit(width, height), NewFruit(width, height)},
		0,
		nil,
		nil,
		false,
		new(sync.RWMutex),
	}

	e.games = append(e.games, &newGame)
	return
}

// NewManagedGame creates a new game where the ticks are
// given manually to the engine to validate.
func (e *Engine) NewManagedGame(width, height int) (ID int) {
	ID = len(e.games)
	newGame := managedGame{
		ID,
		width,
		height,
		newSnake(width, height), // TODO take as input
		[]Fruit{NewFruit(width, height), NewFruit(width, height)}, // TODO take as input
		0, // TODO take as input
	}

	e.managedGames = append(e.managedGames, &newGame)
	return
}

// RunManagedGame runs a managed game to completion given
// a slice of ticks to be executed.
func (e *Engine) RunManagedGame(ID int, ticks []Tick) bool {
	return false
}

// StartGame takes a game ID, starts the game and returns
// a channel to handle input events to the game
func (e *Engine) StartGame(ID int, outputChan chan (GameState)) (chan (GameState), error) {
	var game *game

	game, exists := e.getGame(ID)
	if exists == false {
		return nil, errors.New("no game with given ID")
	}

	game.inputChan = make(chan (KeyCode), 1)
	if outputChan != nil {
		game.outputChan = outputChan
	} else {
		game.outputChan = make(chan (GameState))
	}

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

// DestroyGame stops and then removes the
// game from the engine. Irrecoverable.
func (e *Engine) DestroyGame(ID int) {
	for i, g := range e.games {
		g.stop()
		if g.id == ID {
			// Prevent realloc by swapping element with
			// last and reducing length of array
			e.games[i] = e.games[len(e.games)-1]
			e.games = e.games[:len(e.games)-1]
		}
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

func (e *Engine) getManagedGame(ID int) (*managedGame, bool) {
	for _, g := range e.managedGames {
		if g.id == ID {
			return g, true
		}
	}

	return nil, false
}
