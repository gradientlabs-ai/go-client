package client

import (
	"context"
	"encoding/json"
	"net/http"
)

// SecretsListResponse contains the list of all secrets.
type SecretsListResponse struct {
	Secrets []*Secret `json:"secrets"`
}

// ListSecrets returns all of your secrets.
//
// Note: requires a `Management` API key.
func (c *Client) ListSecrets(ctx context.Context) (*SecretsListResponse, error) {
	rsp, err := c.makeRequest(ctx, http.MethodGet, "secrets", nil)
	if err != nil {
		return nil, err
	}
	defer rsp.Body.Close()

	if err := responseError(rsp); err != nil {
		return nil, err
	}

	var secrets SecretsListResponse
	if err := json.NewDecoder(rsp.Body).Decode(&secrets); err != nil {
		return nil, err
	}
	return &secrets, nil
}
