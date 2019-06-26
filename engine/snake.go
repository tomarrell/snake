package engine

// Snake is... the snake.
type Snake struct {
	Parts  []Part `json:"parts"`
	boundX int
	boundY int
	velX   int
	velY   int
}

// Part is a single piece of the snake
type Part struct {
	X int `json:"x"`
	Y int `json:"y"`
}

func newSnake(boundX, boundY int) Snake {
	// Start a new snake with 3 parts facing East
	return Snake{
		[]Part{
			Part{3, 1},
			Part{2, 1},
			Part{1, 1},
		},
		boundX,
		boundY,
		1,
		0,
	}
}

func (s *Snake) length() int {
	return len(s.Parts)
}

func (s *Snake) head() Part {
	return s.Parts[0]
}

func (s *Snake) eatFruit(size FruitValue) int {
	lastPart := s.Parts[len(s.Parts)-1]

	for i := 0; i < int(size); i++ {
		s.Parts = append(s.Parts, lastPart)
	}

	return len(s.Parts)
}

func (s *Snake) up() {
	if s.velY != 0 {
		return
	}

	s.velX = 0
	s.velY = -1
}

func (s *Snake) down() {
	if s.velY != 0 {
		return
	}

	s.velX = 0
	s.velY = 1
}

func (s *Snake) left() {
	if s.velX != 0 {
		return
	}

	s.velX = -1
	s.velY = 0
}

func (s *Snake) right() {
	if s.velX != 0 {
		return
	}

	s.velX = 1
	s.velY = 0
}

func (s *Snake) update() {
	newHead := s.Parts[0]

	newHead.X += s.velX
	newHead.Y += s.velY

	switch {
	case newHead.X < 0:
		newHead.X = s.boundX - 1
	case newHead.X >= s.boundX:
		newHead.X = 0
	case newHead.Y < 0:
		newHead.Y = s.boundY - 1
	case newHead.Y >= s.boundY:
		newHead.Y = 0
	}

	s.Parts = append([]Part{newHead}, s.Parts[:len(s.Parts)-1]...)
}
