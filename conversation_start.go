package client

import (
	"context"
	"net/http"
	"time"
)

// StartConversationParams are the parameters to Client.StartConversation.
type StartConversationParams struct {
	// ID uniquely identifies the conversation.
	//
	// Can be anything consisting of letters, numbers, or any of the following
	// characters: _ - + =.
	//
	// Tip: use something meaningful to your business (e.g. a ticket number).
	ID string `json:"id"`

	// CustomerID uniquely identifies the customer. Used to build historical
	// context of conversations the agent has had with this customer.
	CustomerID string `json:"customer_id"`

	// AssigneeID optionally identifies who the conversation is assigned to.
	AssigneeID string `json:"assignee_id,omitempty"`

	// AssigneeType optionally identifies which type of participant is currently
	// assigned to respond. Set this to ParticipantTypeAIAgent to assign the conversation
	// to the Gradient Labs AI when starting it.
	AssigneeType ParticipantType `json:"assignee_type,omitempty"`

	// Channel represents the way a customer is getting in touch. It will be used
	// to determine how the agent formats responses, etc.
	Channel Channel `json:"channel"`

	// Metadata is arbitrary metadata that will be attached to the conversation.
	// It will be passed along with webhooks so can be used as action parameters.
	Metadata any `json:"metadata"`

	// Created optionally defines the time when the conversation started.
	// If not given, this will default to the current time.
	Created *time.Time `json:"created,omitempty"`
}

// StartConversation begins a conversation.
func (c *Client) StartConversation(ctx context.Context, p StartConversationParams) (*Conversation, error) {
	rsp, err := c.makeRequest(ctx, http.MethodPost, "conversations", p)
	if err != nil {
		return nil, err
	}
	var conv Conversation
	if err := c.handleResponse(rsp, &conv); err != nil {
		return nil, err
	}
	return &conv, nil
}
