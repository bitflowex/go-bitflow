package types

import (
	"github.com/shopspring/decimal"
	"github.com/tyrenix/goerr"
)

// Amount represents a amount.
type Amount decimal.Decimal

// NewFromInt returns a new amount.
func NewFromInt(i int64) Amount {
	return Amount(decimal.NewFromInt(i))
}

// String returns the string representation of the amount.
func (c Amount) String() string {
	return c.Decimal().String()
}

// Validate validates the amount.
func (a Amount) Validate() error {
	// if amount less then zero
	if a.Decimal().IsNegative() {
		return ErrNegativeAmount
	}

	// return success
	return nil
}

// ValidatePositive validates the amount.
func (a Amount) ValidatePositive() error {
	// validate amount
	if err := a.Validate(); err != nil {
		return err
	}

	// if a not positive
	if !a.Decimal().IsPositive() {
		return ErrNotPositiveAmount
	}

	// return success
	return nil
}

// Decimal returns the decimal representation of the amount.
func (c Amount) Decimal() decimal.Decimal {
	return decimal.Decimal(c)
}

// MarshalJSON marshals the amount.
func (c Amount) MarshalJSON() ([]byte, error) {
	return c.Decimal().MarshalJSON()
}

// UnmarshalJSON unmarshals the amount.
func (c *Amount) UnmarshalJSON(data []byte) error {
	// init decimal
	dc := decimal.Decimal{}

	// unmarshal
	if err := dc.UnmarshalJSON(data); err != nil {
		return err
	}

	// set amount
	*c = Amount(dc)

	// return success
	return nil
}

// Amount error constants.
var (
	ErrNegativeAmount    = goerr.New("amount is negative")
	ErrNotPositiveAmount = goerr.New("amount is not positive")
)
