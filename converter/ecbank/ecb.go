package ecbank

import (
	"converter/money"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"time"
)

const (
	ErrCallingServer      = ecbankError("error calling server")
	ErrTimeout            = ecbankError("timed out when waiting for response")
	ErrUnexpectedFormat   = ecbankError("unexpected response format")
	ErrChangeRateNotFound = ecbankError("couldn't find the exchange rate")
	ErrClientSide         = ecbankError("client side error when contacting ECB")
	ErrServerSide         = ecbankError("server side error when contacting ECB")
	ErrUnknownStatusCode  = ecbankError("unknown status code contacting ECB")
)

type Client struct{
	client *http.Client
}

func NewClient(timeout time.Duration) Client {
	return Client{
		client: &http.Client{Timeout: timeout},
	}
}

func (c Client) FetchExchangeRate(source, target money.Currency) (money.ExchangeRate, error) {
	const euroxrefURL = "http://www.ecb.europa.eu/stats/eurofxref/eurofxref-daily.xml"

	resp, err := c.client.Get(euroxrefURL)
	if err != nil {
		var urlErr *url.Error
		if ok := errors.As(err, &urlErr); ok && urlErr.Timeout() {
			return money.ExchangeRate{}, fmt.Errorf("%w: %s", ErrTimeout, err.Error())
		}
	}

	if err = checkStatusCode(resp.StatusCode); err != nil {
		return money.ExchangeRate{}, err
	}

	rate, err := readRateFromResponse(source.Code(), target.Code(), resp.Body)
	if err != nil {
		return money.ExchangeRate{}, err
	}

	return rate, nil
}

const (
	clientErrorClass = 4
	serverErrorClass = 5
)

func checkStatusCode(statusCode int) error {
	switch {
	case statusCode == http.StatusOK:
		return nil
	case httpStatusClass(statusCode) == clientErrorClass:
		return fmt.Errorf("%w: %d", ErrClientSide, statusCode)
	case httpStatusClass(statusCode) == serverErrorClass:
		return fmt.Errorf("%w: %d", ErrServerSide, statusCode)
	default:
		return fmt.Errorf("%w: %d", ErrUnknownStatusCode, statusCode)
	}
}

func httpStatusClass(statusCode int) int {
	const httpErrorClassSize = 100
	return statusCode / httpErrorClassSize
}