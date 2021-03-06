package engine

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEatFruit(t *testing.T) {
	assert := assert.New(t)

	s := newSnake(100, 100)
	assert.Equal(s.length(), 3)

	s.eatFruit(2)

	assert.Equal(s.length(), 5)
	assert.Equal(s.Parts, []Part{
		{3, 1},
		{2, 1},
		{1, 1},
		{1, 1},
		{1, 1},
	})
}

func TestMoveUp(t *testing.T) {
	assert := assert.New(t)
	s := newSnake(100, 100)

	s.up()
	assert.Equal(s.length(), 3)
	assert.Equal(s.VelY, -1)
	assert.Equal(s.VelX, 0)
}

func TestMoveDown(t *testing.T) {
	assert := assert.New(t)
	s := newSnake(100, 100)

	s.down()
	assert.Equal(s.length(), 3)
	assert.Equal(s.VelY, 1)
	assert.Equal(s.VelX, 0)
}

func TestMoveRight(t *testing.T) {
	assert := assert.New(t)
	s := newSnake(100, 100)

	s.down()
	s.right()
	assert.Equal(s.length(), 3)
	assert.Equal(s.VelY, 0)
	assert.Equal(s.VelX, 1)
}

func TestMoveLeft(t *testing.T) {
	assert := assert.New(t)
	s := newSnake(100, 100)

	s.down()
	s.left()
	assert.Equal(s.length(), 3)
	assert.Equal(s.VelY, 0)
	assert.Equal(s.VelX, -1)
}

func TestUpdate(t *testing.T) {
	assert := assert.New(t)
	s := newSnake(100, 100)

	s.up()
	s.update()
	assert.Equal(s.VelY, -1)
	assert.Equal(s.head().X, 3)
	assert.Equal(s.head().Y, 0)

	s.update()
	assert.Equal(s.head().Y, 99)

	s.update()
	assert.Equal(s.head().Y, 98)

	s.right()
	s.update()
	assert.Equal(s.head().X, 4)
	assert.Equal(s.length(), 3)
}
