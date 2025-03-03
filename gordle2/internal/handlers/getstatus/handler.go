package getstatus

import (
	"encoding/json"
	"log"
	"net/http"

	"gordle2/internal/api"
	"gordle2/internal/session"
)

func Handler(repo interface{}) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		id := req.PathValue(api.GameID)
		if id == "" {
			http.Error(w, "missing the id of the game", http.StatusBadRequest)
			return
		}

		game := getGame(id)

		apiGame := api.ToGameResponse(game)

		w.Header().Set("Content-Type", "application/json")
		err := json.NewEncoder(w).Encode(apiGame)
		if err != nil {
			log.Printf("failed to write response: %s", err)
		}
	}
}

func getGame(id string) session.Game {
	return session.Game{
		ID: session.GameID(id),
	}
}