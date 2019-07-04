package engine

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
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
			val1.fruit,
			0,
			nil,
			nil,
			false,
			new(sync.RWMutex),
		})

	game2 := e.NewGame(40, 80, 120)
	val2, _ := e.getGame(game2)
	assert.Equal(
		val2,
		&game{
			1,
			120,
			40,
			80,
			newSnake(40, 80),
			val2.fruit,
			0,
			nil,
			nil,
			false,
			new(sync.RWMutex),
		})

	assert.Len(e.games, 2)
}

func TestNewManagedGame(t *testing.T) {
	assert := assert.New(t)
	e := NewEngine()

	game1 := e.NewManagedGame(20, 20, 0, newSnake(1, 1), []Fruit{})
	val1, _ := e.getManagedGame(game1)
	assert.Equal(
		val1,
		&ManagedGame{
			0,
			20,
			20,
			newSnake(1, 1),
			val1.Fruit,
			0,
		})

	assert.Len(e.managedGames, 1)
}

func TestEndGame(t *testing.T) {
	assert := assert.New(t)
	e := NewEngine()

	game1 := e.NewGame(1, 1, 60)
	e.getGame(game1)

	game2 := e.NewGame(40, 80, 120)
	e.getGame(game2)

	e.EndGame(0)
	assert.True(e.games[0].isFinished())
	assert.False(e.games[1].isFinished())

	e.EndGame(1)
	assert.True(e.games[0].isFinished())
}

func TestGameDestroy(t *testing.T) {
	assert := assert.New(t)
	e := NewEngine()

	g1 := e.NewGame(1, 1, 60)
	gameToDestroy := e.NewGame(1, 1, 60)
	g2 := e.NewGame(1, 1, 60)

	e.DestroyGame(gameToDestroy)

	assert.Len(e.games, 2)
	assert.NotNil(e.getGame(g1))
	assert.NotNil(e.getGame(g2))
}

func TestEnginePurge(t *testing.T) {
	assert := assert.New(t)
	e := NewEngine()

	e.NewGame(1, 1, 60)
	e.Purge()

	assert.Len(e.games, 0)
}
