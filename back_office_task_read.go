package client

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

// ReadBackOfficeTask retrieves the current state of a back-office task.
//
// Note: requires a Public API key.
func (c *Client) ReadBackOfficeTask(ctx context.Context, taskID string) (*BackOfficeTask, error) {
	rsp, err := c.makeRequest(ctx, http.MethodGet, fmt.Sprintf("back-office-tasks/%s/read", taskID), nil)
	if err != nil {
		return nil, err
	}
	defer rsp.Body.Close()

	if err := responseError(rsp); err != nil {
		return nil, err
	}

	var task BackOfficeTask
	if err := json.NewDecoder(rsp.Body).Decode(&task); err != nil {
		return nil, err
	}
	return &task, nil
}
