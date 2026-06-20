package client

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

// TerminologySubstitutionUpdateParams are the parameters to Client.UpdateTerminologySubstitution.
type TerminologySubstitutionUpdateParams struct {
	Blocked                   string `json:"blocked"`
	BlockedDescription        string `json:"blocked_description"`
	Replacement               string `json:"replacement"`
	ResourceTypeID            string `json:"resource_type_id,omitempty"`
	ResourceAttributeJSONPath string `json:"resource_attribute_json_path,omitempty"`
	ResourceValueToMatch      string `json:"resource_value_to_match,omitempty"`
}

// UpdateTerminologySubstitution replaces an existing terminology substitution.
//
// Note: requires a Management API key.
func (c *Client) UpdateTerminologySubstitution(ctx context.Context, id string, p *TerminologySubstitutionUpdateParams) (*TerminologySubstitution, error) {
	rsp, err := c.makeRequest(ctx, http.MethodPut, fmt.Sprintf("terminology-substitutions/%s", id), p)
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
