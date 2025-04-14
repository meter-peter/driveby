package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/example/driveby/internal/api"
	"github.com/example/driveby/internal/config"
	"github.com/example/driveby/internal/core/services"
	"github.com/example/driveby/internal/queue"
	"github.com/example/driveby/internal/utils/logging"
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
	logger := logging.NewLogger()
	logger.Info("Starting DriveBy API Service")

	// Load configuration
	cfg, err := config.Load()
	if err != nil {
		logger.WithError(err).Fatal("Failed to load configuration")
	}

	// Create context that listens for termination signals
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer cancel()

	// Initialize queue service
	queueService, err := queue.NewRedisQueue(ctx, cfg.Redis)
	if err != nil {
		logger.WithError(err).Fatal("Failed to initialize queue service")
	}
	defer queueService.Close()

	// Initialize services
	serviceManager := services.NewManager(cfg, queueService, logger.Logger)

	// Initialize API server
	server := api.NewServer(cfg, serviceManager, logger.Logger)

	// Start the server in a separate goroutine
	serverErrors := make(chan error, 1)
	go func() {
		logger.WithField("port", cfg.Server.Port).Info("Starting HTTP server")
		serverErrors <- server.Start()
	}()

	// Start the queue worker processor in a separate goroutine
	go func() {
		logger.Info("Starting queue worker")
		err := serviceManager.StartWorkers(ctx)
		if err != nil {
			logger.WithError(err).Error("Queue worker failed")
		}
	}()

	// Wait for interrupt signal or server error
	select {
	case err := <-serverErrors:
		if err != nil {
			logger.WithError(err).Error("Server error")
		}
	case <-ctx.Done():
		logger.Info("Received termination signal")
	}

	// Graceful shutdown
	logger.Info("Shutting down server...")
	shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), cfg.Server.ShutdownTimeout)
	defer shutdownCancel()

	if err := server.Shutdown(shutdownCtx); err != nil {
		logger.WithError(err).Error("Server shutdown error")
	}

	logger.Info("Server gracefully stopped")
}
