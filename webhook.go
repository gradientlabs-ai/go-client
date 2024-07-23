package client

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
	"time"
)

const (
	defaultLeeway   = 5 * time.Minute
	signatureHeader = "X-GradientLabs-Signature"
)

var (
	// ErrInvalidWebhookSignature is returned when the authenticity of the webhook
	// could not be verified using its signature. You should respond with an HTTP
	// 401 status code.
	ErrInvalidWebhookSignature = fmt.Errorf("%s header is invalid", signatureHeader)

	// ErrUnknownWebhookType is returned when the webhook contained an event of an
	// unknwon type. You should generally log these and return an HTTP 200 status
	// code.
	ErrUnknownWebhookType = errors.New("unknown webhook type")
)

// WebhookType indicates the type of webhook event.
type WebhookType string

const (
	// WebhookTypeAgentMessage indicates the agent wants to send the customer a
	// message.
	WebhookTypeAgentMessage WebhookType = "agent.message"

	// WebhookTypeConversationHandOff indicates the agent is escalating or handing
	// the conversation off to a human agent.
	WebhookTypeConversationHandOff WebhookType = "conversation.hand_off"

	// WebhookTypeConversationFinished indicates the agent has concluded the
	// conversation with the customer and you should close any corresponding
	// ticket, etc.
	WebhookTypeConversationFinished WebhookType = "conversation.finished"
)

// Webhook is an event delivered to your webhook endpoint.
type Webhook struct {
	// ID uniquely identifies this event.
	ID string `json:"id"`

	// Type indicates the type of event.
	Type WebhookType `json:"type"`

	// SequenceNumber can be used to establish an order of webhook events.
	// For more information, see: https://api-docs.gradient-labs.ai/#sequence-numbers
	SequenceNumber int `json:"sequence_number"`

	// Timestamp is the time at which this event was generated.
	Timestamp time.Time `json:"timestamp"`

	// Data contains the event data. Use the helper methods (e.g.
	// Webhook.AgentMessage) to access it.
	Data any `json:"-"`
}

// AgentMessage returns the data for an `agent.message` event.
func (w Webhook) AgentMessage() (*AgentMessageEvent, bool) {
	e, ok := w.Data.(*AgentMessageEvent)
	return e, ok
}

// ConversationHandOff returns the data for an `conversation.hand_off` event.
func (w Webhook) ConversationHandOff() (*ConversationHandOffEvent, bool) {
	e, ok := w.Data.(*ConversationHandOffEvent)
	return e, ok
}

// ConversationFinished returns the data for an `conversation.finished` event.
func (w Webhook) ConversationFinished() (*ConversationFinishedEvent, bool) {
	e, ok := w.Data.(*ConversationFinishedEvent)
	return e, ok
}

// AgentMessageEvent contains the data for an `agent.message` webhook event.
type AgentMessageEvent struct {
	// Conversation contains the details of the conversation the event relates to.
	Conversation WebhookConversation `json:"conversation"`

	// Body contains the text of the message the agent wants to send.
	Body string `json:"body"`
}

// ConversationHandOffEvent contains the data for a `conversation.hand_off` event.
type ConversationHandOffEvent struct {
	// Conversation contains the details of the conversation the event relates to.
	Conversation WebhookConversation `json:"conversation"`

	// Target defines where the agent is wants to hand this conversation to.
	Target string `json:"target,omitempty"`

	// Reason is the code that describes why the agent wants to hand off this
	// conversation.
	Reason string `json:"reason_code"`

	// Description is a human-legible description of the Reason code.
	Description string `json:"reason"`
}

// ConversationFinishedEvent contains the data for a `conversation.finished` event.
type ConversationFinishedEvent struct {
	// Conversation contains the details of the conversation the event relates to.
	Conversation WebhookConversation `json:"conversation"`
}

// WebhookConversation contains the details of the conversation the webhook
// relates to.
type WebhookConversation struct {
	// ID is chosen unique identifier for this conversation.
	ID string `json:"id"`

	// CustomerID is your chosen identifier for the customer.
	CustomerID string `json:"customer_id"`

	// Metadata you attached to the conversation with Client.StartConversation.
	Metadata any `json:"metadata"`
}

