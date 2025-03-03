package main

import (
	"gordle2/internal/handlers"
	"gordle2/internal/repository"
	"net/http"
)

func main() {
	db := repository.New()

	err := http.ListenAndServe(":8080", handlers.NewRouter(db))
	if err != nil {
		panic(err)
	}
}