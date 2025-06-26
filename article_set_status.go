package client

import (
	"context"
	"fmt"
	"net/http"
)

type UsageStatus string

const (
	UsageStatusOn  UsageStatus = "on"
	UsageStatusOff UsageStatus = "off"
)

type SetArticleUsageStatusParams struct {
	UsageStatus UsageStatus `json:"usage_status"`
}

// SetArticleUsageStatus updates an article's usage status. Use this to
// make it (un)available for use by the AI agent.
func (c *Client) SetArticleUsageStatus(ctx context.Context, articleID string, p *SetArticleUsageStatusParams) error {
	rsp, err := c.makeRequest(ctx, http.MethodPost, fmt.Sprintf("articles/%s/usage-status", articleID), p)
	if err != nil {
		return err
	}
	defer rsp.Body.Close()

	if err := responseError(rsp); err != nil {
		return err
	}
	return nil
}
