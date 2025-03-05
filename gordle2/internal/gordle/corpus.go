package gordle

import (
	"math/rand"
	"strings"
	"time"
)

const (

	ErrEmptyCorpus = corpusError("corpus is empty")
	ErrPickRandomWord = corpusError("failed to pick a random word")
)

var corpus string

func ParseCorpus() ([]string, error) {
	words := strings.Fields(corpus)

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