package bitflow

import (
	"context"

	"github.com/bitflowex/go-bitflow/bitflow/types"
)

// merchant represents merchant model.
type Merchant struct {
	c *Client
}

func (m *Merchant) GetMe(ctx context.Context) (*types.MerchantResponse, error) {
	// get merchant
	resp, err := get[types.MerchantResponse](ctx, m.c, "merchants/me")
	if err != nil {
		return nil, err
	}

	// return merchant response and success
	return &resp.Data, nil
}
