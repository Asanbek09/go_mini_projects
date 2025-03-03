package newgame

import (
	"encoding/json"
	"gordle2/internal/api"
	"gordle2/internal/session"
	"log"
	"net/http"
)

func Handle(w http.ResponseWriter, req *http.Request) {
	
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	apiGame := api.ToGameResponse(game)
	err := json.NewEncoder(w).Encode(apiGame)
	if err != nil {
		log.Printf("failed to write response: %s", err)
	}
}

func createGame() session.Game {
	return session.Game{}
}