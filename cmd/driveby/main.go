package main

import (
	"context"
	"driveby/internal/api"
	"driveby/internal/config"
	"driveby/internal/core"
	"driveby/internal/core/services"
	"os"
	"os/signal"
	"syscall"

	"github.com/sirupsen/logrus"
)

// @title           DriveBy API
// @version         1.0
// @description     API for testing microservices, validating documentation and performing load tests
// @termsOfService  http://example.com/terms/

// @contact.name   API Support
// @contact.url    http://www.example.com/support
// @contact.email  support@example.com

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /api/v1

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization

func main() {
	// Initialize logger
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})

	// Load configuration
	cfg, err := config.LoadConfig("")
	if err != nil {
		logger.WithError(err).Fatal("Failed to load configuration")
	}

	// Set log level
	level, err := logrus.ParseLevel(cfg.LogLevel)
	if err != nil {
		logger.WithError(err).Fatal("Failed to parse log level")
	}
	logger.SetLevel(level)

	// Initialize service manager
	serviceManager := services.NewServiceManager(cfg, logger)
	if err := serviceManager.Initialize(context.Background()); err != nil {
		logger.WithError(err).Fatal("Failed to initialize services")
	}

	// Initialize testing service
	testingSvc := core.NewTestingService(logger, "0.0.0.0", "8081")

	// Initialize API server
	server := api.NewServer(logger, testingSvc, "0.0.0.0", "8081", "/api/v1", cfg, serviceManager)

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
