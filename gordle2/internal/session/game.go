package session

import "errors"

type Game struct {
	ID GameID
	Attemptsleft byte
	Guesses []Guess
	Status Status
}

type GameID string

type Status string

const (
	StatusPlaying = "Playing"
	StatusWon = "Won"
	StatusLost = "Lost"
)

type Guess struct {
	Word string
	Feedback string
}

var ErrGameOver = errors.New("Game Over")