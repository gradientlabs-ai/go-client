package client

import (
	"context"
	"fmt"
	"net/http"
)

// DeleteResourceSource deletes a resource source by its ID.
//
// Note: requires a `Management` API key.
func (c *Client) DeleteResourceSource(ctx context.Context, id string) error {
	rsp, err := c.makeRequest(ctx, http.MethodDelete, fmt.Sprintf("resource-sources/%s", id), nil)
	if err != nil {
		return err
	}
	defer rsp.Body.Close()

	return responseError(rsp)
}
