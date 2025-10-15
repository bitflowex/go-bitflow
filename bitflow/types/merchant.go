package types

import (
	"time"
)

// MerchantName represents a merchant name.
type MerchantName string

// String returns the string representation of the merchant name.
func (m MerchantName) String() string {
	return string(m)
}

// MerchantProjectURL represents a merchant project URL.
type MerchantProjectURL string

// String returns the string representation of the merchant project URL.
func (m MerchantProjectURL) String() string {
	return string(m)
}

// MerchantResponse represents a merchant response.
type MerchantResponse struct {
	Name              MerchantName       `json:"name"`
	ProjectURL        MerchantProjectURL `json:"project_url"`
	CategoryCode      CategoryCode       `json:"category_code"`
	CommissionPercent CommissionPercent  `json:"commission_percent"`
	CreatedAt         time.Time          `json:"created_at"`
}
