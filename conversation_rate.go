package client

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

type RatingParams struct {

	// SurveyType identifies the type of survey sent to customers.
	SurveyType string `json:"type"`

	// Value is the rating value that was submitted by the customer.
	// It must be a value between [MinValue, MaxValue].
	Value int `json:"value"`

	// MaxValue is the maximum value in the rating scale.
	MaxValue int `json:"max_value"`

	// MinValue is the minimum value of the rating scale.
	MinValue int `json:"min_value"`

	// Comments optionally submits any free-text that was submitted by the
	// customer alongside their rating.
	Comments string `json:"comments,omitempty"`

	// Timestamp optionally defines the time when the conversation was rate.
	// If not given, this will default to the current time.
	Timestamp *time.Time `json:"timestamp,omitempty"`
}

// RateConversation submits a customer (CSAT) rating for a conversation.
func (c *Client) RateConversation(ctx context.Context, conversationID string, p *AssignmentParams) error {
	rsp, err := c.makeRequest(ctx, http.MethodPut, fmt.Sprintf("conversations/%s/rate", conversationID), p)
	if err != nil {
		return err
	}
	defer rsp.Body.Close()

	if err := responseError(rsp); err != nil {
		return err
	}
	return nil
}
