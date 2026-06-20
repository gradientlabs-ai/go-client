package client

import (
	"context"
	"encoding/json"
	"net/http"
)

// TerminologySubstitutionCreateParams are the parameters to Client.CreateTerminologySubstitution.
type TerminologySubstitutionCreateParams struct {
	Blocked                   string `json:"blocked"`
	BlockedDescription        string `json:"blocked_description"`
	Replacement               string `json:"replacement"`
	ResourceTypeID            string `json:"resource_type_id,omitempty"`
	ResourceAttributeJSONPath string `json:"resource_attribute_json_path,omitempty"`
	ResourceValueToMatch      string `json:"resource_value_to_match,omitempty"`
}

// CreateTerminologySubstitution creates a new terminology substitution.
//
// Note: requires a Management API key.
func (c *Client) CreateTerminologySubstitution(ctx context.Context, p *TerminologySubstitutionCreateParams) (*TerminologySubstitution, error) {
	rsp, err := c.makeRequest(ctx, http.MethodPost, "terminology-substitutions", p)
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
