package client

import (
	"context"
	"fmt"
	"net/http"
)

// RevokeSecretParams contains the parameters for revoking (deleting) a secret.
type RevokeSecretParams struct {
	// Name is the unique identifier for the secret to revoke.
	Name string `json:"-"`
}

// RevokeSecret permanently deletes a secret. This action cannot be undone.
//
// Note: requires a `Management` API key.
func (c *Client) RevokeSecret(ctx context.Context, p *RevokeSecretParams) error {
	path := fmt.Sprintf("secrets/%s", p.Name)

	rsp, err := c.makeRequest(ctx, http.MethodDelete, path, nil)
	if err != nil {
		return err
	}
	defer rsp.Body.Close()

	if err := responseError(rsp); err != nil {
		return err
	}
	return nil
}
