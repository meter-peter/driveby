package api

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/example/driveby/internal/config"
	"github.com/example/driveby/internal/core/models"
	"github.com/example/driveby/internal/core/services"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// Server represents the API server
type Server struct {
	router  *gin.Engine
	server  *http.Server
	config  *config.Config
	logger  *logrus.Logger
	manager *services.ServiceManager
}

// NewServer creates a new API server
func NewServer(cfg *config.Config, manager *services.ServiceManager, logger *logrus.Logger) *Server {
	// Set Gin mode
	switch cfg.Server.Mode {
	case "debug":
		gin.SetMode(gin.DebugMode)
	case "release":
		gin.SetMode(gin.ReleaseMode)
	case "test":
		gin.SetMode(gin.TestMode)
	default:
		gin.SetMode(gin.ReleaseMode)
	}

	// Create router
	router := gin.New()

	// Setup middleware
	router.Use(
		gin.Recovery(),
		loggerMiddleware(logger),
		corsMiddleware(),
	)

	server := &Server{
		router:  router,
		config:  cfg,
		logger:  logger,
		manager: manager,
	}

	// Setup routes
	server.setupRoutes()

	// Create HTTP server
	server.server = &http.Server{
		Addr:         fmt.Sprintf("%s:%d", cfg.Server.Host, cfg.Server.Port),
		Handler:      router,
		ReadTimeout:  cfg.Server.Timeout,
		WriteTimeout: cfg.Server.Timeout,
		IdleTimeout:  2 * cfg.Server.Timeout,
	}

	return server
}

// Start starts the server
func (s *Server) Start() error {
	s.logger.WithField("addr", s.server.Addr).Info("Starting HTTP server")
	if err := s.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		return err
	}
	return nil
}

// Shutdown gracefully shuts down the server
func (s *Server) Shutdown(ctx context.Context) error {
	s.logger.Info("Shutting down HTTP server")
	return s.server.Shutdown(ctx)
}

// setupRoutes sets up the API routes
func (s *Server) setupRoutes() {
	// API version grouping
	api := s.router.Group("/api/v1")

	// Health check
	api.GET("/health", s.healthCheckHandler)

	// Documentation routes
	api.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Validation routes
	validation := api.Group("/validation")
	{
		validation.POST("", s.createValidationHandler)
		validation.GET("", s.listValidationsHandler)
		validation.GET("/:id", s.getValidationHandler)
		validation.GET("/:id/report", s.getValidationReportHandler)
	}

	// Load test routes
	loadtest := api.Group("/loadtest")
	{
		loadtest.POST("", s.createLoadTestHandler)
		loadtest.GET("", s.listLoadTestsHandler)
		loadtest.GET("/:id", s.getLoadTestHandler)
		loadtest.GET("/:id/report", s.getLoadTestReportHandler)
	}

	// Acceptance test routes
	acceptance := api.Group("/acceptance")
	{
		acceptance.POST("", s.createAcceptanceTestHandler)
		acceptance.GET("", s.listAcceptanceTestsHandler)
		acceptance.GET("/:id", s.getAcceptanceTestHandler)
		acceptance.GET("/:id/report", s.getAcceptanceTestReportHandler)
	}
}

// authMiddleware handles API authentication
func authMiddleware(logger *logrus.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get authorization header
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Authorization header required",
			})
			return
		}

		// TODO: Implement proper JWT or OAuth authentication
		// For now, just check for a token prefix
		if len(authHeader) < 8 || authHeader[:7] != "Bearer " {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Invalid authorization format",
			})
			return
		}

		// Set user info in context
		c.Set("user_id", "demo-user") // Replace with actual user ID from token

		c.Next()
	}
}

// loggerMiddleware logs HTTP requests
func loggerMiddleware(logger *logrus.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Start timer
		start := time.Now()
		path := c.Request.URL.Path
		method := c.Request.Method

		// Add request ID
		requestID := c.GetHeader("X-Request-ID")
		if requestID == "" {
			requestID = uuid.New().String()
			c.Header("X-Request-ID", requestID)
		}

		// Process request
		c.Next()

		// Calculate latency
		latency := time.Since(start)

		// Log request
		logger.WithFields(logrus.Fields{
			"status":      c.Writer.Status(),
			"method":      method,
			"path":        path,
			"ip":          c.ClientIP(),
			"latency":     latency,
			"request_id":  requestID,
			"user_agent":  c.Request.UserAgent(),
			"error_count": len(c.Errors),
		}).Info("HTTP request")
	}
}

// corsMiddleware handles CORS headers
func corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, X-Request-ID")
		c.Writer.Header().Set("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, X-Request-ID")
		c.Writer.Header().Set("Access-Control-Max-Age", "86400")

		// Handle preflight requests
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		c.Next()
	}
}

