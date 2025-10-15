package bitflow

import (
	"context"

	"github.com/bitflowex/go-bitflow/bitflow/types"
	"github.com/tyrenix/goerr"
)

// Account represents account model.
type Account struct {
	c *Client
}

// Find finds merchant accounts
func (a Account) Find(ctx context.Context, query types.AccountQuery) ([]*types.AccountResponse, error) {
	// get merchant
	resp, err := get[[]*types.AccountResponse](ctx, a.c, "merchants/me/accounts")
	if err != nil {
		return nil, err
	}

	// if status is not success
	if resp.Status != types.ResponseSuccess {
		return nil, goerr.New("failed find accounts", resp.Error, resp.Message)
	}

	// return merchant response and success
	return resp.Data, nil
}
