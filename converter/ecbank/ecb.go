package ecbank

import (
	"converter/money"
	"net/http"
	"fmt"
)

const (
	ErrCallingServer = ecbankError("error calling server")
)

type Client struct{
	url string
}

func (c Client) FetchExchangeRate(source, target money.Currency) (money.ExchangeRate, error) {
	const euroxrefURL = "http://www.ecb.europa.eu/stats/eurofxref/eurofxref-daily.xml"

	if c.url == "" {
		c.url = euroxrefURL
	}

	resp, err := http.Get(c.url)
	if err != nil {
		return money.ExchangeRate{}, fmt.Errorf("%w: %s", ErrCallingServer, err.Error())
	}

	defer resp.Body.Close()

	return money.ExchangeRate{}, nil
}



