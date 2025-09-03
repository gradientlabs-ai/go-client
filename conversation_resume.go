package client

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

type ConversationResumeParams struct {
	// AssigneeID is the optional external identifier of the participant who should
	// handle this conversation when it is resumed. For AI agents, this can be omitted.
	AssigneeID string `json:"assignee_id,omitempty"`

	// AssigneeType specifies what type of participant should handle this conversation.
	// Valid values include "ai", "human", "bot". Cannot be "customer".
	AssigneeType ParticipantType `json:"assignee_type"`

	// Timestamp is an optional timestamp for when this re-opening occurred, in case
	// it needs to be used for a batch upload. If not provided, the current time will be used.
	Timestamp *time.Time `json:"timestamp,omitempty"`

	// Reason is an optional explanation for why this conversation is re-opening.
	// This can be useful for the conversation timeline.
	Reason string `json:"reason,omitempty"`
}

// ConversationResume re-opens a conversation that was previously finished.
func (c *Client) ResumeConversation(ctx context.Context, conversationID string, p *ConversationResumeParams) error {
	rsp, err := c.makeRequest(ctx, http.MethodPut, fmt.Sprintf("conversations/%s/resume", conversationID), p)
	if err != nil {
		return err
	}
	defer rsp.Body.Close()

	if err := responseError(rsp); err != nil {
		return err
	}
	return nil
}
