package engine

import (
	"sync"
)

var (
	e *Engine
)

// Engine controls the entire set
// of games being played within it
type Engine struct {
	games []game
}

type game struct {
	board board
}

type board struct {
	width  int
	height int
}

// NewEngine constructs a new singleton instance
// of Engine, or returns the existing one
func NewEngine() *Engine {
	if e != nil {
		return e
	}

	return &Engine{
		games: []game{},
	}
}
