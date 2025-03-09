package habit

import (
	"testing"
)

func Test_validateAndFillDetails(t testing.T) {
	t.Parallel()

	t.Run("Full", testValidateAndFillDetailsFull)
	t.Run("Partial", testValidateAndFillDetailsPartial)
	t.Run("SpaceName", testValidateAndFillDetailsSpaceName)
}