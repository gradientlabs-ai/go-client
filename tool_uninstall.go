package client

import (
	"context"
	"fmt"
	"net/http"
)

// UninstallTool deletes a tool by uninstalling it. Note: this does not
// (yet) check whether those tools are used in procedures. Use with caution!
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
