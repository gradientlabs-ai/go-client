package client

import (
	"context"
	"fmt"
	"net/http"
)

type SetProcedureExperimentVersionParams struct {
	// MaxDailyConversations limits how many conversations per day can use this version.
	// Allows gradual rollout of a new procedure version.
	MaxDailyConversations int `json:"max_daily_conversations"`

	// If Replace is true, an existing experiment (if any) will be replaced with a new one.
	// Otherwise, if another experiment already exists, an error will be returned.
	Replace bool `json:"replace"`
}

// SetProcedureExperimentVersion sets the experiment version of procedure.
//
// Note: requires a `Management` API key.
func (c *Client) SetProcedureExperimentVersion(ctx context.Context, procedureID string, version int, p *SetProcedureExperimentVersionParams) error {
	rsp, err := c.makeRequest(ctx, http.MethodPost, fmt.Sprintf("procedures/%s/versions/%d/set-experiment", procedureID, version), p)
	if err != nil {
		return err
	}
	defer rsp.Body.Close()

	if err := responseError(rsp); err != nil {
		return err
	}
	return nil
}
