package types

import (
	"net/url"

	"github.com/tyrenix/goerr"
)

// BaseURL represents a base URL.
type BaseURL string

// String returns the string representation of the base URL.
func (b BaseURL) String() string {
	return string(b)
}

// Validate validates the base URL.
func (b BaseURL) Validate() error {
	// if base url is empty
	if b == "" {
		return ErrEmptyBaseURL
	}

	// parse base url
	url, err := url.Parse(b.String())
	if err != nil {
		return goerr.New(ErrInvalidBaseURL, "failed parse url", err)
	}

	// check is absolute
	if !url.IsAbs() {
		return ErrNotAbsoluteBaseURL
	}

	// check scheme
	if url.Scheme != "http" && url.Scheme != "https" {
		return ErrInvalidBaseURLSchema
	}

	// return success
	return nil
}

// BaseURL error constants.
var (
	ErrEmptyBaseURL         = goerr.New("base url is empty")
	ErrInvalidBaseURL       = goerr.New("invalid base url")
	ErrNotAbsoluteBaseURL   = goerr.New("base url must be absolute")
	ErrInvalidBaseURLSchema = goerr.New("base url must be http or https")
)
