package handlers

import (
	"gordle2/internal/api"
	"gordle2/internal/handlers/getstatus"
	"gordle2/internal/handlers/guess"
	"gordle2/internal/handlers/newgame"
	"gordle2/internal/repository"
	"net/http"
)


func NewRouter(db *repository.GameRepository) *http.ServeMux {
	r := http.NewServeMux()

	r.HandleFunc(http.MethodPost + " " + api.NewGameRoute, newgame.Handler(db))
	r.HandleFunc(http.MethodGet + " " + api.GetStatusRoute, getstatus.Handler(db))
	r.HandleFunc(http.MethodPut + " " + api.GuessRoute, guess.Handler(db))

	return r
}