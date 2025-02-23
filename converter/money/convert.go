package money

import "math"

func Convert(amount Amount, to Currency) (Amount, error) {
	return Amount{}, nil
}

type ExchangeRate Decimal

func pow10(power byte) int64 {
	switch power {
	case 0:
		return 1
	case 1:
		return 10
	case 2:
		return 100
	case 3:
		return 1000
	default:
		return int64(math.Pow(10, float64(power)))
	}
}