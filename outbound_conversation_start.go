package client

import (
	"context"
	"encoding/json"
	"net/http"
)

// CustomerSource identifies where customer data originates from.
type CustomerSource string

const (
	CustomerSourceIntercom      CustomerSource = "intercom"
	CustomerSourceFreshchat     CustomerSource = "freshchat"
	CustomerSourceFreshdesk     CustomerSource = "freshdesk"
	CustomerSourcePublicAPI     CustomerSource = "public-api"
	CustomerSourceSalesforce    CustomerSource = "salesforce"
	CustomerSourceZendesk       CustomerSource = "zendesk"
	CustomerSourceVoice         CustomerSource = "livekit"
	CustomerSourceVoiceTwilio   CustomerSource = "twilio"
	CustomerSourceVoiceTalkdesk CustomerSource = "talkdesk"
	CustomerSourceVoiceIntercom CustomerSource = "intercom-voice"
	CustomerSourceWebApp        CustomerSource = "web-app"
	CustomerSourceFile          CustomerSource = "file"
)

// SupportPlatform identifies the support platform where the conversation will be created.
type SupportPlatform string

const (
	SupportPlatformFreshchat  SupportPlatform = "freshchat"
	SupportPlatformFreshdesk  SupportPlatform = "freshdesk"
	SupportPlatformIntercom   SupportPlatform = "intercom"
	SupportPlatformPublicAPI  SupportPlatform = "public-api"
	SupportPlatformSalesforce SupportPlatform = "salesforce"
	SupportPlatformZendesk    SupportPlatform = "zendesk"
)

// StartOutboundConversationParams contains the parameters needed to start a new outbound conversation.
// This kicks off a proactive conversation where your AI agent initiates contact with a customer.
type StartOutboundConversationParams struct {
	// CustomerID is the external identifier for the customer in your support platform.
	// For Intercom, this is the external ID you've defined for the user (e.g., "user-123456").
	// For other platforms, this is the customer identifier used by that platform.
	CustomerID string `json:"customer_id"`

	// CustomerSource is the source of the customer data.
	// For example, a customer ID and phone number might be from Intercom, but the outbound
	// conversation is initiated via Twilio.
	CustomerSource CustomerSource `json:"customer_source"`

	// ProcedureID is the ID of the outbound procedure that defines what the AI agent
	// should accomplish in this conversation. The procedure must be of type "outbound"
	// and must be live (deployed).
	ProcedureID string `json:"procedure_id"`

	// SupportPlatform is the support platform where the conversation should be created.
	// Valid values include "intercom", "zendesk", "freshdesk", "freshchat".
	// If not provided, the system will automatically select the first connected platform
	// in priority order: intercom, zendesk, freshchat, freshdesk, public-api.
	SupportPlatform SupportPlatform `json:"support_platform,omitempty"`

	// Channel specifies the communication channel for this conversation.
	// If not provided, defaults to "email".
	Channel Channel `json:"channel,omitempty"`

	// Subject is the subject line for the initial message (primarily used for email channels).
	// Only used if Body is also provided. If both Subject and Body are omitted, the AI agent
	// will generate the initial message.
	Subject string `json:"subject,omitempty"`

	// Body is the content of the initial message to send to the customer.
	// If provided, this message will be sent instead of having the AI agent generate one.
	// If omitted, the AI agent will generate an appropriate initial message based on the procedure.
	Body string `json:"body,omitempty"`

	// Resources is a JSON object containing structured data that the AI agent
	// can use during the conversation. This should be organized as a map
	// where keys are resource type names and values are the corresponding data.
	// Example: {"customer_profile": {"tier": "premium", "lifetime_value": 5000}}
	// The data will be made available to the AI agent for context during conversation processing.
	Resources json.RawMessage `json:"resources,omitempty"`
}

// StartOutboundConversationResponse contains the response from starting an outbound conversation.
type StartOutboundConversationResponse struct {
	// ConversationID is the internal identifier for the created conversation.
	// You can use this ID with other conversation APIs to check status, send messages, etc.
	ConversationID string `json:"conversation_id"`
}

// StartOutboundConversation creates and starts a new outbound conversation where the AI agent
// proactively initiates contact with a customer. The conversation follows the instructions
// defined in the specified outbound procedure.
//
// If SupportPlatform is not provided, the system will automatically select the highest priority
// platform that has integration settings configured for your company.
//
// If Body and Subject are provided, that message will be sent as the initial message.
// Otherwise, the AI agent will generate an appropriate initial message based on the procedure.
func (c *Client) StartOutboundConversation(ctx context.Context, p StartOutboundConversationParams) (*StartOutboundConversationResponse, error) {
	rsp, err := c.makeRequest(ctx, http.MethodPost, "outbound/conversations", p)
	if err != nil {
		return nil, err
	}
	defer rsp.Body.Close()

	if err := responseError(rsp); err != nil {
		return nil, err
	}

	var result StartOutboundConversationResponse
	if err := json.NewDecoder(rsp.Body).Decode(&result); err != nil {
		return nil, err
	}
	return &result, nil
}