// healthCheckHandler handles health check requests
func (s *Server) healthCheckHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":    "ok",
		"timestamp": time.Now().Format(time.RFC3339),
		"version":   "1.0.0",
	})
}

// createValidationHandler creates a new validation test
func (s *Server) createValidationHandler(c *gin.Context) {
	var request models.ValidationRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Set defaults if not provided
	if request.ComplianceThreshold == nil {
		threshold := s.config.Testing.Validation.ComplianceThreshold
		request.ComplianceThreshold = &threshold
	}
	if request.FailOnValidation == nil {
		failOnValidation := s.config.Testing.Validation.FailOnValidation
		request.FailOnValidation = &failOnValidation
	}

	// Create test
	test := models.NewValidationTest(
		request.Name,
		request.Description,
		request.OpenAPIURL,
		*request.ComplianceThreshold,
	)
	test.FailOnValidation = *request.FailOnValidation
	test.Tags = request.Tags

	// Setup GitHub issue creation if requested
	if request.CreateGitHubIssue && request.GitHubRepo != nil {
		if request.GitHubRepo.Owner == "" {
			request.GitHubRepo.Owner = s.config.GitHub.DefaultOrg
		}
		if request.GitHubRepo.Repository == "" {
			request.GitHubRepo.Repository = s.config.GitHub.DefaultRepo
		}
		test.GitHubIssueRequest = request.GitHubRepo
	}

	// Queue test for processing
	if err := s.manager.GetValidationService().QueueValidationTest(c.Request.Context(), test); err != nil {
		s.logger.WithError(err).Error("Failed to queue validation test")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to queue validation test"})
		return
	}

	// Return response
	c.JSON(http.StatusAccepted, models.ValidationResponse{
		TestID:    test.ID,
		Status:    test.Status,
		CreatedAt: test.CreatedAt,
	})
}

// listValidationsHandler lists all validation tests
func (s *Server) listValidationsHandler(c *gin.Context) {
	tests, err := s.manager.GetValidationService().ListValidationTests(c.Request.Context())
	if err != nil {
		s.logger.WithError(err).Error("Failed to list validation tests")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to list validation tests"})
		return
	}

	// Convert to response objects
	var responses []models.ValidationResponse
	for _, test := range tests {
		responses = append(responses, models.ValidationResponse{
			TestID:    test.ID,
			Status:    test.Status,
			CreatedAt: test.CreatedAt,
			Result:    test.Result,
		})
	}

	c.JSON(http.StatusOK, responses)
}

// getValidationHandler gets a validation test by ID
func (s *Server) getValidationHandler(c *gin.Context) {
	testID := c.Param("id")
	if testID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Test ID is required"})
		return
	}

	test, err := s.manager.GetValidationService().GetValidationTest(c.Request.Context(), testID)
	if err != nil {
		s.logger.WithError(err).Error("Failed to get validation test")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get validation test"})
		return
	}

	if test == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Validation test not found"})
		return
	}

	c.JSON(http.StatusOK, models.ValidationResponse{
		TestID:    test.ID,
		Status:    test.Status,
		CreatedAt: test.CreatedAt,
		Result:    test.Result,
	})
}

// getValidationReportHandler gets a validation test report
func (s *Server) getValidationReportHandler(c *gin.Context) {
	testID := c.Param("id")
	if testID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Test ID is required"})
		return
	}

	// Get test
	test, err := s.manager.GetValidationService().GetValidationTest(c.Request.Context(), testID)
	if err != nil {
		s.logger.WithError(err).Error("Failed to get validation test")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get validation test"})
		return
	}

	if test == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Validation test not found"})
		return
	}

	// Check if test has completed
	if test.Status != models.TestStatusCompleted && test.Status != models.TestStatusFailed {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Test has not completed yet"})
		return
	}

	// Check if test has report
	if test.Result == nil || test.Result.ReportPath == "" {
		// Generate report if not already generated
		reportPath, err := s.manager.GetValidationService().GenerateReport(c.Request.Context(), testID)
		if err != nil {
			s.logger.WithError(err).Error("Failed to generate validation report")
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate validation report"})
			return
		}

		// Update test result with report path
		if test.Result != nil {
			test.Result.ReportPath = reportPath
		}
	}

	// Get report content
	reportContent, err := s.manager.GetStorageService().GetReport(c.Request.Context(), test.Result.ReportPath)
	if err != nil {
		s.logger.WithError(err).Error("Failed to get validation report")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get validation report"})
		return
	}

	// Return report
	c.Header("Content-Type", "text/markdown")
	c.String(http.StatusOK, reportContent)
}

