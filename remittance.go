package gobitflow

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

func (c *Client) CreateRemittance(ctx context.Context, amount float64, cur, recipient string) (bool, error) {
	// create data
	data := map[string]any{
		"amount":        amount,
		"recipient":     recipient,
		"currency_code": cur,
	}

	// marshal data
	body, err := json.Marshal(data)
	if err != nil {
		return false, err
	}

	// create url
	url := fmt.Sprintf("%s/merchant/remittance/create", c.baseURL)

	// create request
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(body))
	if err != nil {
		return false, fmt.Errorf("failed to create request: %w", err)
	}

	// set headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.token))

	// send request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return false, fmt.Errorf("failed to send request: %w", err)
	}

	// if response status code is not 200
	if resp.StatusCode != http.StatusOK {
		return false, fmt.Errorf("failed to create payment: %s", resp.Status)
	}

	// return success
	return true, nil
}
