package client

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

type CancelParams struct {
	// Timestamp optionally defines the time when the conversation was cancelled.
	// If not given, this will default to the current time.
	Timestamp *time.Time `json:"timestamp,omitempty"`
}

// CancelConversation cancels the conversation.
//
// This is intended for cases where the conversation is being explicitly cancelled or terminated.
// Use FinishConversation() when the conversation is has reached a natural 'end' state, such as it being
// resolved or closed due to inactivity.
func (c *Client) CancelConversation(ctx context.Context, conversationID string, p CancelParams) error {
	rsp, err := c.makeRequest(ctx, http.MethodPut, fmt.Sprintf("conversations/%s/cancel", conversationID), p)
	if err != nil {
		return err
	}
	defer rsp.Body.Close()

	if err := responseError(rsp); err != nil {
		return err
	}
	return nil
}