// createLoadTestHandler creates a new load test
func (s *Server) createLoadTestHandler(c *gin.Context) {
	var request models.LoadTestRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Set defaults if not provided
	if request.Timeout == nil {
		timeout := int(s.config.Testing.LoadTest.DefaultTimeout.Seconds())
		request.Timeout = &timeout
	}
	if request.SuccessThreshold == nil {
		threshold := 95.0
		request.SuccessThreshold = &threshold
	}

	// Create test
	duration := time.Duration(request.Duration) * time.Second
	test := models.NewLoadTest(
		request.Name,
		request.Description,
		request.TargetURL,
		request.RequestRate,
		duration,
	)
	test.Timeout = time.Duration(*request.Timeout) * time.Second
	test.SuccessThreshold = *request.SuccessThreshold
	test.Tags = request.Tags
	test.Method = request.Method
	test.Headers = request.Headers
	test.Body = request.Body

	// Add endpoints if provided
	if len(request.Endpoints) > 0 {
		test.Endpoints = request.Endpoints
	}

	// Setup GitHub issue creation if requested
	if request.CreateGitHubIssue && request.GitHubRepo != nil {
		if request.GitHubRepo.Owner == "" {
			request.GitHubRepo.Owner = s.config.GitHub.DefaultOrg
		}
		if request.GitHubRepo.Repository == "" {
			request.GitHubRepo.Repository = s.config.GitHub.DefaultRepo
		}
		test.GitHubIssueRequest = request.GitHubRepo
	}

	// Queue test for processing
	if err := s.manager.GetLoadTestService().QueueLoadTest(c.Request.Context(), test); err != nil {
		s.logger.WithError(err).Error("Failed to queue load test")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to queue load test"})
		return
	}

	// Return response
	c.JSON(http.StatusAccepted, models.LoadTestResponse{
		TestID:    test.ID,
		Status:    test.Status,
		CreatedAt: test.CreatedAt,
	})
}

// listLoadTestsHandler lists all load tests
func (s *Server) listLoadTestsHandler(c *gin.Context) {
	tests, err := s.manager.GetLoadTestService().ListLoadTests(c.Request.Context())
	if err != nil {
		s.logger.WithError(err).Error("Failed to list load tests")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to list load tests"})
		return
	}

	// Convert to response objects
	var responses []models.LoadTestResponse
	for _, test := range tests {
		responses = append(responses, models.LoadTestResponse{
			TestID:    test.ID,
			Status:    test.Status,
			CreatedAt: test.CreatedAt,
			Result:    test.Result,
		})
	}

	c.JSON(http.StatusOK, responses)
}

// getLoadTestHandler gets a load test by ID
func (s *Server) getLoadTestHandler(c *gin.Context) {
	testID := c.Param("id")
	if testID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Test ID is required"})
		return
	}

	test, err := s.manager.GetLoadTestService().GetLoadTest(c.Request.Context(), testID)
	if err != nil {
		s.logger.WithError(err).Error("Failed to get load test")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get load test"})
		return
	}

	if test == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Load test not found"})
		return
	}

	c.JSON(http.StatusOK, models.LoadTestResponse{
		TestID:    test.ID,
		Status:    test.Status,
		CreatedAt: test.CreatedAt,
		Result:    test.Result,
	})
}

// getLoadTestReportHandler gets a load test report
func (s *Server) getLoadTestReportHandler(c *gin.Context) {
	testID := c.Param("id")
	if testID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Test ID is required"})
		return
	}

	// Get test
	test, err := s.manager.GetLoadTestService().GetLoadTest(c.Request.Context(), testID)
	if err != nil {
		s.logger.WithError(err).Error("Failed to get load test")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get load test"})
		return
	}

	if test == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Load test not found"})
		return
	}

	// Check if test has completed
	if test.Status != models.TestStatusCompleted && test.Status != models.TestStatusFailed {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Test has not completed yet"})
		return
	}

	// Check if test has report
	if test.Result == nil || test.Result.ReportPath == "" {
		// Generate report if not already generated
		reportPath, err := s.manager.GetLoadTestService().GenerateReport(c.Request.Context(), testID)
		if err != nil {
			s.logger.WithError(err).Error("Failed to generate load test report")
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate load test report"})
			return
		}

		// Update test result with report path
		if test.Result != nil {
			test.Result.ReportPath = reportPath
		}
	}

	// Get report content
	reportContent, err := s.manager.GetStorageService().GetReport(c.Request.Context(), test.Result.ReportPath)
	if err != nil {
		s.logger.WithError(err).Error("Failed to get load test report")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get load test report"})
		return
	}

	// Return report
	c.Header("Content-Type", "text/markdown")
	c.String(http.StatusOK, reportContent)
}

