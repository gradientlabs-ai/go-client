package client

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

// ReadResourceType retrieves a specific resource type by its ID.
//
// Note: requires a `Management` API key.
func (c *Client) ReadResourceType(ctx context.Context, id string) (*ResourceType, error) {
	rsp, err := c.makeRequest(ctx, http.MethodGet, fmt.Sprintf("resource-types/%s", id), nil)
	if err != nil {
		return nil, err
	}
	defer rsp.Body.Close()

	if err := responseError(rsp); err != nil {
		return nil, err
	}

	var rt ResourceType
	if err := json.NewDecoder(rsp.Body).Decode(&rt); err != nil {
		return nil, err
	}
	return &rt, nil
}
