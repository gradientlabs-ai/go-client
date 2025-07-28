package client

import (
	"context"
	"fmt"
	"net/http"
)

// SetProcedureLiveVersion sets the live version of procedure.
//
// Note: requires a `Management` API key.
func (c *Client) SetProcedureLiveVersion(ctx context.Context, procedureID string, version int) error {
	rsp, err := c.makeRequest(ctx, http.MethodPost, fmt.Sprintf("procedures/%s/versions/%d/set-live", procedureID, version), nil)
	if err != nil {
		return err
	}
	defer rsp.Body.Close()

	if err := responseError(rsp); err != nil {
		return err
	}
	return nil
}