// createAcceptanceTestHandler creates a new acceptance test
func (s *Server) createAcceptanceTestHandler(c *gin.Context) {
	var request models.AcceptanceTestRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Set defaults if not provided
	if request.Timeout == nil {
		timeout := int(s.config.Testing.Acceptance.DefaultTimeout.Seconds())
		request.Timeout = &timeout
	}

	// Create test
	test := models.NewAcceptanceTest(
		request.Name,
		request.Description,
		request.BaseURL,
	)
	test.Timeout = time.Duration(*request.Timeout) * time.Second
	test.Tags = request.Tags
	test.Headers = request.Headers
	test.TestCases = request.TestCases
	
	if request.GlobalVariables != nil {
		test.GlobalVariables = request.GlobalVariables
	}

	// Setup GitHub issue creation if requested
	if request.CreateGitHubIssue && request.GitHubRepo != nil {
		if request.GitHubRepo.Owner == "" {
			request.GitHubRepo.Owner = s.config.GitHub.DefaultOrg
		}
		if request.GitHubRepo.Repository == "" {
			request.GitHubRepo.Repository = s.config.GitHub.DefaultRepo
		}
		test.GitHubIssueRequest = request.GitHubRepo
	}

	// Queue test for processing
	if err := s.manager.GetAcceptanceService().QueueAcceptanceTest(c.Request.Context(), test); err != nil {
		s.logger.WithError(err).Error("Failed to queue acceptance test")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to queue acceptance test"})
		return
	}

	// Return response
	c.JSON(http.StatusAccepted, models.AcceptanceTestResponse{
		TestID:    test.ID,
		Status:    test.Status,
		CreatedAt: test.CreatedAt,
	})
}

// listAcceptanceTestsHandler lists all acceptance tests
func (s *Server) listAcceptanceTestsHandler(c *gin.Context) {
	tests, err := s.manager.GetAcceptanceService().ListAcceptanceTests(c.Request.Context())
	if err != nil {
		s.logger.WithError(err).Error("Failed to list acceptance tests")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to list acceptance tests"})
		return
	}

	// Convert to response objects
	var responses []models.AcceptanceTestResponse
	for _, test := range tests {
		responses = append(responses, models.AcceptanceTestResponse{
			TestID:    test.ID,
			Status:    test.Status,
			CreatedAt: test.CreatedAt,
			Result:    test.Result,
		})
	}

	c.JSON(http.StatusOK, responses)
}

// getAcceptanceTestHandler gets an acceptance test by ID
func (s *Server) getAcceptanceTestHandler(c *gin.Context) {
	testID := c.Param("id")
	if testID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Test ID is required"})
		return
	}

	test, err := s.manager.GetAcceptanceService().GetAcceptanceTest(c.Request.Context(), testID)
	if err != nil {
		s.logger.WithError(err).Error("Failed to get acceptance test")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get acceptance test"})
		return
	}

	if test == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Acceptance test not found"})
		return
	}

	c.JSON(http.StatusOK, models.AcceptanceTestResponse{
		TestID:    test.ID,
		Status:    test.Status,
		CreatedAt: test.CreatedAt,
		Result:    test.Result,
	})
}

// getAcceptanceTestReportHandler gets an acceptance test report
func (s *Server) getAcceptanceTestReportHandler(c *gin.Context) {
	testID := c.Param("id")
	if testID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Test ID is required"})
		return
	}

	// Get test
	test, err := s.manager.GetAcceptanceService().GetAcceptanceTest(c.Request.Context(), testID)
	if err != nil {
		s.logger.WithError(err).Error("Failed to get acceptance test")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get acceptance test"})
		return
	}

	if test == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Acceptance test not found"})
		return
	}

	// Check if test has completed
	if test.Status != models.TestStatusCompleted && test.Status != models.TestStatusFailed {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Test has not completed yet"})
		return
	}

	// Check if test has report
	if test.Result == nil || test.Result.ReportPath == "" {
		// Generate report if not already generated
		reportPath, err := s.manager.GetAcceptanceService().GenerateReport(c.Request.Context(), testID)
		if err != nil {
			s.logger.WithError(err).Error("Failed to generate acceptance test report")
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate acceptance test report"})
			return
		}

		// Update test result with report path
		if test.Result != nil {
			test.Result.ReportPath = reportPath
		}
	}

	// Get report content
	reportContent, err := s.manager.GetStorageService().GetReport(c.Request.Context(), test.Result.ReportPath)
	if err != nil {
		s.logger.WithError(err).Error("Failed to get acceptance test report")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get acceptance test report"})
		return
	}

	// Return report
	c.Header("Content-Type", "text/markdown")
	c.String(http.StatusOK, reportContent)
}