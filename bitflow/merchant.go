package bitflow

import (
	"context"

	"github.com/bitflowex/go-bitflow/bitflow/types"
	"github.com/tyrenix/goerr"
)

// merchant represents merchant model.
type Merchant struct {
	c *Client
}

// GetMe gets the merchant.
func (m Merchant) GetMe(ctx context.Context) (*types.MerchantResponse, error) {
	// get merchant
	resp, err := get[types.MerchantResponse](ctx, m.c, "merchants/me")
	if err != nil {
		return nil, err
	}

	// if status is not success
	if resp.Status != types.ResponseSuccess {
		return nil, goerr.New("failed get merchant", resp.Error, resp.Message)
	}

	// return merchant response and success
	return &resp.Data, nil
}
