package ecbank

import (
	"converter/money"
	"fmt"
)

type envelope struct {
	Rates []currencyRate `xml:"Cube>Cube>Cube"`
}

type currencyRate struct {
	Currency string `xml:"currency,attr"`
	Rate money.ExchangeRate `xml:"rate,attr"`
}

const baseCurrencyCode = "EUR"

func (e envelope) exchangeRates() map[string]money.ExchangeRate {
	rates := make(map[string]money.ExchangeRate, len(e.Rates)+1)

	for _, c := range e.Rates {
		rates[c.Currency] = c.Rate
	}

	rates[baseCurrencyCode] = 1.

	return rates
}

func (e envelope) exchangeRate(source, target string) (money.ExchangeRate, error) {
	if source == target {
		return 1., nil
	}

	rates := e.mappedChangeRates()

	sourceFactor, sourceFound := rates[source]
	if !sourceFound {
		return 0, fmt.Errorf("failed to find the source currency %s", source)
	}

	targetFactor, targetFound := rates[target]
	if !targetFound {
		return 0, fmt.Errorf("failed to find target currency %s", target)
	}

	return targetFactor / sourceFactor, nil
}