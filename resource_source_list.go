package client

import (
	"context"
	"encoding/json"
	"net/http"
)

// ResourceSourceListResponse represents the response from listing resource sources.
type ResourceSourceListResponse struct {
	ResourceSources []*ResourceSource `json:"resource_sources"`
}

// ListResourceSources lists all resource sources accessible to the caller's company.
//
// Note: requires a `Management` API key.
func (c *Client) ListResourceSources(ctx context.Context) (*ResourceSourceListResponse, error) {
	rsp, err := c.makeRequest(ctx, http.MethodGet, "resource-sources", nil)
	if err != nil {
		return nil, err
	}
	defer rsp.Body.Close()

	if err := responseError(rsp); err != nil {
		return nil, err
	}

	var result ResourceSourceListResponse
	if err := json.NewDecoder(rsp.Body).Decode(&result); err != nil {
		return nil, err
	}
	return &result, nil
}
