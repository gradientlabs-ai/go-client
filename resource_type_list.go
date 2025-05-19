package client

import (
	"context"
	"encoding/json"
	"net/http"
)

// ResourceTypeListResponse represents the response from listing resource types.
type ResourceTypeListResponse struct {
	ResourceTypes []*ResourceType `json:"resource_types"`
}

// ListResourceTypes lists all resource types accessible to the caller's company.
//
// Note: requires a `Management` API key.
func (c *Client) ListResourceTypes(ctx context.Context) (*ResourceTypeListResponse, error) {
	rsp, err := c.makeRequest(ctx, http.MethodGet, "resource-types", nil)
	if err != nil {
		return nil, err
	}
	defer rsp.Body.Close()

	if err := responseError(rsp); err != nil {
		return nil, err
	}

	var result ResourceTypeListResponse
	if err := json.NewDecoder(rsp.Body).Decode(&result); err != nil {
		return nil, err
	}
	return &result, nil
}
