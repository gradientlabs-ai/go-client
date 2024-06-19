package client

import (
	"context"
	"fmt"
	"net/http"
)

// AddResource adds (or updates) a resource to the conversation (e.g. the
// customer's order details) so the AI agent can handle customer-specific
// queries.
//
// A resource can be any JSON document, as long it is smaller than 1MB. There
// are no strict requirements on the format/structure of the document, but we
// recommend making attribute names as descriptive as possible.
//
// Over time, the AI agent will learn the structure of your resources - so while
// its fine to add new attributes, you may want to consider using new resource
// names when removing attributes or changing the structure of your resources
// significantly.
//
// Resource names can be anything consisting of letters, numbers, or any of the
// following characters: _ - + =. Names should be descriptive handles that are
// the same for all conversations (e.g. "order-details" and "user-profile") not
// unique identifiers.
func (c *Client) AddResource(ctx context.Context, conversationID string, name string, resource any) error {
	rsp, err := c.makeRequest(ctx, http.MethodPut, fmt.Sprintf("/conversations/%s/resources/%s", conversationID, name), resource)
	if err != nil {
		return err
	}
	defer rsp.Body.Close()

	if err := responseError(rsp); err != nil {
		return err
	}
	return nil
}
