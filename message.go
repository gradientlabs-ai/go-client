package client

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// ParticipantType identifies the type of participant who sent a message.
type ParticipantType string

const (
	// ParticipantTypeCustomer indicates that the message was sent by a
	// customer/end-user.
	ParticipantTypeCustomer ParticipantType = "Customer"

	// ParticipantTypeHumanAgent indicates that the message was sent by a human support
	// agent.
	ParticipantTypeHumanAgent ParticipantType = "Agent"

	// ParticipantTypeBot indicates that the message was sent by an automation/bot
	// other than the Gradient Labs AI agent.
	ParticipantTypeBot ParticipantType = "Bot"
)

// AttachmentType identifies the type of file that has been attached to a
// message. Currently, our AI agent does not support processing attachments,
// and will hand the conversation off to a human agent if it encounters one.
type AttachmentType string

const (
	// AttachmentTypeImage indicates that the attachment is an image.
	AttachmentTypeImage AttachmentType = "image"

	// AttachmentTypeFile indicates that the attachment is a generic file such as
	// a document.
	AttachmentTypeFile AttachmentType = "file"
)

// Attachment represents a file that was uploaded with a message.
type Attachment struct {
	// Type of file that was uploaded.
	Type AttachmentType `json:"type"`

	// FileName is the name of the file that was uploaded.
	FileName string `json:"file_name"`
}

// AddMessageParams are the parameters to Client.AddMessage.
type AddMessageParams struct {
	// ID uniquely identifies this message within the conversation.
	//
	// Can be anything consisting of letters, numbers, or any of the following
	// characters: _ - + =.
	//
	// Tip: use something meaningful to your business.
	ID string `json:"id"`

	// Body contains the message text.
	Body string `json:"body"`

	// ParticipantID identifies the message sender.
	ParticipantID string `json:"participant_id"`

	// ParticipantType identifies the type of participant who sent this message.
	ParticipantType ParticipantType `json:"participant_type"`

	// Created is the time at which the message was sent.
	Created time.Time `json:"created"`

	// Metadata is arbitrary metadata that will be attached to the message.
	Metadata any `json:"metadata"`

	// Attachments contains any files that were uploaded with this message.
	Attachments []*Attachment `json:"attachments,omitempty"`
}

// Message represents a message sent by a customer, human support agent, or the
// API agent.
type Message struct {
	// ID uniquely identifies this message within the conversation.
	//
	// Can be anything consisting of letters, numbers, or any of the following
	// characters: _ - + =.
	//
	// Tip: use something meaningful to your business.
	ID string `json:"id"`

	// Body contains the message text.
	Body string `json:"body"`

	// ParticipantID identifies the message sender.
	ParticipantID string `json:"participant_id"`

	// ParticipantType identifies the type of participant who sent this message.
	ParticipantType ParticipantType `json:"participant_type"`

	// Created is the time at which the message was sent.
	Created time.Time `json:"created"`

	// Metadata is arbitrary metadata attached to the message.
	Metadata any `json:"metadata"`

	// Attachments contains any files that were uploaded with this message.
	Attachment []*Attachment `json:"attachments,omitempty"`
}

// AddMessage records a message sent by the customer or a human agent.
func (c *Client) AddMessage(ctx context.Context, conversationID string, p AddMessageParams) (*Message, error) {
	rsp, err := c.makeRequest(ctx, http.MethodPost, fmt.Sprintf("conversations/%s/messages", conversationID), p)
	if err != nil {
		return nil, err
	}
	defer rsp.Body.Close()

	if rsp.StatusCode < 200 || rsp.StatusCode > 299 {
		return nil, newResponseError(rsp)
	}

	var msg Message
	if err := json.NewDecoder(rsp.Body).Decode(&msg); err != nil {
		return nil, err
	}
	return &msg, nil
}
