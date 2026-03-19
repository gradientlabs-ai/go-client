package client

import (
	"context"
	"encoding/json"
	"net/http"
)

// CreateTrafficGroupParams are the parameters to Client.CreateTrafficGroup.
type CreateTrafficGroupParams struct {
	// Name is the display name for the traffic group.
	Name string `json:"name"`
}

// CreateTrafficGroup creates a new traffic group.
//
// Note: requires a `Management` API key.
func (c *Client) CreateTrafficGroup(ctx context.Context, p *CreateTrafficGroupParams) (*TrafficGroup, error) {
	rsp, err := c.makeRequest(ctx, http.MethodPost, "traffic-groups", p)
	if err != nil {
		return nil, err
	}
	defer rsp.Body.Close()

	if err := responseError(rsp); err != nil {
		return nil, err
	}

	var group TrafficGroup
	if err := json.NewDecoder(rsp.Body).Decode(&group); err != nil {
		return nil, err
	}
	return &group, nil
}
