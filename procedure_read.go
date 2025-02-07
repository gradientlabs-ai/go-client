package client

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

// ReadProcedure returns a procedure.
//
// Note: requires a `Management` API key.
func (c *Client) ReadProcedure(ctx context.Context, procedureID string) (*Procedure, error) {
	rsp, err := c.makeRequest(ctx, http.MethodGet, fmt.Sprintf("procedures/%s", procedureID), nil)
	if err != nil {
		return nil, err
	}
	defer rsp.Body.Close()

	if err := responseError(rsp); err != nil {
		return nil, err
	}

	var proc Procedure
	if err := json.NewDecoder(rsp.Body).Decode(&proc); err != nil {
		return nil, err
	}
	return &proc, nil
}
