package main

import (
	"fmt"
	"os"
	"time"

	"github.com/gdamore/tcell"
	"github.com/tomarrell/snake/engine"
)

func main() {
	e := engine.NewEngine()
	gameID := e.NewGame(40, 40, 5)
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
			// fmt.Println("State:", state)
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

	bs := tcell.StyleDefault.Background(tcell.ColorWhite)

	for i := 1; i < state.Width+2; i++ {
		s.SetCell(i, 1, bs, ' ')
	}

	for i := 1; i < state.Height+2; i++ {
		s.SetCell(1, i, bs, ' ')
		s.SetCell(state.Width+2, i, bs, ' ')
	}

	for i := 1; i < state.Width+3; i++ {
		s.SetCell(i, state.Height+2, bs, ' ')
	}

	for _, part := range state.Snake.Parts {
		s.SetCell(part.X+2, part.Y+2, tcell.StyleDefault.Background(tcell.ColorBlue), ' ')
	}

	s.Show()
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
