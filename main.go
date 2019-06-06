package main

import (
	"fmt"
)

type board struct {
	width  int
	height int
}

type snake struct {
	x      int
	y      int
	length int
}

type fruit struct {
	x int
	y int
}

func main() {
	fmt.Println("Hello World")
}
