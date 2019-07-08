package engine

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestManagedGameRun(t *testing.T) {
	assert := assert.New(t)

	mg := newManagedGame(
		1,
		20,
		20,
		0,
		newSnake(20, 20),
		[]Fruit{{1, 5, 5}, {2, 10, 10}},
	)

	ticks := []Tick{
		{1, 0},
		{1, 0},
		{0, 1},
		{0, 1},
		{0, 1},
		{0, 1},
	}

	valid, err := mg.run(ticks)

	assert.NoError(err)
	assert.True(valid)
}

func TestManagedGameRun_2(t *testing.T) {
	assert := assert.New(t)

	mg := newManagedGame(
		1,
		20,
		20,
		0,
		newSnake(20, 20),
		[]Fruit{{1, 5, 5}, {2, 10, 10}},
	)

	ticks := []Tick{
		{1, 0},
		{0, 1},
		{0, 1},
		{1, 0},
		{0, 1},
		{0, 1},
		{1, 0},
		{1, 0},
		{0, 1},
		{0, 1},
		{0, 1},
		{1, 0},
		{0, 1},
		{0, 1},
		{1, 0},
		{1, 0},
	}

	valid, err := mg.run(ticks)

	assert.NoError(err)
	assert.True(valid)
}

var tickTests = []struct {
	name string
	velX int
	velY int
	out  bool
}{
	{"going right", 1, 0, true},
	{"going left", -1, 0, true},
	{"going down", 0, 1, true},
	{"going up", 0, -1, true},
	{"invalid", -1, -1, false},
	{"sneaky", 3, -2, false},
	{"negative sneaky", -303, 302, false},
}

func TestValidateTick(t *testing.T) {
	assert := assert.New(t)
	for _, tt := range tickTests {
		t.Run(tt.name, func(t *testing.T) {
			valid := validateTick(Tick{tt.velX, tt.velY})
			assert.Equalf(tt.out, valid, "unexpected tick validation: %s", tt.name)
		})
	}
}
