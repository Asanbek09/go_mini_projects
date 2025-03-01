package main

import (
	"gordle2/internal/handlers"
	"net/http"
)

func main() {
	err := http.ListenAndServe(":8080", handlers.Mux())
	if err != nil {
		panic(err)
	}
}