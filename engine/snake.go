package engine

import (
	tm "github.com/buger/goterm"
)

type snake struct {
	parts []part
}

type part struct {
	x int
	y int
}

func (p *part) renderPart() {
	tm.MoveCursor(
		p.x,
		p.y,
	)
	tm.Print("\xe2\x83\x9d")
}

func (s *snake) render() {
	for _, p := range s.parts {
		p.renderPart()
	}
}
