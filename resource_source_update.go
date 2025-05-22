package client

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

// ResourceSourceUpdateParams represents the request to update an existing resource source.
type ResourceSourceUpdateParams struct {
	DisplayName           *string                    `json:"display_name,omitempty"`
	Description           *string                    `json:"description,omitempty"`
	SourceType            *SourceType                `json:"source_type,omitempty"`
	HTTPConfig            *ResourceHTTPDefinition    `json:"http_config,omitempty"`
	WebhookConfig         *ResourceWebhookDefinition `json:"webhook_config,omitempty"`
	AttributeDescriptions map[string]string          `json:"attribute_descriptions,omitempty"`
}

// UpdateResourceSource updates an existing resource source.
//
// Note: requires a `Management` API key.
func (c *Client) UpdateResourceSource(ctx context.Context, id string, req *ResourceSourceUpdateParams) (*ResourceSource, error) {
	rsp, err := c.makeRequest(ctx, http.MethodPut, fmt.Sprintf("resource-sources/%s", id), req)
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
