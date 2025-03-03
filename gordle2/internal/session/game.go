package session

type Game struct {
	ID GameID
	Attemptsleft byte
	Guesses []Guess
	Status Status
}