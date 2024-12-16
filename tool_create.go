package client

import (
	"context"
	"encoding/json"
	"net/http"
)

// CreateTool creates a new tool.
//
// Note: requires a `Management` API key.
func (c *Client) CreateTool(ctx context.Context, p *Tool) (*Tool, error) {
	rsp, err := c.makeRequest(ctx, http.MethodPost, "tools", p)
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
