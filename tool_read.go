package client

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

// ReadTool retrieves a new tool.
//
// Note: requires a `Management` API key.
func (c *Client) ReadTool(ctx context.Context, toolID string) (*Tool, error) {
	rsp, err := c.makeRequest(ctx, http.MethodGet, fmt.Sprintf("tools/%s", toolID), nil)
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
