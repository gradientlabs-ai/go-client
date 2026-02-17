package client

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// ReturnAsyncToolResultParams are the parameters to Client.ReturnAsyncToolResult.
type ReturnAsyncToolResultParams struct {
	// AsyncToolExecutionID is the unique identifier for the async tool execution.
	// This ID is provided in the action.execute webhook event when the agent
	// requests the tool execution.
	AsyncToolExecutionID string `json:"async_tool_execution_id"`

	// Payload is the result data from the tool execution as a JSON object.
	// The structure of this object depends on your tool's implementation.
	// The agent will use AI to extract relevant information from any well-formed
	// JSON object.
	Payload json.RawMessage `json:"payload"`

	// Timestamp optionally defines the time when the result was generated.
	// If not given, this will default to the current time.
	Timestamp *time.Time `json:"timestamp,omitempty"`
}

// ReturnAsyncToolResult returns the result of an async tool execution.
//
// When a tool is configured for asynchronous execution, the agent will request
// the tool execution via an action.execute webhook event, and your system should
// return the result by calling this method.
//
// This allows your system to perform long-running operations without blocking the
// conversation, and return results when they're ready.
//
// Important Notes:
//   - The conversation must be in an ongoing state. You cannot return async tool
//     results to conversations that are finished, failed, or cancelled.
//   - Make sure to use the correct AsyncToolExecutionID from the action.execute
//     webhook event.
//   - The result payload should be a valid JSON object containing the data the
//     agent needs to continue the conversation.
func (c *Client) ReturnAsyncToolResult(ctx context.Context, conversationID string, p ReturnAsyncToolResultParams) error {
	rsp, err := c.makeRequest(ctx, http.MethodPut, fmt.Sprintf("conversations/%s/return-async-tool-result", conversationID), p)
	if err != nil {
		return err
	}
	defer rsp.Body.Close()

	if err := responseError(rsp); err != nil {
		return err
	}
	return nil
}
