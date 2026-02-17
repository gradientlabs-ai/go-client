package client

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type ReadTopicParams struct {
	// SupportPlatform specifies which platform's topic to read.
	// If not provided, defaults to "public-api".
	SupportPlatform SupportPlatform `json:"support_platform,omitempty"`
}

// ReadTopic reads an article topic by ID.
func (c *Client) ReadTopic(ctx context.Context, topicID string, p *ReadTopicParams) (*Topic, error) {
	path := fmt.Sprintf("topic/%s", topicID)
	if p != nil && p.SupportPlatform != "" {
		path = fmt.Sprintf("topic/%s?support_platform=%s", topicID, p.SupportPlatform)
	}

	rsp, err := c.makeRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}
	defer rsp.Body.Close()

	if err := responseError(rsp); err != nil {
		return nil, err
	}

	var topic Topic
	if err := json.NewDecoder(rsp.Body).Decode(&topic); err != nil {
		return nil, err
	}

	return &topic, nil
}
