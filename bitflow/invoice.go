package bitflow

import (
	"context"

	"github.com/bitflowex/go-bitflow/bitflow/types"
	"github.com/tyrenix/goerr"
)

// Invoice represents invoice model.
type Invoice struct {
	c *Client
}

// Create creates an invoice.
func (i *Invoice) Create(ctx context.Context, req types.CreateInvoiceRequest) (*types.InvoiceResponse, error) {
	// validate amount
	if err := req.Amount.ValidatePositive(); err != nil {
		return nil, err
	}
	// validate currency code
	if err := req.CurrencyCode.Validate(); err != nil {
		return nil, err
	}
	// if exists description validate it
	if req.Description != nil {
		if err := req.Description.Validate(); err != nil {
			return nil, err
		}
	}
	// if exists external id validate it
	if req.ExternalID != nil {
		if err := req.ExternalID.Validate(); err != nil {
			return nil, err
		}
	}

	// create invoice
	resp, err := post[types.CreateInvoiceRequest, *types.InvoiceResponse](ctx, i.c, "invoices", req)
	if err != nil {
		return nil, err
	}

	// if status is not success
	if resp.Status != types.ResponseSuccess {
		return nil, goerr.New("failed create invoice", resp.Error, resp.Message)
	}

	// return invoice response and success
	return resp.Data, nil
}

// Get gets invoice by id.
func (i *Invoice) Get(ctx context.Context, id types.ID) (*types.InvoiceResponse, error) {
	// validate id
	if err := id.Validate(); err != nil {
		return nil, err
	}

	// get invoice
	resp, err := get[types.InvoiceResponse](ctx, i.c, "invoices/:"+id.String())
	if err != nil {
		return nil, err
	}

	// if status is not success
	if resp.Status != types.ResponseSuccess {
		return nil, goerr.New("failed get invoice", resp.Error, resp.Message)
	}

	// return invoice response and success
	return &resp.Data, nil
}
