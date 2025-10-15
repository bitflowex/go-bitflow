package types

// AccountName represents an account name.
type AccountName string

// String returns the string representation of the account name.
func (a AccountName) String() string {
	return string(a)
}

// AccountResponse represents an account response.
type AccountResponse struct {
	ID           ID           `json:"id" `
	Name         *AccountName `json:"name"`
	CurrencyCode CurrencyCode `json:"currency_code"`
	Balance      Amount       `json:"balance"`
	HeldBalance  Amount       `json:"held_balance"`
}

// AccountQuery represents an account query.
type AccountQuery struct {
}
