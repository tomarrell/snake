package main

import (
	"github.com/tomarrell/snake/engine"
)

func main() {
	e := engine.NewEngine()

	_, err := e.StartGame(e.NewGame(100, 100, 5))
	if err != nil {
		panic(e)
	}

	_, err = e.StartGame(e.NewGame(100, 100, 10))
	if err != nil {
		panic(e)
	}

	e.EndGame(1)

	e.Start()
}
