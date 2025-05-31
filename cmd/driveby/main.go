package main

import (
	"os"

	"github.com/meter-peter/driveby/internal/cli"
	"github.com/meter-peter/driveby/internal/logger"
)

func main() {
	// Initialize logger with defaults
	if err := logger.Configure(logger.DefaultConfig()); err != nil {
		logger.Fatalf("Failed to configure logger: %v", err)
	}

	// Execute CLI
	if err := cli.Execute(); err != nil {
		logger.WithError(err).Fatal("Application error")
		os.Exit(1)
	}
}
