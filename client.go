package gobitflow

import "strings"

type Client struct {
	baseURL string
	token   string
}

func NewClient(baseURL, token string) *Client {
	// trim baseURL suffix
	baseURL = strings.TrimSuffix(baseURL, "/")

	// create client
	return &Client{
		baseURL: baseURL,
		token:   token,
	}
}
