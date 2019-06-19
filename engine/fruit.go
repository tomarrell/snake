package engine

import "math/rand"

type fruitValue int

const (
	fruitPink   fruitValue = 1
	fruitOrange fruitValue = 2
	fruitGreen  fruitValue = 5
)

// Fruit is the item which adds score and
// length when eaten by the snake
type Fruit struct {
	Value fruitValue
	X     int
	Y     int
}

func newFruit(boundX, boundY int) Fruit {
	fruitValSeed := rand.Intn(10)
	var f fruitValue

	switch {
	case fruitValSeed < 5:
		f = fruitPink
	case fruitValSeed < 8:
		f = fruitOrange
	case fruitValSeed < 10:
		f = fruitGreen
	}

	return Fruit{
		X:     rand.Intn(boundX + 1),
		Y:     rand.Intn(boundY + 1),
		Value: f,
	}
}
