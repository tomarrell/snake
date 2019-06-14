package engine

import "math/rand"

type fruitValue int

const (
	fruitPink   fruitValue = 1
	fruitOrange fruitValue = 2
	fruitGreen  fruitValue = 5
)

type fruit struct {
	value fruitValue
	x     int
	y     int
}

func newFruit(boundX, boundY int) fruit {
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

	return fruit{
		x:     rand.Intn(boundX + 1),
		y:     rand.Intn(boundY + 1),
		value: f,
	}
}
