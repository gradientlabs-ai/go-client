package client

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

// ReadResourceSource retrieves a specific resource source by its ID.
//
// Note: requires a `Management` API key.
func (c *Client) ReadResourceSource(ctx context.Context, id string) (*ResourceSource, error) {
	rsp, err := c.makeRequest(ctx, http.MethodGet, fmt.Sprintf("resource-sources/%s", id), nil)
	if err != nil {
		return nil, err
	}
	defer rsp.Body.Close()

	if err := responseError(rsp); err != nil {
		return nil, err
	}

	var rs ResourceSource
	if err := json.NewDecoder(rsp.Body).Decode(&rs); err != nil {
		return nil, err
	}
	return &rs, nil
}
