package engine

import "log"

// Tick is the snake's velocity at a single
// game state transition
type Tick struct {
	VelX int
	VelY int
}

type managedGame struct {
	id     int
	width  int
	height int
	snake  Snake
	fruit  []Fruit
	score  int
}

func newManagedGame(id, width, height, score int, snake Snake, fruit []Fruit) *managedGame {
	return &managedGame{
		id,
		width,
		height,
		snake,
		fruit,
		score,
	}
}

func (mg *managedGame) run(ticks []Tick) bool {
	for _, t := range ticks {
		if !validateTick(t) {
			return false
		}

		mg.snake.velX = t.VelX
		mg.snake.velY = t.VelY

		mg.snake.update()
		log.Println("Ticked snake: ", mg.snake)
	}

	log.Println("Checking collisions", mg.fruit)
	i, ok := mg.checkCollision()
	if !ok {
		return false
	}

	mg.score += int(mg.fruit[i].Value)
	return true
}

func (mg *managedGame) checkCollision() (int, bool) {
	snakeHead := mg.snake.head()
	log.Println("Snake head:", snakeHead)

	for i, fruit := range mg.fruit {
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
