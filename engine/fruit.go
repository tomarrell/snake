package engine

import "math/rand"

// FruitValue represents the score that the
// repspective piece of fruit will credit
type FruitValue int

const (
	// FruitPink is the least valuable commodity
	FruitPink FruitValue = 1
	// FruitOrange is a moderately valued commodity
	FruitOrange FruitValue = 2
	// FruitGreen is the most valuable commodity
	FruitGreen FruitValue = 5
)

// Fruit is the item which adds score and
// length when eaten by the snake
type Fruit struct {
	Value FruitValue `json:"value"`
	X     int        `json:"x"`
	Y     int        `json:"y"`
}

func newFruit(boundX, boundY int) Fruit {
	fruitValSeed := rand.Intn(10)
	var f FruitValue

	switch {
	case fruitValSeed < 5:
		f = FruitPink
	case fruitValSeed < 8:
		f = FruitOrange
	case fruitValSeed < 10:
		f = FruitGreen
	}

	return Fruit{
		X:     rand.Intn(boundX),
		Y:     rand.Intn(boundY),
		Value: f,
	}
}
