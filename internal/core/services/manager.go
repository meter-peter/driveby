package services

import (
	"context"

	"driveby/internal/config"

	"github.com/sirupsen/logrus"
)

// ServiceManager manages all the services and provides a unified API
type ServiceManager struct {
	config *config.Config
	logger *logrus.Logger

	// Service instances
	validationService ValidationService
	loadTestService   LoadTestService
	acceptanceService AcceptanceTestService
	githubService     GitHubService
}

// NewServiceManager creates a new service manager
func NewServiceManager(cfg *config.Config, logger *logrus.Logger) *ServiceManager {
	manager := &ServiceManager{
		config: cfg,
		logger: logger,
	}

	// Initialize services
	manager.Initialize(context.Background())

	return manager
}

// Initialize initializes all services
func (m *ServiceManager) Initialize(ctx context.Context) error {
	// Initialize GitHub service
	m.githubService = NewGitHubService(m.config, m.logger)
	m.logger.Info("GitHub service initialized")

	// Initialize validation service
	m.validationService = NewValidationService(
		m.config,
		m.logger,
		m.githubService,
	)
	m.logger.Info("Validation service initialized")

	// Initialize load test service (placeholder for now)
	m.loadTestService = nil

	// Initialize acceptance test service (placeholder for now)
	m.acceptanceService = nil

	m.logger.Info("Services initialized")
	return nil
}

// GetValidationService returns the validation service
func (m *ServiceManager) GetValidationService() ValidationService {
	return m.validationService
}

// GetLoadTestService returns the load test service
func (m *ServiceManager) GetLoadTestService() LoadTestService {
	return m.loadTestService
}

// GetAcceptanceService returns the acceptance test service
func (m *ServiceManager) GetAcceptanceService() AcceptanceTestService {
	return m.acceptanceService
}

// GetGitHubService returns the GitHub service
func (m *ServiceManager) GetGitHubService() GitHubService {
	return m.githubService
}
