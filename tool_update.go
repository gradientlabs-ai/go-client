package client

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type UpdateToolParams struct {
	ID          string                    `json:"id"`
	Description string                    `json:"description"`
	Parameters  []ToolParameter           `json:"parameters"`
	Webhook     *ToolWebhookConfiguration `json:"webhook,omitempty"`
	HTTP        *HTTPDefinition           `json:"http,omitempty"`
	Mock        bool                      `json:"mock,omitempty"`
}

// UpdateTool updates an existing tool. It allows callers to convert mock tools
// into real tools, but not the other way around.
//
// Note: requires a `Management` API key.
func (c *Client) UpdateTool(ctx context.Context, p *UpdateToolParams) (*Tool, error) {
	rsp, err := c.makeRequest(ctx, http.MethodPut, fmt.Sprintf("tools/%s", p.ID), p)
	if err != nil {
		return nil, err
	}
	defer rsp.Body.Close()

	if err := responseError(rsp); err != nil {
		return nil, err
	}

	var tool Tool
	if err := json.NewDecoder(rsp.Body).Decode(&tool); err != nil {
		return nil, err
	}
	return &tool, nil
}
