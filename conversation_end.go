package client

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

type EndParams struct {
	// Timestamp optionally defines the time when the conversation ended.
	// If not given, this will default to the current time.
	Timestamp *time.Time `json:"timestamp,omitempty"`
}

// EndConversation ends a conversation.
func (c *Client) EndConversation(ctx context.Context, conversationID string, p EndParams) error {
	rsp, err := c.makeRequest(ctx, http.MethodPut, fmt.Sprintf("/conversations/%s/end", conversationID), p)
	if err != nil {
		return err
	}
	return c.handleResponse(rsp, nil)
}
