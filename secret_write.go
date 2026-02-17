package client

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// WriteSecretParams contains the parameters for writing (creating or updating) a secret.
type WriteSecretParams struct {
	// Name is the unique identifier for the secret.
	Name string `json:"-"`

	// Value is the secret value to store.
	Value string `json:"value"`

	// Expiry is the optional expiration time for the secret.
	Expiry *time.Time `json:"expiry,omitempty"`

	// RefreshMechanismHTTP is the optional configuration for automatically refreshing
	// the secret value using an HTTP request (e.g., OAuth token refresh).
	RefreshMechanismHTTP *RefreshMechanismHTTP `json:"refresh_mechanism_http,omitempty"`
}

// WriteSecret creates or updates a secret. If a secret with the given name already exists,
// it will be updated with the new value and configuration.
//
// Note: requires a `Management` API key.
func (c *Client) WriteSecret(ctx context.Context, p *WriteSecretParams) (*Secret, error) {
	path := fmt.Sprintf("secrets/%s", p.Name)

	// Prepare the request body without the Name field (it's in the path)
	body := struct {
		Value                string                `json:"value"`
		Expiry               *time.Time            `json:"expiry,omitempty"`
		RefreshMechanismHTTP *RefreshMechanismHTTP `json:"refresh_mechanism_http,omitempty"`
	}{
		Value:                p.Value,
		Expiry:               p.Expiry,
		RefreshMechanismHTTP: p.RefreshMechanismHTTP,
	}

	rsp, err := c.makeRequest(ctx, http.MethodPut, path, body)
	if err != nil {
		return nil, err
	}
	defer rsp.Body.Close()

	if err := responseError(rsp); err != nil {
		return nil, err
	}

	var secret Secret
	if err := json.NewDecoder(rsp.Body).Decode(&secret); err != nil {
		return nil, err
	}
	return &secret, nil
}
