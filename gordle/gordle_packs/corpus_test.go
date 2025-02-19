package gordlepacks_test

import (
	gordlepacks "gordle/gordle_packs"
	"testing"
)

func TestReadCorpus(t *testing.T) {
	tt := map[string]struct {
		file string
		length int
		err error
	}{
		"Englsih corpus": {
			file: "../corpus/english.txt",
			length: 35,
			err: nil,
		},
		"empty corpus": {
			file: "../corpus/empty.txt",
			length: 0,
			err: gordlepacks.ErrCorpusIsEmpty,
		},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			words, err := gordlepacks.ReadCorpus(tc.file)
			if tc.err != err {
				t.Errorf("expected err %v, got %v", tc.err, err)
			}

			if tc.length != len(words) {
				t.Errorf("expected %d, got %d", tc.length, len(words))
			}
		})
	}
}