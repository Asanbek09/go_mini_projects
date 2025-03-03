package newgame

import (
	"encoding/json"
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

func createGame(db gameAdder) session.Game {
	return session.Game{}
}