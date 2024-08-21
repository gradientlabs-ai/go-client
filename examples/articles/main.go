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
		glabs.WithAPIKey(os.Getenv("GLABS_API_KEY")),
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

	err := client.UpsertArticleTopic(ctx, &glabs.UpsertArticleTopicParams{
		ID:                "topic-1",
		Name:              "Account Management",
		Description:       "How to manage your account in our amazing app!",
		Visibility:        glabs.VisibilityPublic,
		PublicationStatus: glabs.PublicationStatusPublished,
		Created:           time.Now().UTC(),
		LastEdited:        time.Now().UTC(),
	})
	if err != nil {
		return err
	}
	fmt.Println("Created topic: topic-1")

	err = client.UpsertArticleTopic(ctx, &glabs.UpsertArticleTopicParams{
		ID:                "topic-1a",
		ParentID:          "topic-1",
		Name:              "Personal details",
		Description:       "How to change your personal details",
		Visibility:        glabs.VisibilityPublic,
		PublicationStatus: glabs.PublicationStatusPublished,
		Created:           time.Now().UTC(),
		LastEdited:        time.Now().UTC(),
	})
	if err != nil {
		return err
	}
	fmt.Println("Created topic: topic-1a")

	err = client.UpsertArticle(ctx, &glabs.UpsertArticleParams{
		AuthorID:          "neal@gradient-labs.ai",
		ID:                "article-1",
		Title:             "Change my address",
		Body:              "Go to the settings screen in the app, and then tap 'update my address.'",
		Visibility:        glabs.VisibilityPublic,
		TopicID:           "topic-1a",
		PublicationStatus: glabs.PublicationStatusPublished,
		Created:           time.Now().UTC(),
		LastEdited:        time.Now().UTC(),
	})
	if err != nil {
		return err
	}
	fmt.Println("Created article: article-1")
	return nil
}
