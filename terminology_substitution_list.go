package client

import (
	"context"
	"encoding/json"
	"net/http"
)

// TerminologySubstitutionListResponse is the response from Client.ListTerminologySubstitutions.
type TerminologySubstitutionListResponse struct {
	Substitutions []*TerminologySubstitution `json:"substitutions"`
}

// ListTerminologySubstitutions returns all terminology substitutions configured
// for your company.
//
// Note: requires a Management API key.
func (c *Client) ListTerminologySubstitutions(ctx context.Context) (*TerminologySubstitutionListResponse, error) {
	rsp, err := c.makeRequest(ctx, http.MethodGet, "terminology-substitutions", nil)
	if err != nil {
		return nil, err
	}
	defer rsp.Body.Close()

	if err := responseError(rsp); err != nil {
		return nil, err
	}

	var result TerminologySubstitutionListResponse
	if err := json.NewDecoder(rsp.Body).Decode(&result); err != nil {
		return nil, err
	}
	return &result, nil
}
