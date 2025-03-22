package client

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type Argument struct {
	// Name is the parameter name.
	Name string `json:"name"`

	// Value is the value of the argument.
	// It is a string here, but it will be converted to the
	// appropriate type when the tool is called.
	Value string `json:"value"`
}

type ExecutionParams struct {
	ID string `json:"id"`

	Arguments []Argument `json:"arguments"`
}

type ExecuteResult struct {
	ID string `json:"id"`

	// Result is the JSON-encoded result of the tool execution.
	RawResult json.RawMessage `json:"raw"`
}

// ExecuteTool executes a tool with the provided arguments, so that you can
// test a tool end-to-end.
//
// Note: requires a `Management` API key.
func (c *Client) ExecuteTool(ctx context.Context, p *ExecutionParams) (*ExecuteResult, error) {
	rsp, err := c.makeRequest(ctx, http.MethodPost, fmt.Sprintf("tools/%s/execute", p.ID), p)
	if err != nil {
		return nil, err
	}
	defer rsp.Body.Close()

	if err := responseError(rsp); err != nil {
		return nil, err
	}

	var toolResponse ExecuteResult
	if err := json.NewDecoder(rsp.Body).Decode(&toolResponse); err != nil {
		return nil, err
	}
	return &toolResponse, nil
}
