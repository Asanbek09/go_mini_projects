package newgame

import (
	"encoding/json"
	"fmt"
	"gordle2/internal/api"
	"gordle2/internal/session"
	"gordle2/internal/gordle"
	"log"
	"net/http"
	"github.com/oklog/ulid/v2"
)

type gameAdder interface {
	Add(game session.Game) error
}

var corpora = map[string]string{
	"en": "./../../../corpus/english.txt",
	"he": "./../../../corpus/greek.txt",
	"cr": "./../../../corpus/cree.txt",
}

func Handler(adder gameAdder) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		lang := r.URL.Query().Get(api.Lang)
		corpusPath, ok := corpora[lang]
		if !ok {
			corpusPath = corpora["en"]
		}

		game, err := createGame(adder, corpusPath)
		if err != nil {
			log.Printf("unable to create a new game: %s", err)
			http.Error(w, "failed to create a new game", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)

		apiGame := api.ToGameResponse(game)
		err = json.NewEncoder(w).Encode(apiGame)
		if err != nil {
			log.Printf("failed to write response: %s", err)
		}
	}
}

const maxAttempts = 5

func createGame(adder gameAdder, corpusPath string) (session.Game, error) {
	corpus, err := gordle.ReadCorpus(corpusPath)
	if err != nil {
		return session.Game{}, fmt.Errorf("unable to read corpus: %w", err)
	}

	game, err := gordle.New(corpus)
	if err != nil {
		return session.Game{}, fmt.Errorf("failed to create a new gordle game")
	}

	g := session.Game{
		ID:           session.GameID(ulid.Make().String()),
		Gordle:       *game,
		AttemptsLeft: maxAttempts,
		Guesses:      []session.Guess{},
		Status:       session.StatusPlaying,
	}

	err = adder.Add(g)
	if err != nil {
		return session.Game{}, fmt.Errorf("failed to save the new game")
	}

	return g, nil
}