package engine

type KeyCode int

const (
	KeyLeft KeyCode = iota
	KeyRight
	KeyUp
	KeyDown
)

type Game struct {
	ID        int
	Tickrate  int
	Width     int
	Height    int
	inputChan chan (Command)
}
