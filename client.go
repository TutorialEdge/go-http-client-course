package client

import (
	"net/http"
	"time"
)

// Client - Our magical Client
type Client struct {
	httpClient *http.Client
}

// Option - a handy type definition that will
// save a bit of repetitive typing
type Option func(c *Client)

// NewClient - returns a new instance of our http client
func NewClient(
	opts ...Option,
) *Client {
	client := &Client{
		httpClient: http.DefaultClient,
	}

	for _, o := range opts {
		o(client)
	}

	return client
}

// WithTimeout - Allows us to define custom timeouts for
// our HTTP Client.
func WithTimeout(timeout time.Duration) Option {
	return func(c *Client) {
		c.httpClient.Timeout = timeout
	}
}

// WithHttpClient - Maybe we want incredibly specific configuration
// for our http client
func WithHttpClient(httpClient *http.Client) Option {
	return func(c *Client) {
		c.httpClient = httpClient
	}
}
