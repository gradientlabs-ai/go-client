package client

import (
	"context"
	"fmt"
	"net/http"
)

// DeleteResourceType deletes a resource type by its ID.
//
// Note: requires a `Management` API key.
func (c *Client) DeleteResourceType(ctx context.Context, id string) error {
	rsp, err := c.makeRequest(ctx, http.MethodDelete, fmt.Sprintf("resource-types/%s", id), nil)
	if err != nil {
		return err
	}
	defer rsp.Body.Close()

	return responseError(rsp)
}
