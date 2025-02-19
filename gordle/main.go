package main

import (
	"bufio"
	"fmt"
	"gordle/gordle_packs"
	"os"
)

const maxAttempts = 6

func main() {
    corpus, err := gordlepacks.ReadCorpus("corpus/english.txt")
    if err != nil {
        _, _ = fmt.Fprintf(os.Stderr, "unable to read corpus: %s", err)
        return
    }

    g, err := gordlepacks.New(bufio.NewReader(os.Stdin), corpus, maxAttempts)
    if err != nil {
        _, _ = fmt.Fprintf(os.Stderr, "unable to start game: %s", err)
        return
    }

    g.Play()
}