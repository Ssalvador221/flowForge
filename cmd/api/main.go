package main

import (
	"FlowForge/internal/observability"
	"context"
	"log"
)

func main() {
	ctx := context.Background()

	shutDown, err := observability.Init(ctx, "flowforge-api")
	if err != nil {
		log.Fatal(err)
	}

	defer shutDown(ctx)

	logger := observability.NewLogger()
	logger.Info("FlowForge API is ON!")
}
