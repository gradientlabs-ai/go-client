package client

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type GetDefaultHandOffTargetParams struct {
	// Channel is the conversation channel for which to get the default hand-off target.
	Channel Channel `json:"channel"`
}

type GetDefaultHandOffTargetResponse struct {
	// ID is the unique identifier for the default hand-off target.
	// Empty string if no default is set.
	ID string `json:"id"`
}

// GetDefaultHandOffTarget gets the current default hand-off target for the company.
//
// Note: requires a `Management` API key.
func (c *Client) GetDefaultHandOffTarget(ctx context.Context, p *GetDefaultHandOffTargetParams) (*GetDefaultHandOffTargetResponse, error) {
	path := fmt.Sprintf("hand-off-targets/default?channel=%s", p.Channel)
	rsp, err := c.makeRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}
	defer rsp.Body.Close()

	if err := responseError(rsp); err != nil {
		return nil, err
	}

	var response GetDefaultHandOffTargetResponse
	if err := json.NewDecoder(rsp.Body).Decode(&response); err != nil {
		return nil, err
	}
	return &response, nil
}
