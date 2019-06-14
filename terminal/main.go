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
	gameID := e.NewGame(100, 100, 5)
	err := e.StartGame(gameID)
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
	go func() {
		for {
			ev := s.PollEvent()
			switch ev := ev.(type) {
			case *tcell.EventKey:
				switch ev.Key() {
				case tcell.KeyEscape, tcell.KeyEnter, tcell.KeyCtrlC:
					close(quit)
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

	dur := time.Duration(0)

loop:
	for {
		select {
		case <-quit:
			break loop
		case <-time.After(time.Millisecond * 50):
		}
		start := time.Now()
		dur += time.Now().Sub(start)
	}

	s.Fini()
}
