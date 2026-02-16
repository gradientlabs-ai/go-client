package client

import (
	"context"
	"fmt"
	"net/http"
)

// DeleteNote marks a note as deleted.
func (c *Client) DeleteNote(ctx context.Context, noteID string) error {
	rsp, err := c.makeRequest(ctx, http.MethodDelete, fmt.Sprintf("notes/%s", noteID), nil)
	if err != nil {
		return err
	}
	defer rsp.Body.Close()

	if err := responseError(rsp); err != nil {
		return err
	}
	return nil
}
