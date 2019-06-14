package engine

type snake struct {
	parts  []part
	boundX int
	boundY int
	velX   int
	velY   int
}

type part struct {
	x int
	y int
}

// NewSnake creates a new snake
func newSnake(boundX, boundY int) snake {
	// Start a new snake with 3 parts facing East
	return snake{
		[]part{
			part{3, 1},
			part{2, 1},
			part{1, 1},
		},
		boundX,
		boundY,
		1,
		0,
	}
}

func (s *snake) length() int {
	return len(s.parts)
}

func (s *snake) head() part {
	return s.parts[0]
}

func (s *snake) eatFruit(size fruitValue) int {
	lastPart := s.parts[len(s.parts)-1]

	for i := 0; i < int(size); i++ {
		s.parts = append(s.parts, lastPart)
	}

	return len(s.parts)
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
	newHead := s.parts[0]

	newHead.x += s.velX
	newHead.y += s.velY

	switch {
	case newHead.x < 0:
		newHead.x = s.boundX
	case newHead.x > s.boundX:
		newHead.x = 0
	case newHead.y < 0:
		newHead.y = s.boundY
	case newHead.y > s.boundY:
		newHead.y = 0
	}

	s.parts = append([]part{newHead}, s.parts[:len(s.parts)-1]...)
}
