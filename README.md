# Bitflow Go Client

A robust and type-safe Go client library for interacting with the Bitflow API.

## Table of Contents

- [Features](#features)
- [Installation](#installation)
- [Usage](#usage)
- [Configuration](#configuration)
- [Error Handling](#error-handling)
- [Project Structure](#project-structure)
- [License](#license)

## Features

-   **Type-Safe API Interactions:** Utilizes custom Go types for all API requests and responses, ensuring strong type checking at compile time.
-   **Robust Input Validation:** Comprehensive validation for client configuration (Token, BaseURL) and API request parameters.
-   **Service-Oriented Architecture:** API endpoints are organized into distinct services (Merchants, Invoices, Accounts, Transfers) for clear separation of concerns.
-   **Flexible Configuration:** Employs the functional options pattern for easy and extensible client setup.
-   **Consistent Error Handling:** Uses `goerr` for structured and predictable error reporting.
-   **Arbitrary Precision Arithmetic:** Uses `shopspring/decimal` for financial amounts and percentages to ensure accuracy.

## Installation

To install the Bitflow Go Client, use `go get`:

```bash
go get github.com/bitflowex/go-bitflow
```

## Usage

Here's a quick example of how to initialize the client and fetch merchant information:

```go
package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/bitflowex/go-bitflow/bitflow"
	"github.com/bitflowex/go-bitflow/bitflow/types"
)

func main() {
	// Replace with your actual API token and base URL
	apiToken := os.Getenv("BITFLOW_API_TOKEN")
	baseURL := os.Getenv("BITFLOW_BASE_URL") // e.g., "https://bitflow.ws/api/v1"

	if apiToken == "" || baseURL == "" {
		log.Fatal("BITFLOW_API_TOKEN and BITFLOW_BASE_URL environment variables must be set")
	}

	// Initialize the client with functional options
	client, err := bitflow.NewClient(
		bitflow.WithToken(types.Token(apiToken)),
		bitflow.WithBaseURL(types.BaseURL(baseURL)),
	)
	if err != nil {
		log.Fatalf("Failed to create Bitflow client: %v", err)
	}

	ctx := context.Background()

	// Example: Get authenticated merchant information
	merchant, err := client.Merchants.GetMe(ctx)
	
	if err != nil {
		log.Fatalf("Failed to get merchant info: %v", err)
	}

	fmt.Printf("Authenticated Merchant: %s (ID: %s)\n", merchant.Name, merchant.ID)

	// Example: Create an invoice
	createInvoiceReq := types.CreateInvoiceRequest{
		Amount:       types.NewFromInt(10000), // 100.00 units
		CurrencyCode: types.CurrencyRUB,
		ExpiresAt:    time.Now().Add(24 * time.Hour),
		Description:  types.InvoiceDescription("Test invoice from Go client"),
	}
	invoice, err := client.Invoices.Create(ctx, createInvoiceReq)
	if err != nil {
		log.Fatalf("Failed to create invoice: %v", err)
	}
	fmt.Printf("Created Invoice ID: %s, Status: %s\n", invoice.ID, invoice.Status)
}
```

## Configuration

The `bitflow.NewClient` function accepts functional options to configure the client:

-   `bitflow.WithToken(token types.Token)`: Sets the API authentication token.
-   `bitflow.WithBaseURL(baseURL types.BaseURL)`: Sets the base URL for the Bitflow API.

## Error Handling

The client returns standard Go `error` types. Custom errors are defined within the `bitflow/types` package, following the `Err[Reason][Entity]` naming convention. You can check for specific error types using `errors.Is()` or `errors.As()`.

API responses that indicate an error (e.g., `resp.Status != types.ResponseSuccess`) will return an error containing details from the API's error message.

## Project Structure

The library is structured as follows:

```
bitflow-client/
├── go.mod
├── README.md
└── bitflow/
    ├── client.go             # Main client initialization and configuration
    ├── http.go               # Generic HTTP request helpers (get, post)
    ├── account.go            # Account service methods
    ├── invoice.go            # Invoice service methods
    ├── merchant.go           # Merchant service methods
    ├── transfer.go           # Transfer service methods
    └── types/
        ├── account.go
        ├── amount.go
        ├── base_url.go
        ├── category.go
        ├── commission.go
        ├── currency.go
        ├── id.go
        ├── invoice.go
        ├── merchant.go
        ├── response.go
        ├── token.go
        ├── transfer.go
        └── user.go
```

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
