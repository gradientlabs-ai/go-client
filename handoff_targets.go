package client

import (
	"context"
	"encoding/json"
	"net/http"
)

type HandOffTargets struct {
	Targets []*HandOffTarget `json:"targets"`
}

// ListHandOffTargets returns all of your hand off targets.
//
// Note: requires a `Management` API key.
func (c *Client) ListHandOffTargets(ctx context.Context) (*HandOffTargets, error) {
	rsp, err := c.makeRequest(ctx, http.MethodGet, "hand-off-targets", nil)
	if err != nil {
		return nil, err
	}
	defer rsp.Body.Close()

	if err := responseError(rsp); err != nil {
		return nil, err
	}

	var targets HandOffTargets
	if err := json.NewDecoder(rsp.Body).Decode(&targets); err != nil {
		return nil, err
	}
	return &targets, nil
}
