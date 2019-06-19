package engine

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewFruit(t *testing.T) {
	assert := assert.New(t)

	f := newFruit(100, 100)

	assert.True(f.X < 100)
	assert.True(f.Y < 100)
	assert.IsType(FruitValue(0), f.Value)
}
