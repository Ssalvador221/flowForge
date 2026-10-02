package worker

import (
	"FlowForge/internal/observability"
	"context"
	"log"
)

func main() {
	ctx := context.Background()

	shutDown, err := observability.Init(ctx, "flowForge-worker")
	if err != nil {
		log.Fatal(err)
	}

	defer shutDown(ctx)

	logger := observability.NewLogger()

	logger.Info("Worker is ON!")
}
