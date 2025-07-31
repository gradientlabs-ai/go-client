package client

import (
	"context"
	"net/http"
)

type HandOffTargetDeleteParams struct {
	// ID is the unique identifier for the hand-off target to delete
	ID string `json:"id"`
}

// DeleteHandOffTarget deletes a hand-off target. This will fail if the hand off target
// is in use - either in a procedure, or in an intent.
//
// Note: requires a `Management` API key.
func (c *Client) DeleteHandOffTarget(ctx context.Context, p *HandOffTargetDeleteParams) error {
	rsp, err := c.makeRequest(ctx, http.MethodDelete, "hand-off-targets", p)
	if err != nil {
		return err
	}
	defer rsp.Body.Close()

	if err := responseError(rsp); err != nil {
		return err
	}
	return nil
}
