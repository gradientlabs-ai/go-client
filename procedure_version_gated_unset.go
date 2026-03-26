package client

import (
	"context"
	"fmt"
	"net/http"
)

// UnsetProcedureGatedVersion unsets the gated version of the procedure.
//
// Note: UnsetProcedureGatedVersion requires a `Management` API key.
func (c *Client) UnsetProcedureGatedVersion(ctx context.Context, procedureID string, version int) error {
	rsp, err := c.makeRequest(ctx, http.MethodPost, fmt.Sprintf("procedures/%s/versions/%d/unset-gated", procedureID, version), nil)
	if err != nil {
		return err
	}
	defer rsp.Body.Close()

	if err := responseError(rsp); err != nil {
		return err
	}
	return nil
}
