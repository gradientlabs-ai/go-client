package client

import (
	"context"
	"fmt"
	"net/http"
)

// SetNoteStatusParams contains the parameters for setting a note's status.
type SetNoteStatusParams struct {
	// Status describes whether the note is draft, live, or deleted.
	Status NoteStatus `json:"status"`
}

// SetNoteStatus updates a note's status.
func (c *Client) SetNoteStatus(ctx context.Context, noteID string, p *SetNoteStatusParams) error {
	rsp, err := c.makeRequest(ctx, http.MethodPost, fmt.Sprintf("notes/%s/status", noteID), p)
	if err != nil {
		return err
	}
	defer rsp.Body.Close()

	if err := responseError(rsp); err != nil {
		return err
	}
	return nil
}
