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

	rsp, err := client.ListProcedures(ctx, &glabs.ProcedureListParams{})
	if err != nil {
		return err
	}

	fmt.Printf("Listed %v procedures\n", len(rsp.Procedures))
	for _, proc := range rsp.Procedures {
		prc, err := client.ReadProcedure(ctx, proc.ID)
		if err != nil {
			return err
		}

		fmt.Printf("Procedure (%v, %v): %v (%v)'\n", prc.ID, prc.Status, prc.Name, prc.Description)
	}
	return nil
}
