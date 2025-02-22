package money

type Decimal struct {
	subunits  int64
	precision byte
}

const (
	ErrInvalidDecimal = Error("unable to convert the decimal")

	ErrTooLarge = Error("quantity over 10^12 is too large")
)
