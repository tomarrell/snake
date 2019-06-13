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

// Command is a combined input key
// for a given game ID
type Command struct {
	gameID int
	key    KeyCode
}

// Engine controls the entire set
// of games being played within it
type Engine struct {
	inputChan chan (Command)
	games     []*Game
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
			games:     []*Game{},
			inputChan: nil,
		}
	})

	return e
}

// NewGame creates a new game of snake to be run by the engine
func (e *Engine) NewGame(height, width, tickrate int) int {
	mutex.Lock()
	defer mutex.Unlock()

	ID := len(e.games)

	newGame := Game{
		ID,
		tickrate,
		width,
		height,
		nil,
		false,
	}

	e.games = append(e.games, &newGame)

	return ID
}

// StartGame takes a game ID, starts the game and returns
// a channel to handle input events to the game
func (e *Engine) StartGame(ID int) (chan (Command), error) {
	var game *Game

	for _, g := range e.games {
		if g.ID == ID {
			game = g
			break
		}
	}

	if game == nil {
		return nil, errors.New("no game with given ID")
	}

	game.inputChan = make(chan (Command))
	wg.Add(1)
	go game.run(&wg)

	return game.inputChan, nil
}

// EndGame stops a game with the given ID from
// running. It however does not remove it from
// the Engine.
func (e *Engine) EndGame(ID int) {
	for i := range e.games {
		if e.games[i].ID == ID {
			e.games[i].stop()
		}
	}
}

func (e *Engine) getGame(ID int) (*Game, bool) {
	for _, g := range e.games {
		if g.ID == ID {
			return g, true
		}
	}

	return nil, false
}
