package guess

import (
	"gordle2/internal/api"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)


func TestHandle(t *testing.T) {
	handle := Handler(nil)

	req, err := http.NewRequest(http.MethodPost, "/games/", strings.NewReader(`{"guess":"pocket"}`))
	require.NoError(t, err)

	// add path parameters
	req.SetPathValue(api.GameID, "123456")

	recorder := httptest.NewRecorder()

	handle(recorder, req)

	assert.Equal(t, http.StatusOK, recorder.Code)
	assert.Equal(t, "application/json", recorder.Header().Get("Content-Type"))
	assert.JSONEq(t, `{"id":"123456","attempts_left":0,"guesses":[],"word_length":0,"status":""}`, recorder.Body.String())
}