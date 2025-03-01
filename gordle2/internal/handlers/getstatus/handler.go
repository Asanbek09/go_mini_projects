package getstatus

import (
	"encoding/json"
	"log"
	"net/http"

	"gordle2/internal/api"
)

func Handle(w http.ResponseWriter, req *http.Request) {
	id := req.PathValue(api.GameID)
	if id == "" {
		http.Error(w, "missing the id of the game", http.StatusBadRequest)
		return
	}

	apiGame := api.GameResponse{
		ID: id,
	}

	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(apiGame)
	if err != nil {
		log.Printf("failed to write response: %s", err)
	}
}