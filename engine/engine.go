package engine

import (
	"errors"
	"sync"
)

var (
	e     *Engine
	once  sync.Once
	mutex sync.Mutex
)

type Command struct {
	gameID int
	key    KeyCode
}

// Engine controls the entire set
// of games being played within it
type Engine struct {
	inputChan chan (Command)
	games     []Game
}

// NewEngine constructs a new singleton instance
// of Engine, or returns the existing one
func NewEngine() *Engine {
	once.Do(func() {
		e = &Engine{
			games:     []Game{},
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
	}

	e.games = append(e.games, newGame)

	return ID
}

// StartGame takes a game ID, starts the game and returns
// a channel to handle input events to the game
func (e *Engine) StartGame(ID int) (chan (Command), error) {
	var game *Game

	for _, g := range e.games {
		if g.ID == ID {
			game = &g
			break
		}
	}

	if game == nil {
		return nil, errors.New("no game with given ID")
	}

	game.inputChan = make(chan (Command))
	game.Run()

	return game.inputChan, nil
}

func (e *Engine) getGame(ID int) (*Game, bool) {
	for _, g := range e.games {
		if g.ID == ID {
			val := &g
			return val, true
		}
	}

	return nil, false
}
