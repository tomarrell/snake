package engine

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

func (mg *managedGame) run() bool {
	return true
}
