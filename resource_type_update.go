package client

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

// ResourceTypeUpdateParams represents the request to update an existing resource type.
type ResourceTypeUpdateParams struct {
	DisplayName     *string          `json:"display_name,omitempty"`
	Description     *string          `json:"description,omitempty"`
	Scope           *Scope           `json:"scope,omitempty"`
	RefreshStrategy *RefreshStrategy `json:"refresh_strategy,omitempty"`
	SourceConfig    *SourceConfig    `json:"source_config,omitempty"`
	IsEnabled       *bool            `json:"is_enabled,omitempty"`
}

// UpdateResourceType updates an existing resource type.
//
// Note: requires a `Management` API key.
func (c *Client) UpdateResourceType(ctx context.Context, id string, req *ResourceTypeUpdateParams) (*ResourceType, error) {
	rsp, err := c.makeRequest(ctx, http.MethodPut, fmt.Sprintf("resource-types/%s", id), req)
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
