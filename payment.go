package gobitflow

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type CreatePaymentResponse struct {
	StatusCode int      `json:"statusCode"`
	Status     string   `json:"status"`
	Data       *Payment `json:"data"`
}

type Payment struct {
	ID                  string    `json:"id"`
	MerchantID          string    `json:"merchant_id"`
	Amount              int       `json:"amount"`
	CurrencyCode        string    `json:"currency_code"`
	PaymentAmount       int       `json:"payment_amount"`
	PaymentCurrencyCode string    `json:"payment_currency_code"`
	ReceivedAmount      int       `json:"received_amount"`
	NetworkCode         string    `json:"network_code"`
	Address             string    `json:"address"`
	Link                string    `json:"link"`
	Status              string    `json:"status"`
	PaidAt              time.Time `json:"paid_at"`
	CreatedAt           time.Time `json:"created_at"`
}

func (c *Client) CreatePayment(ctx context.Context, amount float64, cur, net string) (*Payment, error) {
	// create data
	data := map[string]interface{}{
		"amount":        amount,
		"currency_code": cur,
		"network_code":  net,
	}

	// marshal data
	body, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	// create url
	url := fmt.Sprintf("%s/merchant/payment/create", c.baseURL)

	// create request
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(body))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	// set headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.token))

	// send request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %w", err)
	}

	// if response status code is not 200
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to create payment: %s", resp.Status)
	}

	// create response
	var res CreatePaymentResponse
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	// return payment and success
	return res.Data, nil
}

func (c *Client) GetPayment(ctx context.Context, id string) (*Payment, error) {
	// create url
	url := fmt.Sprintf("%s/merchant/payment/%s", c.baseURL, id)

	// create request
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	// set headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.token))

	// send request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %w", err)
	}

	// if response status code is not 200
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to get payment: %s", resp.Status)
	}

	// create response
	var res CreatePaymentResponse
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	// return payment and success
	return res.Data, nil
}
