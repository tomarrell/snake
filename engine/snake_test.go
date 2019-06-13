package engine

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEatFood(t *testing.T) {
	assert := assert.New(t)

	s := newSnake(100, 100)
	assert.Equal(s.length(), 3)

	s.eatFood(2)

	assert.Equal(s.length(), 5)
	assert.Equal(s.parts, []part{
		part{3, 1},
		part{2, 1},
		part{1, 1},
		part{1, 1},
		part{1, 1},
	})
}

func TestMoveUp(t *testing.T) {
	assert := assert.New(t)
	s := newSnake(100, 100)

	s.up()
	assert.Equal(s.length(), 3)
	assert.Equal(s.velY, -1)
	assert.Equal(s.velX, 0)
}

func TestMoveDown(t *testing.T) {
	assert := assert.New(t)
	s := newSnake(100, 100)

	s.down()
	assert.Equal(s.length(), 3)
	assert.Equal(s.velY, 1)
	assert.Equal(s.velX, 0)
}

func TestMoveRight(t *testing.T) {
	assert := assert.New(t)
	s := newSnake(100, 100)

	s.down()
	s.right()
	assert.Equal(s.length(), 3)
	assert.Equal(s.velY, 0)
	assert.Equal(s.velX, 1)
}

func TestMoveLeft(t *testing.T) {
	assert := assert.New(t)
	s := newSnake(100, 100)

	s.down()
	s.left()
	assert.Equal(s.length(), 3)
	assert.Equal(s.velY, 0)
	assert.Equal(s.velX, -1)
}

func TestUpdate(t *testing.T) {
	assert := assert.New(t)
	s := newSnake(100, 100)

	s.up()
	s.update()
	assert.Equal(s.velY, -1)
	assert.Equal(s.head().x, 3)
	assert.Equal(s.head().y, 0)

	s.update()
	assert.Equal(s.head().y, 100)

	s.update()
	assert.Equal(s.head().y, 99)

	s.right()
	s.update()
	assert.Equal(s.head().x, 4)
	assert.Equal(s.length(), 3)
}
