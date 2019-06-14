package engine

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewFruit(t *testing.T) {
	assert := assert.New(t)

	f := newFruit(100, 100)

	assert.True(f.x < 100)
	assert.True(f.y < 100)
	assert.IsType(fruitValue(0), f.value)
}
