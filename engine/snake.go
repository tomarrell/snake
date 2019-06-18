package engine

type snake struct {
	Parts  []Part
	boundX int
	boundY int
	velX   int
	velY   int
}

type Part struct {
	X int
	Y int
}

// NewSnake creates a new snake
func newSnake(boundX, boundY int) snake {
	// Start a new snake with 3 parts facing East
	return snake{
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

func (s *snake) length() int {
	return len(s.Parts)
}

func (s *snake) head() Part {
	return s.Parts[0]
}

func (s *snake) eatFruit(size fruitValue) int {
	lastPart := s.Parts[len(s.Parts)-1]

	for i := 0; i < int(size); i++ {
		s.Parts = append(s.Parts, lastPart)
	}

	return len(s.Parts)
}

func (s *snake) up() {
	if s.velY != 0 {
		return
	}

	s.velX = 0
	s.velY = -1
}

func (s *snake) down() {
	if s.velY != 0 {
		return
	}

	s.velX = 0
	s.velY = 1
}

func (s *snake) left() {
	if s.velX != 0 {
		return
	}

	s.velX = -1
	s.velY = 0
}

func (s *snake) right() {
	if s.velX != 0 {
		return
	}

	s.velX = 1
	s.velY = 0
}

func (s *snake) update() {
	newHead := s.Parts[0]

	newHead.X += s.velX
	newHead.Y += s.velY

	switch {
	case newHead.X < 0:
		newHead.X = s.boundX
	case newHead.X > s.boundX:
		newHead.X = 0
	case newHead.Y < 0:
		newHead.Y = s.boundY
	case newHead.Y > s.boundY:
		newHead.Y = 0
	}

	s.Parts = append([]Part{newHead}, s.Parts[:len(s.Parts)-1]...)
}