// ParseWebhook parses the request, verifies its signature, and returns the
// webhook data.
func (c *Client) ParseWebhook(req *http.Request) (*Webhook, error) {
	if err := c.VerifyWebhookRequest(req); err != nil {
		return nil, err
	}

	var payload struct {
		Webhook

		Data json.RawMessage `json:"data"`
	}
	if err := json.NewDecoder(req.Body).Decode(&payload); err != nil {
		return nil, err
	}

	switch payload.Type {
	case WebhookTypeAgentMessage:
		var am AgentMessageEvent
		if err := json.Unmarshal(payload.Data, &am); err != nil {
			return nil, err
		}
		payload.Webhook.Data = &am
	case WebhookTypeConversationHandOff:
		var ho ConversationHandOffEvent
		if err := json.Unmarshal(payload.Data, &ho); err != nil {
			return nil, err
		}
		payload.Webhook.Data = &ho
	case WebhookTypeConversationFinished:
		var fin ConversationFinishedEvent
		if err := json.Unmarshal(payload.Data, &fin); err != nil {
			return nil, err
		}
		payload.Webhook.Data = &fin
	default:
		return nil, fmt.Errorf("unknown webhook event type received: %q", payload.Type)
	}

	return &payload.Webhook, nil
}

// VerifyWebhookRequest verifies the authenticity of the given request using
// its signature header. You do not need to call it if you're already using
// Client.ParseWebhook.
func (c *Client) VerifyWebhookRequest(req *http.Request) error {
	return c.webhookVerifier.VerifyRequest(req)
}

// WebhookVerifier verifies the authenticity of requests to your webhook
// endpoint using the X-GradientLabs-Signature header.
type WebhookVerifier struct {
	secret []byte
	leeway time.Duration
}

// VerifyRequest verifies the authenticity of the given request using its
// signature header.
func (v WebhookVerifier) VerifyRequest(req *http.Request) error {
	body, err := io.ReadAll(req.Body)
	if err != nil {
		return err
	}
	if err := req.Body.Close(); err != nil {
		return err
	}
	req.Body = io.NopCloser(bytes.NewReader(body))
	return v.VerifySignature(body, req.Header.Get(signatureHeader))
}

// VerifySignature is a lower level variant of VerifyRequest
func (v WebhookVerifier) VerifySignature(body []byte, sig string) error {
	ts, sigs, err := v.parseHeader(sig)
	if err != nil {
		return ErrInvalidWebhookSignature
	}

	if time.Since(ts).Abs() > v.leeway {
		return ErrInvalidWebhookSignature
	}

	expected, err := v.computeSignature(ts, body)
	if err != nil {
		return err
	}

	for _, sig := range sigs {
		if hmac.Equal(expected, sig) {
			return nil
		}
	}

	return ErrInvalidWebhookSignature
}

func (v WebhookVerifier) computeSignature(ts time.Time, body []byte) ([]byte, error) {
	mac := hmac.New(sha256.New, v.secret)

	if _, err := io.WriteString(mac, strconv.Itoa(int(ts.Unix()))); err != nil {
		return nil, err
	}

	if _, err := io.WriteString(mac, "."); err != nil {
		return nil, err
	}

	if _, err := mac.Write(body); err != nil {
		return nil, err
	}

	return mac.Sum(nil), nil
}

func (v WebhookVerifier) parseHeader(header string) (time.Time, [][]byte, error) {
	var (
		ts   time.Time
		sigs [][]byte
	)
	for _, pair := range strings.Split(header, ",") {
		parts := strings.SplitN(pair, "=", 2)
		if len(parts) != 2 {
			return ts, nil, fmt.Errorf("invalid %s header", signatureHeader)
		}

		switch parts[0] {
		case "t":
			unix, err := strconv.ParseInt(parts[1], 10, 64)
			if err != nil {
				return ts, nil, fmt.Errorf("invalid timestamp component %s", parts[1])
			}
			ts = time.Unix(unix, 0)
		case "v1":
			sig, err := hex.DecodeString(parts[1])
			if err != nil {
				return ts, nil, errors.New("invalid signature")
			}
			sigs = append(sigs, sig)
		}
	}

	if ts.IsZero() {
		return ts, nil, fmt.Errorf("%s header contains no timestamp component", signatureHeader)
	}

	return ts, sigs, nil
}
