package client

import (
	"context"
	"encoding/json"
	"net/http"
	"time"
)

type UpsertArticleParams struct {
	// AuthorID optionally identifies the (current) author of the article
	AuthorID string `json:"author_id"`

	// ID is the identifier that the company uses for this article
	ID string `json:"id"`

	// Title is the article's title. It may be empty.
	Title string `json:"title"`

	// Description is an article's tagline. It may be empty.
	Description string `json:"description"`

	// Body is the main contents of an article. It may be empty.
	Body string `json:"body"`

	// Visibility describes who can access this article, ranging from the
	// whole world (public) through to employees only (internal).
	Visibility Visibility `json:"visibility"`

	// TopicExternalID optionally identifies the topic that this
	// article is associated with. If given, you must have created
	// the topic first
	TopicID string `json:"topic_id"`

	// Status describes whether this article is published or not.
	Status PublicationStatus `json:"status"`

	// Data optionally gives additional meta-data about the article.
	Data json.RawMessage `json:"data"`

	// Created is when the article was first authored.
	Created time.Time `json:"created"`

	// Updated is when the article was last changed.
	Updated time.Time `json:"updated"`
}

// UpsertArticle inserts or updates a help article
func (c *Client) UpsertArticle(ctx context.Context, p *UpsertArticleParams) error {
	rsp, err := c.makeRequest(ctx, http.MethodPost, "articles", p)
	if err != nil {
		return err
	}
	defer rsp.Body.Close()

	if err := responseError(rsp); err != nil {
		return err
	}
	return nil
}
