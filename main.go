package main

import (
	"bufio"
	"os"
	"time"

	_ "./engine"

	tm "github.com/buger/goterm"
)

const (
	// WALL_WIDTH is The non-playable offset that is built
	// into the coordinate system
	WallWidth = 1
)

type fruit struct {
	x int
	y int
}

func main() {
	h, _ := height()
	w, _ := width()

	board := board{
		width:  int(w) - 2*WallWidth,
		height: int(h) - 2*WallWidth,
	}

	snake := snake{
		parts: []part{
			part{2, 3},
			part{3, 3},
			part{4, 3},
		},
	}

	tm.Clear()

	for {
		tm.MoveCursor(1, 1)

		board.render()
		snake.render()

		tm.Flush()
		time.Sleep(time.Second)
	}
}

var output = bufio.NewWriter(os.Stdout)
