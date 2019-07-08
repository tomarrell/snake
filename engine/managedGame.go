package engine

import "errors"

// Tick is the snake's velocity at a single
// game state transition
type Tick struct {
	VelX int
	VelY int
}

// ManagedGame is a game where the
// individual ticks are pre-defined
// by the client.
type ManagedGame struct {
	id     int
	Width  int
	Height int
	Snake  Snake
	Fruit  []Fruit
	Score  int
}

func newManagedGame(id, width, height, score int, snake Snake, fruit []Fruit) *ManagedGame {
	return &ManagedGame{
		id,
		width,
		height,
		snake,
		fruit,
		score,
	}
}

func (mg *ManagedGame) run(ticks []Tick) (bool, error) {
	for _, t := range ticks {
		if !validateTick(t) {
			return false, errors.New("one or more ticks are not valid")
		}

		mg.Snake.VelX = t.VelX
		mg.Snake.VelY = t.VelY

		mg.Snake.update()
		if mg.checkSelfCollision() {
			return false, errors.New("collision with self occurs")
		}
	}

	i, ok := mg.checkFruitCollision()
	if !ok {
		return false, errors.New("tick path does not finish on a fruit")
	}

	mg.Snake.eatFruit(mg.Fruit[i].Value)
	mg.Score += int(mg.Fruit[i].Value)
	mg.Fruit[i] = NewFruit(mg.Width, mg.Height)

	return true, nil
}

func (mg *ManagedGame) checkSelfCollision() bool {
	head := mg.Snake.head()

	for _, p := range mg.Snake.Parts[1:] {
		if p.X == head.X && p.Y == head.Y {
			return true
		}
	}

	return false
}

func (mg *ManagedGame) checkFruitCollision() (int, bool) {
	snakeHead := mg.Snake.head()

	for i, fruit := range mg.Fruit {
		if snakeHead.X == fruit.X && snakeHead.Y == fruit.Y {
			return i, true
		}
	}

	return -1, false
}

func validateTick(t Tick) bool {
	sum := t.VelX + t.VelY

	if sum != -1 && sum != 1 {
		return false
	}

	switch t.VelX {
	case -1:
	case 0:
	case 1:
		break
	default:
		return false
	}

	switch t.VelY {
	case -1:
	case 0:
	case 1:
		break
	default:
		return false
	}

	return true
}
