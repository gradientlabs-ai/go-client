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

	procedureID := os.Getenv("PROCEDURE_ID")

	rsp, err := client.ListProcedureVersions(ctx, procedureID)
	if err != nil {
		return err
	}

	if len(rsp.Versions) == 0 {
		fmt.Println("No versions found")
		return nil
	}

	fmt.Printf("Listed %v procedure versions\n", len(rsp.Versions))
	for _, proc := range rsp.Versions {
		fmt.Printf("Version: %v; Live: %t, Gated: %t\n", proc.Version, proc.Live, proc.Gated)
	}

	ver := rsp.Versions[0]

	fmt.Println("Set gated version")
	err = client.SetProcedureGatedVersion(ctx, procedureID, ver.Version, &glabs.SetProcedureGatedVersionParams{
		MaxDailyConversations: 10,
		Replace:               true,
	})
	if err != nil {
		return err
	}

	fmt.Println("Unset gated version")
	err = client.UnsetProcedureGatedVersion(ctx, procedureID, ver.Version)
	if err != nil {
		return err
	}

	fmt.Println("Set live version")
	err = client.SetProcedureLiveVersion(ctx, procedureID, ver.Version)
	if err != nil {
		return err
	}

	fmt.Println("Unset live version")
	err = client.UnsetProcedureLiveVersion(ctx, procedureID, ver.Version)
	if err != nil {
		return err
	}

	return nil
}
