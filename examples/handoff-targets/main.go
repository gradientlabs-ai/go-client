package main

import (
	"context"
	"fmt"
	"log"
	"os"

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

	rsp, err := client.ListHandOffTargets(ctx)
	if err != nil {
		return err
	}

	fmt.Printf("Listed %v handoff targets\n", len(rsp.Targets))
	for _, tgt := range rsp.Targets {
		fmt.Printf("\tTarget id: %v", tgt.ID)
	}
	return nil
}
