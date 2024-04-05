package client

import (
	"context"
	"fmt"
	"net/http"
)

type AssignmentParams struct {
	// AssigneeID optionally identifies the specific user that the conversation
	// is being assigned to.
	AssigneeID string `json:"assignee_id,omitempty"`

	// AssigneeType identifies the type of participant that this conversation is
	// being assigned to. Set this to ParticipantTypeAIAgent to assign the conversation
	// to the Gradient Labs AI agent.
	AssigneeType ParticipantType `json:"assignee_type"`
}

// AssignConversation assigns a conversation to a participant.
func (c *Client) AssignConversation(ctx context.Context, conversationID string, p *AssignmentParams) error {
	rsp, err := c.makeRequest(ctx, http.MethodPut, fmt.Sprintf("/conversations/%s/assignee", conversationID), p)
	if err != nil {
		return err
	}
	defer rsp.Body.Close()

	if err := responseError(rsp); err != nil {
		return err
	}
	return nil
}
