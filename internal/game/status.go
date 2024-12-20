package game

type Status int

const (
    StatusMenu Status = iota
	StatusPlaying
	StatusPaused
	StatusGameOver
	StatusWon
)
