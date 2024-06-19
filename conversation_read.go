package client

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

// ReadConversation returns a conversation.
func (c *Client) ReadConversation(ctx context.Context, conversationID string) (*Conversation, error) {
	rsp, err := c.makeRequest(ctx, http.MethodGet, fmt.Sprintf("conversations/%s", conversationID), nil)
	if err != nil {
		return nil, err
	}
	defer rsp.Body.Close()

	if err := responseError(rsp); err != nil {
		return nil, err
	}

	var conv Conversation
	if err := json.NewDecoder(rsp.Body).Decode(&conv); err != nil {
		return nil, err
	}
	return &conv, nil
}
