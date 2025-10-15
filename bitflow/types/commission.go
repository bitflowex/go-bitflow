package types

import (
	"github.com/shopspring/decimal"
)

// CommissionPercent represents a commission percent.
type CommissionPercent decimal.Decimal

// String returns the string representation of the commission percent.
func (c CommissionPercent) String() string {
	return c.Decimal().String()
}

// Decimal returns the decimal representation of the commission percent.
func (c CommissionPercent) Decimal() decimal.Decimal {
	return decimal.Decimal(c)
}

// MarshalJSON marshals the commission percent.
func (c CommissionPercent) MarshalJSON() ([]byte, error) {
	return c.Decimal().MarshalJSON()
}

// UnmarshalJSON unmarshals the commission percent.
func (c *CommissionPercent) UnmarshalJSON(data []byte) error {
	// init decimal
	dc := decimal.Decimal{}

	// unmarshal
	if err := dc.UnmarshalJSON(data); err != nil {
		return err
	}

	// set commission percent
	*c = CommissionPercent(dc)

	// return success
	return nil
}
