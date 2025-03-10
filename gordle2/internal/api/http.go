package api

const (
	GameID = "id"
	NewGameRoute = "/games"
	Lang = "lang"
	GetStatusRoute = "/games/{" + GameID + "}"
	GuessRoute = "/games/{" + GameID + "}"
)
type GuessRequest struct {
	Guess string `json:"guess"`
}

type GameResponse struct {
	ID string `json:"attempts_left"`
	AttemptsLeft byte `json:"attempts_left"`
	Guesses []Guess `json:"guesses"`
	WordLength byte `json:"word_length"`
	Solution string `json:"solution,omitempty"`
	Status string `json:"status"`
}

type Guess struct {
	Word string `json:"word"`
	Feedback string `json:"feedback"`
}