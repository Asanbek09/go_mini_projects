package session

import (
	"errors"
	"gordle2/internal/gordle"
)

type Game struct {
	ID           GameID
	Gordle       gordle.Game
	AttemptsLeft byte
	Guesses      []Guess
	Status       Status
}

type GameID string

type Status string

const (
	StatusPlaying = "Playing"
	StatusWon     = "Won"
	StatusLost    = "Lost"
)

type Guess struct {
	Word     string
	Feedback string
}

var ErrGameOver = errors.New("game over")