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
	client, err := glabs.NewClient(
		glabs.WithAPIKey(os.Getenv("GLABS_MANAGEMENT_API_KEY")),
	)
	if err != nil {
		log.Fatal(err)
	}

	if err := run(client); err != nil {
		log.Fatal(err)
	}
}

func run(client *glabs.Client) error {
	ctx := context.Background()

	// Create a note with body text
	note, err := client.CreateNote(ctx, &glabs.CreateNoteParams{
		ID:        "note-1",
		Title:     "Product Update: New Features",
		Body:      "We're excited to announce several new features including advanced search and custom workflows.",
		StartTime: timePtr(time.Now().UTC()),
	})
	if err != nil {
		return err
	}
	fmt.Printf("Created note: %s (status: %s)\n", note.ExternalID, note.Status)

	// Update the first note
	updatedNote, err := client.UpdateNote(ctx, note.ExternalID, &glabs.UpdateNoteParams{
		Title: "Product Update: New Features [Updated]",
		Body:  "We're excited to announce several new features including advanced search, custom workflows, and improved analytics.",
	})
	if err != nil {
		return err
	}
	fmt.Printf("Updated note: %s (status: %s)\n", updatedNote.ExternalID, updatedNote.Status)

	// Set note status to draft (i.e., unpublish)
	err = client.SetNoteStatus(ctx, note.ExternalID, &glabs.SetNoteStatusParams{
		Status: glabs.NoteStatusDraft,
	})
	if err != nil {
		return err
	}
	fmt.Printf("Set note %s status to: live\n", note.ExternalID)

	// Delete a note
	err = client.DeleteNote(ctx, note.ExternalID)
	if err != nil {
		return err
	}
	fmt.Printf("Deleted note: %s\n", note.ExternalID)
	return nil
}

func timePtr(t time.Time) *time.Time {
	return &t
}
