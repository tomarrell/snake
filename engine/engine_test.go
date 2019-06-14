package engine

import (
	"github.com/stretchr/testify/assert"
	"sync"
	"testing"
)

func TestNewGame_MultipleGames(t *testing.T) {
	assert := assert.New(t)
	e := NewEngine()

	game1 := e.NewGame(1, 1, 60)
	val1, _ := e.getGame(game1)
	assert.Equal(
		val1,
		&game{
			0,
			60,
			1,
			1,
			newSnake(1, 1),
			nil,
			false,
			new(sync.Mutex),
		})

	game2 := e.NewGame(40, 80, 120)
	val2, _ := e.getGame(game2)
	assert.Equal(
		val2,
		&game{
			1,
			120,
			80,
			40,
			newSnake(40, 80),
			nil,
			false,
			new(sync.Mutex),
		})

	assert.Len(e.games, 2)
}

func TestEndGame(t *testing.T) {
	assert := assert.New(t)
	e := NewEngine()

	game1 := e.NewGame(1, 1, 60)
	e.getGame(game1)

	game2 := e.NewGame(40, 80, 120)
	e.getGame(game2)

	e.EndGame(0)
	assert.True(e.games[0].isStopped())
	assert.False(e.games[1].isStopped())

	e.EndGame(1)
	assert.True(e.games[0].isStopped())
}
