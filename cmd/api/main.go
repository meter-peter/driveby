package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/sirupsen/logrus"

	"driveby/internal/api"
	"driveby/internal/config"
	"driveby/internal/core"
	"driveby/internal/core/services"
)

func main() {
	// Initialize logger
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})
	logger.SetLevel(logrus.InfoLevel)

	// Get configuration from environment
	apiHost := getEnv("API_HOST", "localhost")
	apiPort := getEnv("API_PORT", "8080")
	apiBasePath := getEnv("API_BASE_PATH", "")

	// Load config
	cfg, err := config.LoadConfig("")
	if err != nil {
		logger.WithError(err).Fatal("Failed to load config")
	}

	// Initialize service manager
	manager := services.NewServiceManager(cfg, logger)

	// Initialize testing service
	testingSvc := core.NewTestingService(logger, apiHost, apiPort)

	// Initialize API server
	server := api.NewServer(logger, testingSvc, apiHost, apiPort, apiBasePath, cfg, manager)

	// Start server in a goroutine
	go func() {
		if err := server.Start(); err != nil {
			logger.WithError(err).Fatal("Failed to start server")
		}
	}()

	// Wait for interrupt signal
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	// Create shutdown context
	ctx, cancel := context.WithTimeout(context.Background(), 5)
	defer cancel()

	// Shutdown server
	if err := server.Shutdown(ctx); err != nil {
		logger.WithError(err).Fatal("Server forced to shutdown")
	}

	logger.Info("Server exiting")
}

// getEnv gets an environment variable or returns a default value
func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}
