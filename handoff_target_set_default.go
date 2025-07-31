package client

import (
	"context"
	"net/http"
)

type SetDefaultHandOffTargetParams struct {
	// ID is the unique identifier for the hand-off target to set as default.
	// This should match an existing hand-off target's ID.
	// Set to empty string to clear the default.
	ID string `json:"id"`

	// Channel is the conversation channel for which to set the default hand-off target.
	Channel Channel `json:"channel"`
}

// SetDefaultHandOffTarget sets the default hand off target that the AI agent will
// use when handing off the conversation, if there is no specific target for that intent
// or procedure.
//
// Note: requires a `Management` API key.
func (c *Client) SetDefaultHandOffTarget(ctx context.Context, p *SetDefaultHandOffTargetParams) error {
	rsp, err := c.makeRequest(ctx, http.MethodPut, "hand-off-targets/default", p)
	if err != nil {
		return err
	}
	defer rsp.Body.Close()

	if err := responseError(rsp); err != nil {
		return err
	}
	return nil
}
