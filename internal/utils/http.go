package utils

import (
	"github.com/go-resty/resty/v2"
)

// Get function to send a GET request
func Get(resty *resty.Client, authorizeToken string, body []byte, url string) (*resty.Response, error) {
	resp, err := resty.R().
		SetHeader("Authorization", "Bearer "+authorizeToken).
		SetHeader("Content-Type", "application/json").
		SetBody(body).
		Get(url)

	if err != nil {
		return nil, err
	}

	return resp, nil
}

// Post function to send POST request
func Post(resty *resty.Client, authorizeToken string, body []byte, url string) (*resty.Response, error) {
	resp, err := resty.R().
		SetHeader("Authorization", "Bearer "+authorizeToken).
		SetHeader("Content-Type", "application/json").
		SetBody(body).
		Post(url)

	if err != nil {
		return nil, err
	}

	return resp, nil
}

// Wrapper function to retry the operation
func Retry(operation func() (*resty.Response, error)) (*resty.Response, error) {
	var resp *resty.Response
	var err error

	// Retry logic
	restyError := resty.Backoff(func() (*resty.Response, error) {
		resp, err = operation()
		return resp, err
	})

	return resp, restyError
}
