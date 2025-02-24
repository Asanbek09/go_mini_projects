package ecbank

import "converter/money"

type Client struct{}

func (c Client) FetchExchangeRate(source, target money.Currency) (money.ExchangeRate, error) {
	return money.ExchangeRate{}, nil
}