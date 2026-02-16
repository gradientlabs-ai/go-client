package client

import (
	"context"
	"encoding/json"
	"net/http"
	"time"
)

// CreateNoteParams contains the parameters for creating a note.
type CreateNoteParams struct {
	// AuthorID optionally identifies the (current) author of the note.
	AuthorID string `json:"author_id,omitempty"`

	// ID is your identifier for this note.
	ID string `json:"id"`

	// Title is the note's title.
	Title string `json:"title"`

	// Body is the main contents of a note. This is mutually exclusive with WebpageURL.
	Body string `json:"body,omitempty"`

	// WebpageURL optionally points to a webpage to use as the note body.
	// This is mutually exclusive with Body.
	WebpageURL string `json:"webpage_url,omitempty"`

	// StartTime is when the note becomes relevant.
	StartTime *time.Time `json:"start_time,omitempty"`

	// EndTime is when the note is no longer relevant.
	EndTime *time.Time `json:"end_time,omitempty"`
}

// CreateNote creates a new note.
func (c *Client) CreateNote(ctx context.Context, p *CreateNoteParams) (*Note, error) {
	rsp, err := c.makeRequest(ctx, http.MethodPost, "notes", p)
	if err != nil {
		return nil, err
	}
	defer rsp.Body.Close()

	if err := responseError(rsp); err != nil {
		return nil, err
	}

	var note Note
	if err := json.NewDecoder(rsp.Body).Decode(&note); err != nil {
		return nil, err
	}
	return &note, nil
}
