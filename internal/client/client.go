package client

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"
)

// API contains client, url and apikey
type API struct {
	Client  *http.Client
	BaseURL string
}

const (
	DefaultTimeout      = 10 * time.Second
	MaxIdleConnsPerHost = 50
	MaxConnsPerHost     = 50
)

type Client struct {
	HttpClient *http.Client
}

func New() *Client {
	t := http.DefaultTransport.(*http.Transport).Clone()
	t.MaxIdleConnsPerHost = MaxIdleConnsPerHost
	t.MaxConnsPerHost = MaxConnsPerHost

	newClient := &http.Client{
		Timeout:   DefaultTimeout,
		Transport: t,
	}

	return &Client{
		HttpClient: newClient,
	}
}

func (c *Client) Get() (*Client, error) {
	if c.HttpClient == nil {
		return nil, errors.New("http client not available")
	}

	return c, nil
}

func (c *Client) SetHeader(token string, request *http.Request) *http.Request {
	request.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
	request.Header.Set("Content-type", "application/json")

	return request
}

// RoundTripFunc .
type RoundTripFunc func(req *http.Request) *http.Response

// RoundTrip .
func (f RoundTripFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return f(req), nil
}

// TestClient returns *http.Client with Transport replaced to avoid making real calls
func TestClient(fn RoundTripFunc) *http.Client {
	return &http.Client{
		Transport: fn,
	}
}

// GetTestClient returns a TestClient
func GetTestClient(statusCode int, response string) *Client {
	var client = TestClient(func(req *http.Request) *http.Response {
		// Test request parameters
		response := &http.Response{
			StatusCode: statusCode,
			// Send response to be tested
			Body: io.NopCloser(bytes.NewBufferString(response)),
			// Must be set to non-nil value or it panics
			Header: make(http.Header),
		}
		response.Header.Set("Content-Type", "application/json")
		return response
	})

	return &Client{
		HttpClient: client,
	}
}
