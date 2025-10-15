package types

import (
	"time"

	"github.com/tyrenix/goerr"
)

// InvoiceStatus represents an invoice status.
type InvoiceStatus string

// Invoice statuses.
const (
	InvoicePending InvoiceStatus = "pending"
	InvoicePaid    InvoiceStatus = "paid"
	InvoiceExpired InvoiceStatus = "expired"
)

// InvoiceDescription represents an invoice description.
type InvoiceDescription string

// String returns the string representation of the invoice description.
func (d InvoiceDescription) String() string {
	return string(d)
}

// Validate validates the invoice description.
func (d InvoiceDescription) Validate() error {
	// get runes
	r := []rune(d)

	// if runes less than 3
	if len(r) < 3 {
		return ErrTooShortInvoiceDescription
	}
	// if runes more than 120
	if len(r) > 120 {
		return ErrTooLongInvoiceDescription
	}

	// return success
	return nil
}

// InvoiceExternalID represents an invoice external identifier.
type InvoiceExternalID string

// String returns the string representation of the invoice external identifier.
func (e InvoiceExternalID) String() string {
	return string(e)
}

// Validate validates the invoice external identifier.
func (e InvoiceExternalID) Validate() error {
	// get runes
	r := []rune(e)

	// if runes less than 8
	if len(r) < 8 {
		return ErrTooShortInvoiceExternalID
	}
	// if runes more than 64
	if len(r) > 64 {
		return ErrTooLongInvoiceExternalID
	}

	// return success
	return nil
}

// InvoiceResponse represents an invoice response.
type InvoiceResponse struct {
	ID               ID                  `json:"id"`
	Status           InvoiceStatus       `json:"status"`
	BaseAmount       Amount              `json:"base_amount"`
	BaseCurrencyCode CurrencyCode        `json:"base_currency_code"`
	Description      *InvoiceDescription `json:"description,omitempty"`
	ExternalID       *InvoiceExternalID  `json:"external_id,omitempty"`

	PaymentAmount       *Amount       `json:"payment_amount,omitempty" `
	PaymentCurrencyCode *CurrencyCode `json:"payment_currency_code,omitempty" `
	PaymentCommission   *Amount       `json:"payment_commission,omitempty" `
	PaymentNetworkCode  *NetworkCode  `json:"payment_network_code,omitempty"`
	PaymentAddress      *Address      `json:"payment_address,omitempty"`
	PaymentURL          *string       `json:"payment_url,omitempty" `

	PaidAmount *Amount    `json:"paid_amount,omitempty"`
	PaidAt     *time.Time `json:"paid_at,omitempty"`

	ExpiresAt time.Time `json:"expires_at"`
	CreatedAt time.Time `json:"created_at"`
}

// CreateInvoiceRequest represents a create invoice request.
type CreateInvoiceRequest struct {
	Amount       Amount              `json:"base_amount"`
	CurrencyCode CurrencyCode        `json:"base_currency_code"`
	Description  *InvoiceDescription `json:"description"`
	ExternalID   *InvoiceExternalID  `json:"external_id"`
	ExpiresAt    time.Time           `json:"expires_at"`
}

// Invoice description error constants.
var (
	ErrTooShortInvoiceDescription = goerr.New("invoice description too short")
	ErrTooLongInvoiceDescription  = goerr.New("invoice description too long")
)

// Invoice external ID error constants.
var (
	ErrTooShortInvoiceExternalID = goerr.New("invoice external id too short")
	ErrTooLongInvoiceExternalID  = goerr.New("invoice external id too long")
)
