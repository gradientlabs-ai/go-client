package client

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type ReadParams struct {
	// SupportPlatform is the name of the support platform where the
	// conversation was started (e.g. Intercom).
	//
	// Leave empty if the conversation was started via the Gradient
	// Labs API.
	SupportPlatform string `json:"support_platform,omitempty"`
}

func (c *Client) ReadConversation(ctx context.Context, conversationID string, p *ReadParams) (*Conversation, error) {
	url := fmt.Sprintf("conversations/%s/read", conversationID)
	if p.SupportPlatform != "" {
		url = fmt.Sprintf("%s?support_platform=%s", url, p.SupportPlatform)
	}

	rsp, err := c.makeRequest(ctx, http.MethodGet, url, nil)
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
