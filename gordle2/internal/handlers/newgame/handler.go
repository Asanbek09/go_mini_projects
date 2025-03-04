package newgame

import (
	"encoding/json"
	"fmt"
	"gordle2/internal/api"
	"gordle2/internal/session"
	"log"
	"net/http"
)

type gameAdder interface {
	Add(game session.Game) error
}

func Handler(adder gameAdder) http.HandlerFunc {
	return func(w http.ResponseWriter, _ *http.Request) {
		game := createGame(adder)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)

		apiGame := api.ToGameResponse(game)
		err := json.NewEncoder(w).Encode(apiGame)
		if err != nil {
			log.Printf("failed to write response: %s", err)
		}
	}
}

func createGame(db gameAdder) (session.Game, error) {
	corpus, err := gordle2.ReadCorpus("corpus/english.txt")
	if err != nil {
		return session.Game{}, fmt.Errorf("unable to read corpus: %w", err)
	}

	game, err := gordle2.New(corpus)
	if err != nil {
		return session.Game{}, fmt.Errorf("failed to create a new gordle game")
	}

	g := session.Game{
		ID: session.GameID(ulid.Make().String()),
		Gordle: *game,
		AttemptsLeft: maxAttempts,
		Guesses: []session.Guess{},
		Status: session.StatusPlaying,
	}

	err = db.Add(g)
	if err != nil {
		return session.Game{}, fmt.Errorf("failed to save the new game")
	}

	return g, nil
}