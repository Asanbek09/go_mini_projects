package gordlepacks

import "testing"

func inCorpus(corpus []string, word string) bool {
	for _, corpusWord := range corpus {
		if corpusWord == word {
			return true
		}
	}

	return false
}

func TestPickWord(t *testing.T) {
	corpus := []string{"HELLO", "SALUT", "ПРИВЕТ", "XAIPE"}
	word := pickWord(corpus)

	if !inCorpus(corpus, word) {
		t.Errorf("Expected a word in the corpus, got %q", word)
	}
}