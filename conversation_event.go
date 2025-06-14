package client

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

// ConversationEventType describes an event that occurred within the conversation.
type ConversationEventType string

const (
	// ConversationEventInternalNote means that a human agent (or bot) has added
	// a message into the conversation that is not visible to the customer.
	ConversationEventInternalNote ConversationEventType = "internal-note"

	// ConversationEventTypeJoin means the customer or human agent joined the
	// conversation.
	ConversationEventTypeJoin ConversationEventType = "join"

	// ConversationEventTypeLeave means the customer or human agent left the
	// conversation.
	ConversationEventTypeLeave ConversationEventType = "leave"

	// ConversationEventTypeMessageDelivered means that a message has been delivered
	// to a participant
	ConversationEventTypeMessageDelivered ConversationEventType = "delivered"

	// ConversationEventTypeMessageDelivered means that a message has been read
	// by the participant it was delivered to
	ConversationEventTypeMessageRead ConversationEventType = "read"

	// ConversationEventTypeTyping means the customer or human agent started typing.
	ConversationEventTypeTyping ConversationEventType = "typing"
)

type EventParams struct {
	// Type identifies the type of event (see: ConversationEventType).
	Type ConversationEventType `json:"type"`

	// ParticipantID identifies the message sender.
	ParticipantID string `json:"participant_id"`

	// ParticipantType identifies the type of participant who sent this message.
	ParticipantType ParticipantType `json:"participant_type"`

	// MessageID optionally identifies the message that this event relates to
	MessageID *string `json:"message_id,omitempty"`

	// Timestamp optionally defines the time when the conversation was assigned.
	// If not given, this will default to the current time.
	Timestamp *time.Time `json:"timestamp,omitempty"`

	// IdempotencyKey optionally enables you to safely retry requests
	IdempotencyKey string `json:"idempotency_key,omitempty"`
}

// AddConversationEvent records an event such as the customer started typing.
func (c *Client) AddConversationEvent(ctx context.Context, conversationID string, p *EventParams) error {
	rsp, err := c.makeRequest(ctx, http.MethodPost, fmt.Sprintf("conversations/%s/events", conversationID), p)
	if err != nil {
		return err
	}
	defer rsp.Body.Close()

	if err := responseError(rsp); err != nil {
		return err
	}
	return nil
}
