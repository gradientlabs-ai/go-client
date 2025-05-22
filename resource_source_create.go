package client

import (
	"context"
	"encoding/json"
	"net/http"
)

// ResourceSourceCreateParams represents the request to create a new resource source.
type ResourceSourceCreateParams struct {
	DisplayName           string                     `json:"display_name"`
	Description           string                     `json:"description,omitempty"`
	SourceType            SourceType                 `json:"source_type"`
	HTTPConfig            *ResourceHTTPDefinition    `json:"http_config,omitempty"`
	WebhookConfig         *ResourceWebhookDefinition `json:"webhook_config,omitempty"`
	AttributeDescriptions map[string]string          `json:"attribute_descriptions,omitempty"`
}

// CreateResourceSource creates a new resource source.
//
// Note: requires a `Management` API key.
func (c *Client) CreateResourceSource(ctx context.Context, req *ResourceSourceCreateParams) (*ResourceSource, error) {
	rsp, err := c.makeRequest(ctx, http.MethodPost, "resource-sources", req)
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
