package types

// ExecuteTransferRequest represents an execute transfer request.
type ExecuteTransferRequest struct {
	AccountID ID     `json:"account_id"`
	Amount    Amount `json:"amount"`
	Recipient User   `json:"recipient"`
}
