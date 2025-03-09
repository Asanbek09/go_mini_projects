package habit

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_validateAndFillDetails(t testing.T) {
	t.Parallel()

	t.Run("Full", testValidateAndFillDetailsFull)
	t.Run("Partial", testValidateAndFillDetailsPartial)
	t.Run("SpaceName", testValidateAndFillDetailsSpaceName)
}

func testValidateAndFillDetailsFull(t testing.T) {
	t.Parallel()

	h := Habit{
		ID: "987",
		Name: "laugh",
		WeeklyFrequency: 256,
		CreationTime: time.Date(2024, 03, 9, 1, 5, 0, 0, time.UTC),
	}

	got, err := validateAndFillDetails(h)
	require.NoError(t, err)
	assert.Equal(t, h, got)
}

func testValidateAndFillDetailsPartial(t *testing.T) {
	t.Parallel()

	h := Habit{
		Name:            "laugh",
		WeeklyFrequency: 256,
	}

	got, err := validateAndFillDetails(h)
	require.NoError(t, err)
	assert.Equal(t, h.Name, got.Name)
	assert.Equal(t, h.WeeklyFrequency, got.WeeklyFrequency)
	assert.NotEmpty(t, got.ID)
	assert.NotEmpty(t, got.CreationTime)
}