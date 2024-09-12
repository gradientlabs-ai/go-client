package client

import (
	"context"
	"net/http"
)

type UpsertHandOffTargetParams struct {
	// ID is your identifier of choice for this hand-off target. Can be anything consisting
	// of letters, numbers, or any of the following characters: `_` `-` `+` `=`.
	ID string `json:"id"`

	// Name is the hand-off target’s name. This cannot be nil.
	Name string `json:"name"`
}

// UpsertHandOffTarget inserts or updates a hand-off target
func (c *Client) UpsertHandOffTarget(ctx context.Context, p *UpsertHandOffTargetParams) error {
	rsp, err := c.makeRequest(ctx, http.MethodPost, "hand-off-targets", p)
	if err != nil {
		return err
	}
	defer rsp.Body.Close()

	if err := responseError(rsp); err != nil {
		return err
	}
	return nil
}
