package client

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

// UpdateTrafficGroupParams are the parameters to Client.UpdateTrafficGroup.
type UpdateTrafficGroupParams struct {
	// ID is the unique identifier of the traffic group to update.
	ID string `json:"-"`

	// Name is the new display name for the traffic group.
	Name string `json:"name"`
}

// UpdateTrafficGroup updates an existing traffic group.
//
// Note: requires a `Management` API key.
func (c *Client) UpdateTrafficGroup(ctx context.Context, p *UpdateTrafficGroupParams) (*TrafficGroup, error) {
	rsp, err := c.makeRequest(ctx, http.MethodPut, fmt.Sprintf("traffic-groups/%s", p.ID), p)
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
