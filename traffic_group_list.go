package client

import (
	"context"
	"encoding/json"
	"net/http"
)

// ListTrafficGroups retrieves all traffic groups.
//
// Note: requires a `Management` API key.
func (c *Client) ListTrafficGroups(ctx context.Context) ([]*TrafficGroup, error) {
	rsp, err := c.makeRequest(ctx, http.MethodGet, "traffic-groups", nil)
	if err != nil {
		return nil, err
	}
	defer rsp.Body.Close()

	if err := responseError(rsp); err != nil {
		return nil, err
	}

	var list trafficGroupList
	if err := json.NewDecoder(rsp.Body).Decode(&list); err != nil {
		return nil, err
	}

	return list.TrafficGroups, nil
}
