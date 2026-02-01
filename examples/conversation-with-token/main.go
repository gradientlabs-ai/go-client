package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	glabs "github.com/gradientlabs-ai/go-client"
)

// This example is similar to ./conversation/main.go, but it makes use of Conversation-Scoped tokens.
// When a conversation is started, a token is provided in the request.
// When handling a webhook event, we will expect that token to be echoed back and we will perform
// further validations on the token.

func main() {
	client, err := glabs.NewClient(
		glabs.WithAPIKey(os.Getenv("GLABS_API_KEY")),
		glabs.WithWebhookSigningKey(os.Getenv("GLABS_WEBHOOK_KEY")),
	)
	if err != nil {
		log.Fatal(err)
	}

	if err := run(client); err != nil {
		log.Fatal(err)
	}

	if err := http.ListenAndServe(":4321", webhookHandler(client)); err != nil {
		log.Fatal(err)
	}
}

type token struct {
	payload        string
	conversationID string
	userID         string
	expiry         time.Time
}

const tokenPayload = "123456789"
const customerID = "user-1234"
const conversationID = "conversation-1234"

// Set-up a mock database of conversation-scoped tokens
var conversationTokensDatabase = map[string]token{
	tokenPayload: {
		payload:        tokenPayload,
		conversationID: conversationID,
		userID:         customerID,
		expiry:         time.Now().Add(1 * time.Hour),
	},
}

func run(client *glabs.Client) error {
	ctx := context.Background()

	conv, err := client.StartConversation(ctx, glabs.StartConversationParams{
		ID:         conversationID,
		CustomerID: customerID,
		Channel:    glabs.ChannelWeb,
		Resources: map[string]any{
			"user_profile": map[string]any{
				"name":         "Jane Doe",
				"subscription": "premium",
			},
			"transaction": map[string]any{
				"id":       123,
				"outbound": true,
			},
			"source": "website",
		},
		// Include token when starting a conversation
		ConversationToken: tokenPayload,
	})
	if err != nil {
		return err
	}
	fmt.Printf("Conversation: %#v\n", conv)

	msg, err := client.AddMessage(ctx, conv.ID, glabs.AddMessageParams{
		ID:              "message-1234",
		Body:            "Hello! I need some help setting up my toaster oven, please.",
		ParticipantID:   "user-1234",
		ParticipantType: glabs.ParticipantTypeCustomer,
		Created:         time.Now(),
		Attachments: []*glabs.Attachment{
			{
				Type:     glabs.AttachmentTypeImage,
				FileName: "toaster.jpg",
			},
		},
	})
	if err != nil {
		return err
	}
	fmt.Printf("Message: %#v\n", msg)

	err = client.AssignConversation(ctx, conv.ID, &glabs.AssignmentParams{
		AssigneeType: glabs.ParticipantTypeAIAgent,
	})
	if err != nil {
		return err
	}

	err = client.AddConversationEvent(ctx, conv.ID, &glabs.EventParams{
		Type:            glabs.ConversationEventInternalNote,
		ParticipantID:   "user-1234",
		ParticipantType: glabs.ParticipantTypeCustomer,
		Body:            "This customer has bought a toaster from someone else",
	})
	if err != nil {
		return err
	}

	if err := client.FinishConversation(ctx, conv.ID, glabs.FinishParams{}); err != nil {
		return err
	}

	readRsp, err := client.ReadConversation(ctx, conv.ID, &glabs.ReadParams{})
	if err != nil {
		return err
	}
	fmt.Printf("%#v\n", readRsp)

	return nil
}

func webhookHandler(client *glabs.Client) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		webhook, token, err := client.ParseWebhook(r)
		switch {
		case errors.Is(err, glabs.ErrInvalidWebhookSignature):
			w.WriteHeader(http.StatusUnauthorized)
			return
		case errors.Is(err, glabs.ErrUnknownWebhookType):
			log.Printf("unknown webhook type: %q", webhook.Type)
			return
		case err != nil:
			w.WriteHeader(http.StatusInternalServerError)
			log.Printf("failed to parse webhook: %v", err)
			return
		}

		// Validate the conversation-scoped token
		if !isValidConversationToken(token) {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		if am, ok := webhook.AgentMessage(); ok {
			log.Printf("agent message: %s", am.Body)
			return
		}

		if ho, ok := webhook.ConversationHandOff(); ok {
			log.Printf("hand off: %s", ho.Conversation.ID)
			return
		}

		if fin, ok := webhook.ConversationFinished(); ok {
			log.Printf("hand off: %s", fin.Conversation.ID)
			return
		}

	})
}

func isValidConversationToken(token string) bool {
	if token != tokenPayload {
		// Webhook returned a token we did not expect
		return false
	}
	tokenData, ok := conversationTokensDatabase[token]
	if !ok {
		return false
	}

	switch {
	case tokenData.userID != customerID:
		// Token is for a different customer
		return false
	case tokenData.conversationID != conversationID:
		// Token is for a different conversation
		return false
	case tokenData.expiry.Before(time.Now()):
		// Token has expired
		return false
	}

	return true
}
