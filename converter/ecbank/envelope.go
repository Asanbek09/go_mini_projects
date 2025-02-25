package ecbank

import "converter/money"

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