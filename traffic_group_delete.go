package client

import (
	"context"
	"fmt"
	"net/http"
)

// DeleteTrafficGroup deletes a traffic group and all associated targets.
//
// Note: requires a `Management` API key.
func (c *Client) DeleteTrafficGroup(ctx context.Context, trafficGroupID string) error {
	rsp, err := c.makeRequest(ctx, http.MethodDelete, fmt.Sprintf("traffic-groups/%s", trafficGroupID), nil)
	if err != nil {
		return err
	}
	defer rsp.Body.Close()

	if err := responseError(rsp); err != nil {
		return err
	}
	return nil
}
