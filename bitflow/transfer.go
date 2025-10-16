package bitflow

import (
	"context"

	"github.com/bitflowex/go-bitflow/bitflow/types"
	"github.com/tyrenix/goerr"
)

// Transfer represents transfer model.
type Transfer struct {
	c *Client
}

// Execute executes a transfer.
func (t *Transfer) Execute(ctx context.Context, req types.ExecuteTransferRequest) (bool, error) {
	// validate account id
	if err := req.AccountID.Validate(); err != nil {
		return false, err
	}
	// validate amount
	if err := req.Amount.ValidatePositive(); err != nil {
		return false, err
	}
	// if exists recipient id, validate it
	if req.Recipient.ID != nil {
		if err := req.Recipient.ID.Validate(); err != nil {
			return false, err
		}
	}
	// if exists recipient username, validate it
	if req.Recipient.Username != nil {
		if err := req.Recipient.Username.Validate(); err != nil {
			return false, err
		}
	}

	// execute transfer
	resp, err := post[types.ExecuteTransferRequest, bool](ctx, t.c, "merchants/me/transfers", req)
	if err != nil {
		return false, err
	}

	// if status is not success
	if resp.Status != types.ResponseSuccess {
		return false, goerr.New("failed execute transfer", resp.Error, resp.Message)
	}

	// return success
	return true, nil
}
