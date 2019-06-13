package engine

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewGame_MultipleGames(t *testing.T) {
	assert := assert.New(t)
	e := NewEngine()

	game1 := e.NewGame(1, 1, 60)
	val1, _ := e.getGame(game1)
	assert.Equal(val1, &Game{0, 60, 1, 1, nil, false})

	game2 := e.NewGame(40, 80, 120)
	val2, _ := e.getGame(game2)
	assert.Equal(val2, &Game{1, 120, 80, 40, nil, false})

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
	assert.True(e.games[0].end)
	assert.False(e.games[1].end)

	e.EndGame(1)
	assert.True(e.games[0].end)
}
