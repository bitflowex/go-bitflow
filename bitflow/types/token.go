package types

import (
	"github.com/btcsuite/btcd/btcutil/base58"
	"github.com/tyrenix/goerr"
)

// Token represents a token.
type Token string

// String returns the string representation of the token.
func (t Token) String() string {
	return string(t)
}

// Validate validates the token.
func (t Token) Validate() error {
	// if token is empty
	if t == "" {
		return ErrEmptyToken
	}

	// check prefix
	if t[0:3] != "sk-" {
		return ErrRequiredTokenPrefix
	}

	// decode base58
	if len(base58.Decode(t.String()[3:])) != 32 {
		return ErrInvalidTokenLength
	}

	// return success
	return nil
}

// Token error constants.
var (
	ErrEmptyToken          = goerr.New("token is empty")
	ErrRequiredTokenPrefix = goerr.New("token must start with 'sk-'")
	ErrInvalidTokenLength  = goerr.New("token must be 32 characters long")
)
