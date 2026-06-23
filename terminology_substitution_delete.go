package client

import (
	"context"
	"fmt"
	"net/http"
)

// DeleteTerminologySubstitution removes a terminology substitution.
//
// Note: requires a Management API key.
func (c *Client) DeleteTerminologySubstitution(ctx context.Context, id string) error {
	rsp, err := c.makeRequest(ctx, http.MethodDelete, fmt.Sprintf("terminology-substitutions/%s", id), nil)
	if err != nil {
		return err
	}
	defer rsp.Body.Close()

	if err := responseError(rsp); err != nil {
		return err
	}
	return nil
}
