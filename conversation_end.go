package client

import (
	"context"
	"fmt"
	"net/http"
)

// EndConversation ends a conversation.
func (c *Client) EndConversation(ctx context.Context, conversationID string) error {
	rsp, err := c.makeRequest(ctx, http.MethodPost, fmt.Sprintf("/conversations/%s/end", conversationID), nil)
	if err != nil {
		return err
	}
	defer rsp.Body.Close()

	if err := responseError(rsp); err != nil {
		return err
	}
	return nil
}
