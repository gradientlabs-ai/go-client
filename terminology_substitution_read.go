package client

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

// ReadTerminologySubstitution retrieves a single terminology substitution by ID.
//
// Note: requires a Management API key.
func (c *Client) ReadTerminologySubstitution(ctx context.Context, id string) (*TerminologySubstitution, error) {
	rsp, err := c.makeRequest(ctx, http.MethodGet, fmt.Sprintf("terminology-substitutions/%s", id), nil)
	if err != nil {
		return nil, err
	}
	defer rsp.Body.Close()

	if err := responseError(rsp); err != nil {
		return nil, err
	}

	var sub TerminologySubstitution
	if err := json.NewDecoder(rsp.Body).Decode(&sub); err != nil {
		return nil, err
	}
	return &sub, nil
}
