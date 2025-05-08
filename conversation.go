package client

import (
	"time"
)

// Channel represents the way a customer is getting in touch. It will be used to
// determine how the agent formats responses, etc.
type Channel string

const (
	// ChannelWeb indicates the customer is getting in touch via a live instant
	// message chat.
	ChannelWeb Channel = "web"

	// ChannelEmail indicates the customer is getting in touch via an email.
	ChannelEmail Channel = "email"
)

// Status describes the current state of the conversation.
type Status string

const (
	// StatusObserving indicates the agent is following the conversation but not
	// participating (e.g. when it has been handed-off to a human).
	StatusObserving Status = "observing"

	// StatusActive indicates the agent is actively participating in the
	// conversation.
	StatusActive Status = "active"

	// StatusCancelled indicates the conversation has been prematurely brought to
	// a close (e.g. because a human has taken it over) and the AI agent can no
	// longer participate in it.
	StatusCancelled Status = "cancelled"

	// StatusFinished indicates the conversation has been closed because the
	// customer's issue has been resolved.
	StatusFinished Status = "finished"

	// StatusFailed indicates the agent encountered an irrecoverable error, such as
	// not being able to deliver a message to your webhook endpoint after multiple
	// retries.
	StatusFailed Status = "failed"
)

// Conversation represents a series of messages between a customer, human agent,
// and the AI Agent.
type Conversation struct {
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

	// Channel represents the way a customer is getting in touch. It will be used
	// to determine how the agent formats responses, etc.
	Channel Channel `json:"channel"`

	// Metadata is arbitrary metadata that will be attached to the conversation.
	// It will be passed along with webhooks so can be used as action parameters.
	Metadata any `json:"metadata"`

	// Created is the time at which the conversation was created.
	Created time.Time `json:"created"`

	// Updated is the time at which the conversation was last updated.
	Updated time.Time `json:"updated"`

	// Status describes the current state of the conversation.
	Status Status `json:"status"`

	// IsActive is true if the AI agent is currently assigned to response
	// in this conversation.
	IsActive bool `json:"agent_is_active"`

	// AgentMetadata contains several fields that come from the agent.
	AgentMetadata *AgentMetadata `json:"latest_agent_metadata,omitempty"`
}

type AgentMetadata struct {
	// Intent is the name of the *latest* intent that the agent has
	// classified for this conversation.
	Intent string `json:"intent,omitempty"`

	// IntentHandOffTarget is the ID of the hand off target that
	// is currently associated with the latest intent.
	IntentHandOffTarget string `json:"intent_handoff_target,omitempty"`

	// HandOffReason is the coded reason why the agent has handed off
	// the conversation (if at all).
	HandOffReason string `json:"handoff_reason,omitempty"`

	// HandOffNote is the free-text note that the agent generated
	// to summarize what it has done so far when handing off the
	// conversation.
	HandOffNote string `json:"handoff_note,omitempty"`
}
