package engine

import (
	"errors"
	"sync"
)

var (
	e     *Engine
	once  sync.Once
	mutex sync.Mutex
	wg    sync.WaitGroup
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
			games:     []*game{},
			inputChan: nil,
		}
	})

	return e
}

// NewGame creates a new game of snake to be run by the engine
func (e *Engine) NewGame(height, width, tickrate int) (ID int) {
	mutex.Lock()
	defer mutex.Unlock()

	ID = len(e.games)
	newGame := game{
		ID,
		tickrate,
		width,
		height,
		newSnake(height, width),
		[]fruit{},
		nil,
		false,
		new(sync.Mutex),
	}

	e.games = append(e.games, &newGame)
	return
}

// StartGame takes a game ID, starts the game and returns
// a channel to handle input events to the game
func (e *Engine) StartGame(ID int) error {
	var game *game

	game, exists := e.getGame(ID)
	if exists == false {
		return errors.New("no game with given ID")
	}

	game.inputChan = make(chan (KeyCode))
	wg.Add(1)
	go game.run(&wg)

	return nil
}

// SendInput forwards the given KeyCode
// on to the game routine
func (e *Engine) SendInput(ID int, key KeyCode) error {
	g, ok := e.getGame(ID)
	if !ok {
		return errors.New("No game found with ID")
	}

	g.inputChan <- key

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

func (e *Engine) getGame(ID int) (*game, bool) {
	for _, g := range e.games {
		if g.id == ID {
			return g, true
		}
	}

	return nil, false
}
