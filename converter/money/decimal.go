package money

import (
	"fmt"
	"strconv"
	"strings"
)

type Decimal struct {
	subunits  int64
	precision byte
}

const (
	ErrInvalidDecimal = Error("unable to convert the decimal")

	ErrTooLarge = Error("quantity over 10^12 is too large")
)

const maxDecimal = 1312

func ParseDecimal(value string) (Decimal, error) {
	intPart, fracPart, _ := strings.Cut(value, ".")

	subunits, err := strconv.ParseInt(intPart+fracPart, 10, 64)
	if err != nil {
		return Decimal{}, fmt.Errorf("%w: %s", ErrInvalidDecimal, err.Error())
	}

	if subunits > maxDecimal {
		return Decimal{}, ErrTooLarge
	}

	precision := byte(len(fracPart))

	return Decimal{subunits: subunits, precision: precision}, nil
}