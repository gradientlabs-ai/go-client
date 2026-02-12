package client

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// UpdateNoteParams contains the parameters for updating a note.
type UpdateNoteParams struct {
	// AuthorID optionally identifies the (current) author of the note.
	AuthorID string `json:"author_id,omitempty"`

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

// UpdateNote updates an existing note's contents.
func (c *Client) UpdateNote(ctx context.Context, noteID string, p *UpdateNoteParams) (*Note, error) {
	rsp, err := c.makeRequest(ctx, http.MethodPost, fmt.Sprintf("notes/%s", noteID), p)
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
