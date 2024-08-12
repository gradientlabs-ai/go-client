package client

import (
	"context"
	"encoding/json"
	"net/http"
	"time"
)

type UpsertArticleTopicParams struct {
	// ID is the identifier that the company uses for this topic
	ID string `json:"id"`

	// ParentID is the identifier for this topic's parent topic (if any).
	ParentID string `json:"parent_id"`

	// Name is the topic's name.
	Name string `json:"title"`

	// Description is an topic's tagline. It may be empty.
	Description string `json:"description"`

	// Visibility describes who can see this topic, ranging from the
	// whole world (public) through to employees only (internal).
	Visibility Visibility `json:"visibility"`

	// Status describes whether this article is published or not.
	Status PublicationStatus `json:"status"`

	// Data optionally gives additional meta-data about the topic.
	Data json.RawMessage `json:"data"`

	// Created is when the topic was first created.
	Created time.Time `json:"created"`

	// Updated is when the topic was last changed.
	Updated time.Time `json:"updated"`
}

func (c *Client) UpsertArticleTopic(ctx context.Context, p *UpsertArticleTopicParams) error {
	rsp, err := c.makeRequest(ctx, http.MethodPost, "topics", p)
	if err != nil {
		return err
	}
	defer rsp.Body.Close()

	if err := responseError(rsp); err != nil {
		return err
	}
	return nil
}
