package main

import (
	"os"
    "gordle/gordle_packs"
)

const maxAttempts = 6

func main() {
    solution := "Hello"

    g := gordlepacks.New(os.Stdin, solution, maxAttempts)
    g.Play()
}