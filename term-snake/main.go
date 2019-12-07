package main

import (
	"flag"
	"fmt"
	"log"
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
	widthPtr := flag.Int("width", 80, "the width of the snake arena")
	heightPtr := flag.Int("height", 40, "the height of the snake arena")

	flag.Parse()

	e := engine.NewEngine()

	var gameID int
	var output chan engine.GameState
	gameID = e.NewGame(*widthPtr, *heightPtr, 10)
	output, err := e.StartGame(gameID, nil)
	if err != nil {
		log.Fatal(err.Error())
	}

	s, err := tcell.NewScreen()
	if err != nil {
		log.Fatalf(err.Error())
	}

	if err = s.Init(); err != nil {
		log.Fatal(err.Error())
	}

	quit := make(chan struct{})
	restart := make(chan struct{})
	initKeyHandlers(s, e, gameID, quit, restart)
	dur := time.Duration(0)

loop:
	for {
		select {
		case state := <-output:
			if state.Finished {
				renderGameOver(s, state)
			} else {
				renderState(s, state)
			}
		case <-quit:
			break loop
		case <-restart:
			e.DestroyGame(gameID)
			gameID = e.NewGame(*widthPtr, *heightPtr, 10)
			output, err = e.StartGame(gameID, nil)
			if err != nil {
				panic(e)
			}
		case <-time.After(time.Millisecond * 50):
		}

		start := time.Now()
		dur += time.Since(start)
	}

	s.Fini()
}

func renderState(s tcell.Screen, state engine.GameState) {
	s.Clear()

	renderOutline(s, state)
	renderSnake(s, state.Snake)
	renderFruit(s, state.Fruit)
	renderText(s, 2, state.Height+offset+1, "Press SPACE to restart.")
	renderText(s, 2, state.Height+offset+2, fmt.Sprint("Score: ", state.Score))

	s.Show()
}

func renderGameOver(s tcell.Screen, state engine.GameState) {
	s.Clear()

	renderOutline(s, state)
	renderSnake(s, state.Snake)
	renderFruit(s, state.Fruit)
	renderText(s, 2, state.Height+offset+2, fmt.Sprint("Score: ", state.Score))
	renderText(s, 2, state.Height+offset+1, "Game Over. Press SPACE to restart.")

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
	for i := inset; i <= state.Width+offset-1; i++ {
		s.SetCell(i, inset, bs, ' ')
	}

	// Let and right borders
	for i := inset + borderWidth; i < state.Height+offset-1; i++ {
		s.SetCell(inset, i, bs, ' ')
		s.SetCell(state.Width+offset-1, i, bs, ' ')
	}

	// Bottom border
	for i := inset; i <= state.Width+offset-1; i++ {
		s.SetCell(i, state.Height+offset-1, bs, ' ')
	}
}

// Render each piece of the snake
func renderSnake(s tcell.Screen, snake engine.Snake) {
	for _, part := range snake.Parts {
		s.SetCell(
			part.X+inset+borderWidth,
			part.Y+inset+borderWidth,
			tcell.StyleDefault.Background(tcell.ColorBlue),
			' ',
		)
	}
}

// Render each fruit
func renderFruit(s tcell.Screen, fruit []engine.Fruit) {
	for _, f := range fruit {
		var style tcell.Style

		switch f.Value {
		case engine.FruitPink:
			style = tcell.StyleDefault.Background(tcell.ColorPink)
		case engine.FruitOrange:
			style = tcell.StyleDefault.Background(tcell.ColorOrange)
		case engine.FruitGreen:
			style = tcell.StyleDefault.Background(tcell.ColorGreen)
		}

		s.SetCell(
			f.X+inset+borderWidth,
			f.Y+inset+borderWidth,
			style,
			' ',
		)
	}
}

func initKeyHandlers(s tcell.Screen, e *engine.Engine, gameID int, c chan struct{}, rst chan struct{}) {
	go func() {
		for {
			ev := s.PollEvent()
			switch ev := ev.(type) {
			case *tcell.EventKey:
				if ev.Rune() == 'q' {
					close(c)
					return
				}
				if ev.Rune() == ' ' {
					rst <- struct{}{}
					continue
				}

				// Special keys
				switch ev.Key() {
				case tcell.KeyEscape, tcell.KeyCtrlC:
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
