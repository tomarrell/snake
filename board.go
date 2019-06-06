package main

import (
	"strings"

	tm "github.com/buger/goterm"
)

type board struct {
	width  int
	height int
}

func (b *board) render() {
	block := "\xE2\x96\x88"

	// Header
	tm.Println(block + strings.Repeat(block, b.width) + block)

	for r := 0; r < b.height; r++ {
		tm.Println(block + strings.Repeat(" ", b.width) + block)
	}

	// Footer
	tm.Print(block + strings.Repeat(block, b.width) + block)
}
