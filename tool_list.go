package client

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

// ListTools retrieves all tools.
//
// Note: requires a `Management` API key.
func (c *Client) ListTools(ctx context.Context) ([]*Tool, error) {
	rsp, err := c.makeRequest(ctx, http.MethodGet, fmt.Sprintf("tools"), nil)
	if err != nil {
		return nil, err
	}
	defer rsp.Body.Close()

	if err := responseError(rsp); err != nil {
		return nil, err
	}

	var toolList ToolList
	if err := json.NewDecoder(rsp.Body).Decode(&toolList); err != nil {
		return nil, err
	}

	return toolList.Tools, nil
}
