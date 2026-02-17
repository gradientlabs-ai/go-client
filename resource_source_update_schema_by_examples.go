package client

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

// SchemaUpdateStrategy represents how the new schema should be applied.
type SchemaUpdateStrategy string

const (
	// SchemaUpdateStrategyMerge merges the inferred schema with the existing schema,
	// preserving existing fields and adding new ones.
	SchemaUpdateStrategyMerge SchemaUpdateStrategy = "merge"

	// SchemaUpdateStrategyReplace completely replaces the existing schema with the newly inferred one.
	SchemaUpdateStrategyReplace SchemaUpdateStrategy = "replace"
)

// UpdateResourceSourceSchemaByExamplesParams represents the request to update a resource source schema by examples.
type UpdateResourceSourceSchemaByExamplesParams struct {
	// Examples is an array of example data payloads that represent the structure
	// of the data your resource source returns.
	Examples []any `json:"examples"`

	// SchemaUpdateStrategy controls how the new schema is applied.
	// Defaults to "merge" if not specified.
	SchemaUpdateStrategy SchemaUpdateStrategy `json:"schema_update_strategy,omitempty"`
}

// UpdateResourceSourceSchemaByExamples updates a resource source schema by providing example data payloads.
// Instead of manually defining the JSON schema structure, you send representative examples of the data
// your resource source returns, and the system automatically infers the schema from these examples.
//
// Note: requires a `Management` API key.
func (c *Client) UpdateResourceSourceSchemaByExamples(ctx context.Context, id string, req *UpdateResourceSourceSchemaByExamplesParams) (*ResourceSource, error) {
	rsp, err := c.makeRequest(ctx, http.MethodPost, fmt.Sprintf("resource-sources/%s/schema-by-examples", id), req)
	if err != nil {
		return nil, err
	}
	defer rsp.Body.Close()

	if err := responseError(rsp); err != nil {
		return nil, err
	}

	var rs ResourceSource
	if err := json.NewDecoder(rsp.Body).Decode(&rs); err != nil {
		return nil, err
	}
	return &rs, nil
}
