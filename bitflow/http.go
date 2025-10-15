package bitflow

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"path"
	"strings"

	"github.com/bitflowex/go-bitflow/bitflow/types"
	"github.com/tyrenix/goerr"
)

// get makes a GET request.
func get[T any](ctx context.Context, c *Client, method string) (*types.Response[T], error) {
	_ = ctx
	// create http request
	path.Join(c.baseURL.String(), method)
	req, err := http.NewRequest(http.MethodGet, joinURL(c.baseURL.String(), method), nil)
	if err != nil {
		return nil, goerr.New("failed create http request", err)
	}

	// set authorization header
	req.Header.Set("Authorization", "Bearer "+c.token.String())

	// make http request
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, goerr.New("failed make http request", err)
	}

	// read data
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, goerr.New("failed read data", err)
	}

	// parse json
	res := new(types.Response[T])
	if err := json.Unmarshal(data, resp); err != nil {
		return nil, goerr.New("failed parse json", err)
	}

	// return response and success
	return res, nil
}

// joinURL joins the base URL and method.
func joinURL(baseURL string, method string) string {
	return baseURL + "/" + strings.TrimPrefix(method, "/")
}
