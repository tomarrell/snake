package main

import (
	"fmt"
	"os"
	"time"

	"github.com/gdamore/tcell"
	"github.com/tomarrell/snake/engine"
)

const (
	inset       = 1
	borderWidth = 1
	offset      = inset + (2 * borderWidth)
)

func main() {
	e := engine.NewEngine()
	gameID := e.NewGame(80, 40, 5)
	output, err := e.StartGame(gameID)
	if err != nil {
		panic(e)
	}

	s, err := tcell.NewScreen()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", e)
		os.Exit(1)
	}

	if err = s.Init(); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", e)
		os.Exit(1)
	}

	quit := make(chan struct{})
	initKeyHandlers(s, e, gameID, quit)

	dur := time.Duration(0)

loop:
	for {
		select {
		case state := <-output:
			renderState(s, state)
		case <-quit:
			break loop
		case <-time.After(time.Millisecond * 50):
		}
		start := time.Now()
		dur += time.Now().Sub(start)
	}

	s.Fini()
}

func renderState(s tcell.Screen, state engine.GameState) {
	s.Clear()

	renderOutline(s, state)
	renderSnake(s, state.Snake.Parts)
	renderText(s, 2, 45, fmt.Sprint("Score: ", state.Score))

	s.Show()
}

// Render a string of text starting at x, y
func renderText(s tcell.Screen, x, y int, text string) {
	for i, r := range text {
		s.SetCell(x+i, y, tcell.StyleDefault, r)
	}
}

// Render the arena
func renderOutline(s tcell.Screen, state engine.GameState) {
	bs := tcell.StyleDefault.Background(tcell.ColorWhite)

	// Top border
	for i := inset; i <= state.Width+offset; i++ {
		s.SetCell(i, inset, bs, ' ')
	}

	// Let and right borders
	for i := inset + borderWidth; i < state.Height+offset; i++ {
		s.SetCell(inset, i, bs, ' ')
		s.SetCell(state.Width+offset, i, bs, ' ')
	}

	// Bottom border
	for i := inset; i <= state.Width+offset; i++ {
		s.SetCell(i, state.Height+offset, bs, ' ')
	}
}

// Render the snake on screen
func renderSnake(s tcell.Screen, snake []engine.Part) {
	for _, part := range snake {
		s.SetCell(
			part.X+inset+borderWidth,
			part.Y+inset+borderWidth,
			tcell.StyleDefault.Background(tcell.ColorBlue),
			' ',
		)
	}
}

func initKeyHandlers(s tcell.Screen, e *engine.Engine, gameID int, c chan (struct{})) {
	go func() {
		for {
			ev := s.PollEvent()
			switch ev := ev.(type) {
			case *tcell.EventKey:
				switch ev.Key() {
				case tcell.KeyEscape, tcell.KeyEnter, tcell.KeyCtrlC:
					close(c)
					return
				case tcell.KeyDown:
					e.SendInput(gameID, engine.KeyDown)
				case tcell.KeyRight:
					e.SendInput(gameID, engine.KeyRight)
				case tcell.KeyUp:
					e.SendInput(gameID, engine.KeyUp)
				case tcell.KeyLeft:
					e.SendInput(gameID, engine.KeyLeft)
				case tcell.KeyCtrlL:
					s.Sync()
				}
			case *tcell.EventResize:
				s.Sync()
			}
		}
	}()
}
