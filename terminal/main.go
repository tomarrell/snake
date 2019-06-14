package main

import (
	"github.com/tomarrell/snake/engine"
)

func main() {
	e := engine.NewEngine()

	err := e.StartGame(e.NewGame(100, 100, 5))
	if err != nil {
		panic(e)
	}

	e.SendInput(0, engine.KeyRight)
}
