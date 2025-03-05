package gordle

import (
	"math/rand"
	"strings"
	"time"
	"os"
	"fmt"
)

const (
	ErrInaccessibleCorpus = corpusError("corpus can't be opened")
	ErrEmptyCorpus = corpusError("corpus is empty")
	ErrPickRandomWord = corpusError("failed to pick a random word")
)

var corpus string

func ReadCorpus(path string) ([]string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("unable to open %q for reading (%s): %w", path, err, ErrInaccessibleCorpus)
	}

	// we expect the corpus to be a line- or space-separated list of words
	words := strings.Fields(string(data))

	if len(words) == 0 {
		return nil, ErrEmptyCorpus
	}

	return words, nil
}

func pickRandomWord(corpus []string) string {
	rand.Seed(time.Now().UTC().UnixNano())
	index := rand.Intn(len(corpus))

	return corpus[index]
}