package money

import (
	"fmt"
	"math"
)

func Convert(amount Amount, to Currency, rates exchangeRates) (Amount, error) {
	r, err := rates.FetchExchangeRate(amount.currency, to)
	if err != nil {
		return Amount{}, fmt.Errorf("cannot get change rate: %w", err)
	}

	convertedValue := applyExchangeRate(amount, to, r)

	if err := convertedValue.validate(); err != nil {
		return Amount{}, err
	}

	return convertedValue, nil
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

func applyExchangeRate(a Amount, target Currency, rate ExchangeRate) Amount {
	converted := multiply(a.quantity, rate)

	switch {
	case converted.precision > target.precision:
		converted.subunits = converted.subunits / pow10(converted.precision-target.precision)
	case converted.precision < target.precision:
		converted.subunits = converted.subunits * pow10(target.precision-a.currency.precision)
	}

	converted.precision = target.precision

	return Amount{
		currency: target,
		quantity: converted,
	}
}

func multiply(d Decimal, r ExchangeRate) Decimal {
	return Decimal{
		subunits: d.subunits * r.subunits,
		precision: d.precision + r.precision,
	}
}