package newgame

import (
	"gordle2/internal/session"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestHandle(t *testing.T) {
	handleFunc := Handler(gameAdderStub{})

	req, err := http.NewRequest(http.MethodPost, "/games", nil)
	require.NoError(t, err)

	recorder := httptest.NewRecorder()

	handleFunc(recorder, req)

	assert.Equal(t, http.StatusCreated, recorder.Code)
	assert.Equal(t, "application/json", recorder.Header().Get("Content-Type"))
	assert.JSONEq(t, `{"id":"","attempts_left":0,"guesses":[],"word_length":0,"status":""}`, recorder.Body.String())
}

func Test_createGame(t *testing.T) {
	corpusPath  = "testdata/corpus.txt"

	g, err := createGame(gameCreatorStub{nil})
	require.NoError(t, err)

	assert.Regexp(t, "[A-Z0-9]+", g.ID)
	assert.Equal(t, uint8(5), g.AttemptsLeft)
	assert.Equal(t, 0, len(g.Guesses))
}

type gameAdderStub struct {
	err error
}

func (g gameAdderStub) Add(_ session.Game) error {
	return g.err
}