package handlers

import (
	"gordle2/internal/api"
	"gordle2/internal/handlers/newgame"
	"net/http"
)


func Mux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc(api.NewGameRoute, newgame.Handle)
	return mux
}

func NewRouter() *http.ServeMux {
	r := http.NewServeMux()
	r.HandleFunc(http.MethodPost + " " + api.NewGameRoute, newgame.Handle)

	return r
}