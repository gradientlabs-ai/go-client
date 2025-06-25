package client

import (
	"context"
	"fmt"
	"net/http"
)

// DeleteArticle marks an article as deleted. Copies of the article are kept
// in case they are needed to render citations.
func (c *Client) DeleteArticle(ctx context.Context, articleID string) error {
	rsp, err := c.makeRequest(ctx, http.MethodDelete, fmt.Sprintf("articles/%s", articleID), nil)
	if err != nil {
		return err
	}
	defer rsp.Body.Close()

	if err := responseError(rsp); err != nil {
		return err
	}
	return nil
}
