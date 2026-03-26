package client

import (
	"context"
	"fmt"
	"net/http"
)

type SetProcedureGatedVersionParams struct {
	// MaxDailyConversations limits how many conversations per day can use the gated version.
	// MaxDailyConversations allows gradual rollout of a new procedure version.
	MaxDailyConversations int `json:"max_daily_conversations"`

	// Replace controls whether an existing gated version (if any) will be replaced with a new one.
	// When Replace is false and another gated version already exists, an error will be returned.
	Replace bool `json:"replace"`
}

// SetProcedureGatedVersion sets the gated version of the procedure.
//
// Note: SetProcedureGatedVersion requires a `Management` API key.
func (c *Client) SetProcedureGatedVersion(ctx context.Context, procedureID string, version int, p *SetProcedureGatedVersionParams) error {
	rsp, err := c.makeRequest(ctx, http.MethodPost, fmt.Sprintf("procedures/%s/versions/%d/set-gated", procedureID, version), p)
	if err != nil {
		return err
	}
	defer rsp.Body.Close()

	if err := responseError(rsp); err != nil {
		return err
	}
	return nil
}
