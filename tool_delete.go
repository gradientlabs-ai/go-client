package client

import (
	"context"
	"fmt"
	"net/http"
)

// DeleteTool deletes a tool. Note: You can't delete a tool that is eing used in a live procedure.
//
// Note: requires a `Management` API key.
func (c *Client) UninstallTool(ctx context.Context, toolID string) error {
	rsp, err := c.makeRequest(ctx, http.MethodDelete, fmt.Sprintf("tools/%s", toolID), nil)
	if err != nil {
		return err
	}
	defer rsp.Body.Close()

	if err := responseError(rsp); err != nil {
		return err
	}
	return nil
}
