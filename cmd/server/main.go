package main

import (
	"context"
	"github.com/aaronland/go-http-maps/app/server"
	"log"
)

func main() {

	ctx := context.Background()
	logger := log.Default()

	err := server.Run(ctx, logger)

	if err != nil {
		logger.Fatalf("Failed to run server application, %v", err)
	}
}
