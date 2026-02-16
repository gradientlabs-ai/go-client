# Gradient Labs Go
[![Go Reference](https://pkg.go.dev/badge/github.com/gradientlabs-ai/go-client.svg)](https://pkg.go.dev/github.com/gradientlabs-ai/go-client)

Go bindings for the [Gradient Labs API](https://api-docs.gradient-labs.ai).

## Requirements

- Go 1.20 or later

## Installation

```bash
go get github.com/gradientlabs-ai/go-client
```

## Documentation

- [API Documentation](https://api-docs.gradient-labs.ai)
- [Go Package Reference](https://pkg.go.dev/github.com/gradientlabs-ai/go-client)

## Example Usage

### Starting a Conversation

```go
package main

import (
    "context"
    "fmt"
    "log"
    "os"
    "time"

    glabs "github.com/gradientlabs-ai/go-client"
)

func main() {
    // Create a new client with your API key
    client, err := glabs.NewClient(
        glabs.WithAPIKey(os.Getenv("GLABS_API_KEY")),
    )
    if err != nil {
        log.Fatal(err)
    }

    ctx := context.Background()

    // Start a new conversation
    conv, err := client.StartConversation(ctx, glabs.StartConversationParams{
        ID:         "conversation-1234",
        CustomerID: "user-1234",
        Channel:    glabs.ChannelWeb,
        Resources: map[string]any{
            "user_profile": map[string]any{
                "name":         "Jane Doe",
                "subscription": "premium",
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("Started conversation: %s\n", conv.ID)

    // Add a message to the conversation
    msg, err := client.AddMessage(ctx, conv.ID, glabs.AddMessageParams{
        ID:              "message-1234",
        Body:            "Hello! I need some help.",
        ParticipantID:   "user-1234",
        ParticipantType: glabs.ParticipantTypeCustomer,
        Created:         time.Now(),
    })
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("Added message: %s\n", msg.ID)

    // Assign the conversation to an AI agent
    err = client.AssignConversation(ctx, conv.ID, &glabs.AssignmentParams{
        AssigneeType: glabs.ParticipantTypeAIAgent,
    })
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println("Conversation assigned to AI agent")
}
```

For more examples, see the [examples](./examples) directory.
