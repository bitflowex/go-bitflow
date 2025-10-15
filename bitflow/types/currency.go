package types

import (
	"unicode"

	"github.com/tyrenix/goerr"
)

// CurrencyCode represents a currency code.
type CurrencyCode string

// Currency codes.
const (
	CurrencyRUB  CurrencyCode = "RUB"
	CurrencyUSD  CurrencyCode = "USD"
	CurrencyUSDT CurrencyCode = "USDT"
	CurrencyTON  CurrencyCode = "TON"
	CurrencyTRX  CurrencyCode = "TRX"
)

// String returns the string representation of the currency code.
func (c CurrencyCode) String() string {
	return string(c)
}

// Validate validates the currency code.
func (c CurrencyCode) Validate() error {
	// if currency code is empty
	if c == "" {
		return ErrEmptyCurrencyCode
	}

	// check all symbols
	for _, r := range c {
		// check if symbol is a letter
		if !unicode.IsLetter(r) || !unicode.IsUpper(r) {
			return ErrInvalidCurrencyCodeFormat
		}
	}

	// return success
	return nil
}

// CurrencyCode error constants.
var (
	ErrEmptyCurrencyCode         = goerr.New("currency code is empty")
	ErrInvalidCurrencyCodeFormat = goerr.New("currency code must be a string of uppercase letters")
)
