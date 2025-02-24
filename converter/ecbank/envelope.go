package ecbank

import "converter/money"

type envelope struct {
	Rates []currencyRate `xml:"Cube>Cube>Cube"`
}

type currencyRate struct {
	Currency string `xml:"currency,attr"`
	Rate money.ExchangeRate `xml:"rate,attr"`
}