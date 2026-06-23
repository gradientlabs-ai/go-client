package client

import (
	"context"
	"fmt"
	"net/http"
)

// DeleteTrafficGroupExclusion removes an exclusion from a traffic group.
//
// Note: requires a Management API key.
func (c *Client) DeleteTrafficGroupExclusion(ctx context.Context, trafficGroupID string, targetID string) error {
	rsp, err := c.makeRequest(ctx, http.MethodDelete, fmt.Sprintf("traffic-groups/%s/exclusions/%s", trafficGroupID, targetID), nil)
	if err != nil {
		return err
	}
	defer rsp.Body.Close()

	if err := responseError(rsp); err != nil {
		return err
	}
	return nil
}
