package bitflow

import (
	"errors"
	"strings"

	"github.com/bitflowex/go-bitflow/bitflow/types"
)

// Client model.
type Client struct {
	baseURL types.BaseURL
	token   types.Token

	Merchants Merchant
}

// Option is a function that configures a Client.
type Option func(*Client)

// NewClient creates a new Client.
func NewClient(opts ...Option) (*Client, error) {
	// init client
	c := &Client{}
	for _, opt := range opts {
		opt(c)
	}

	// validate base url
	if err := c.baseURL.Validate(); err != nil {
		if errors.Is(err, types.ErrEmptyBaseURL) {
			c.baseURL = types.BaseURL("https://bitflow.ws/api/v1")
		} else {
			return nil, err
		}
	}
	// trim token suffix
	c.baseURL = types.BaseURL(strings.TrimSuffix(c.baseURL.String(), "/"))

	// validate token
	if err := c.token.Validate(); err != nil {
		return nil, err
	}

	// set clients
	c.Merchants = Merchant{c}

	// return client and success
	return c, nil
}

// WithBaseURL sets the base URL for the client.
func WithBaseURL(baseURL types.BaseURL) Option {
	return func(c *Client) {
		c.baseURL = baseURL
	}
}

// WithToken sets the token for the client.
func WithToken(token types.Token) Option {
	return func(c *Client) {
		c.token = token
	}
}
