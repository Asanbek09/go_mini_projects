package session

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