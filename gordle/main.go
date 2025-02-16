package main

import (
	"os"
    "gordle/gordle_packs"
)

func main() {
    g := gordlepacks.New(os.Stdin)
    g.Play()
}