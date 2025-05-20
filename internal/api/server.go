package api

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"driveby/internal/config"
	"driveby/internal/core"
	"driveby/internal/core/models"
	"driveby/internal/core/services"
	"driveby/internal/types"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

// @title           Driveby Testing API
// @version         1.0.0
// @description     Documentation-driven API testing service. Validates OpenAPI documentation, runs integration and load tests, and enforces quality gates.
// @termsOfService  http://example.com/terms/

// @contact.name   API Support
// @contact.url    http://www.example.com/support
// @contact.email  support@example.com

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8081
// @BasePath  /api/v1

// Server represents the API server
type Server struct {
	router      *mux.Router
	logger      *logrus.Logger
	testingSvc  *core.TestingService
	apiHost     string
	apiPort     string
	apiBasePath string
	config      *config.Config
	manager     *services.ServiceManager
}

// NewServer creates a new API server
func NewServer(logger *logrus.Logger, testingSvc *core.TestingService, apiHost, apiPort, apiBasePath string, cfg *config.Config, manager *services.ServiceManager) *Server {
	s := &Server{
		router:      mux.NewRouter(),
		logger:      logger,
		testingSvc:  testingSvc,
		apiHost:     apiHost,
		apiPort:     apiPort,
		apiBasePath: apiBasePath,
		config:      cfg,
		manager:     manager,
	}

	s.setupRoutes()
	return s
}

// setupRoutes configures the API routes
func (s *Server) setupRoutes() {
	// API routes
	apiRouter := s.router.PathPrefix(s.apiBasePath).Subrouter()

	// Health check endpoint under API base path
	apiRouter.HandleFunc("/health", s.handleHealthCheck).Methods(http.MethodGet)

	// OpenAPI documentation endpoints
	apiRouter.HandleFunc("/docs", s.handleSwaggerUI).Methods(http.MethodGet)
	apiRouter.HandleFunc("/openapi.json", s.handleOpenAPISpec).Methods(http.MethodGet)

	// Testing endpoints
	apiRouter.HandleFunc("/tests", s.handleRunTests).Methods(http.MethodPost)
	apiRouter.HandleFunc("/tests/{test_id}", s.handleGetTestResult).Methods(http.MethodGet)

	// Validation routes
	apiRouter.HandleFunc("/validation", s.createValidationHandler).Methods(http.MethodPost)
	apiRouter.HandleFunc("/validation", s.listValidationsHandler).Methods(http.MethodGet)
	apiRouter.HandleFunc("/validation/{id}", s.getValidationHandler).Methods(http.MethodGet)
	apiRouter.HandleFunc("/validation/{id}/report", s.getValidationReportHandler).Methods(http.MethodGet)

	// Load test routes
	apiRouter.HandleFunc("/loadtest", s.createLoadTestHandler).Methods(http.MethodPost)
	apiRouter.HandleFunc("/loadtest", s.listLoadTestsHandler).Methods(http.MethodGet)
	apiRouter.HandleFunc("/loadtest/{id}", s.getLoadTestHandler).Methods(http.MethodGet)
	apiRouter.HandleFunc("/loadtest/{id}/report", s.getLoadTestReportHandler).Methods(http.MethodGet)

	// Acceptance test routes
	apiRouter.HandleFunc("/acceptance", s.createAcceptanceTestHandler).Methods(http.MethodPost)
	apiRouter.HandleFunc("/acceptance", s.listAcceptanceTestsHandler).Methods(http.MethodGet)
	apiRouter.HandleFunc("/acceptance/{id}", s.getAcceptanceTestHandler).Methods(http.MethodGet)
	apiRouter.HandleFunc("/acceptance/{id}/report", s.getAcceptanceTestReportHandler).Methods(http.MethodGet)
}

// @Summary     Health check endpoint
// @Description Returns the health status of the API
// @Tags        health
// @Produce     json
// @Success     200 {object} HealthResponse
// @Router      /health [get]
func (s *Server) handleHealthCheck(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"status":    "healthy",
		"version":   "1.0.0",
		"timestamp": time.Now().Format(time.RFC3339),
	})
}

// @Summary     Run full test suite
// @Description Runs documentation, integration, and load tests against a target API using its OpenAPI spec
// @Tags        tests
// @Accept      json
// @Produce     json
// @Param       request body TestRequest true "Test configuration"
// @Success     200 {object} TestResponse
// @Failure     400 {object} ErrorResponse
// @Failure     500 {object} ErrorResponse
// @Router      /tests [post]
func (s *Server) handleRunTests(w http.ResponseWriter, r *http.Request) {
	var req types.TestRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		s.logger.WithError(err).Error("Failed to decode test request")
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Run the tests
	result, err := s.testingSvc.RunTests(r.Context(), req)
	if err != nil {
		s.logger.WithError(err).Error("Failed to run tests")
		http.Error(w, "Failed to run tests", http.StatusInternalServerError)
		return
	}

	// Return the results
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}

// handleGetTestResult handles retrieving test results
func (s *Server) handleGetTestResult(w http.ResponseWriter, r *http.Request) {
	_ = mux.Vars(r)["test_id"] // Ignore test_id for now
	http.Error(w, "Test result retrieval not implemented", http.StatusNotImplemented)
}

// Start starts the API server
func (s *Server) Start() error {
	addr := s.apiHost + ":" + s.apiPort
	s.logger.Infof("Starting API server on %s", addr)
	return http.ListenAndServe(addr, s.router)
}

// Shutdown gracefully shuts down the server
func (s *Server) Shutdown(ctx context.Context) error {
	// TODO: Implement graceful shutdown
	return nil
}

// @Summary     Create and run validation tests
// @Description Validates API implementation against OpenAPI spec
// @Tags        validation
// @Accept      json
// @Produce     json
// @Param       request body ValidationRequest true "Validation test configuration"
// @Success     200 {object} ValidationResult
// @Failure     400 {object} ErrorResponse
// @Failure     500 {object} ErrorResponse
// @Router      /validation [post]
func (s *Server) createValidationHandler(w http.ResponseWriter, r *http.Request) {
	var req struct {
		OpenAPISpec string `json:"openapi_spec"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		s.logger.WithError(err).Error("Failed to decode validation test request")
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if req.OpenAPISpec == "" {
		s.logger.Error("OpenAPI spec URL is required")
		http.Error(w, "OpenAPI spec URL is required", http.StatusBadRequest)
		return
	}

	// Create validation test
	test := models.NewValidationTest(
		"OpenAPI Validation",
		"Validating API implementation against OpenAPI specification",
		req.OpenAPISpec,
		95.0, // Default compliance threshold
	)

	// Run validation
	result, err := s.manager.GetValidationService().ValidateOpenAPI(r.Context(), test)
	if err != nil {
		s.logger.WithError(err).Error("Failed to validate OpenAPI spec")
		// Return a more detailed error response
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"error":   "Failed to validate OpenAPI spec",
			"details": err.Error(),
		})
		return
	}

	// Return the results
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}

// listValidationsHandler lists all validation tests
func (s *Server) listValidationsHandler(w http.ResponseWriter, r *http.Request) {
	tests, err := s.manager.GetValidationService().ListValidationTests(r.Context())
	if err != nil {
		s.logger.WithError(err).Error("Failed to list validation tests")
		http.Error(w, "Failed to list validation tests", http.StatusInternalServerError)
		return
	}

	// Return the tests
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tests)
}

// getValidationHandler gets a validation test by ID
func (s *Server) getValidationHandler(w http.ResponseWriter, r *http.Request) {
	testID := mux.Vars(r)["id"]
	if testID == "" {
		http.Error(w, "Test ID is required", http.StatusBadRequest)
		return
	}

	test, err := s.manager.GetValidationService().GetValidationTest(r.Context(), testID)
	if err != nil {
		s.logger.WithError(err).Error("Failed to get validation test")
		http.Error(w, "Failed to get validation test", http.StatusInternalServerError)
		return
	}

	if test == nil {
		http.Error(w, "Validation test not found", http.StatusNotFound)
		return
	}

	// Return the test
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(test)
}

// getValidationReportHandler gets a validation test report
func (s *Server) getValidationReportHandler(w http.ResponseWriter, r *http.Request) {
	testID := mux.Vars(r)["id"]
	if testID == "" {
		http.Error(w, "Test ID is required", http.StatusBadRequest)
		return
	}

	report, err := s.manager.GetValidationService().GenerateReport(r.Context(), testID)
	if err != nil {
		s.logger.WithError(err).Error("Failed to generate validation report")
		http.Error(w, "Failed to generate validation report", http.StatusInternalServerError)
		return
	}

	// Return the report
	w.Header().Set("Content-Type", "text/markdown")
	w.Write([]byte(report))
}

// @Summary     Create and run load tests
// @Description Performs load testing with configurable parameters
// @Tags        loadtest
// @Accept      json
// @Produce     json
// @Param       request body LoadTest true "Load test configuration"
// @Success     200 {object} LoadTestResult
// @Failure     400 {object} ErrorResponse
// @Failure     500 {object} ErrorResponse
// @Router      /loadtest [post]
func (s *Server) createLoadTestHandler(w http.ResponseWriter, r *http.Request) {
	var test models.LoadTest
	if err := json.NewDecoder(r.Body).Decode(&test); err != nil {
		s.logger.WithError(err).Error("Failed to decode load test request")
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Set defaults if not provided
	if test.Timeout == 0 {
		test.Timeout = s.config.Testing.LoadTest.DefaultTimeout
	}
	if test.RequestRate == 0 {
		test.RequestRate = s.config.Testing.LoadTest.DefaultRPS
	}
	if test.Duration == 0 {
		test.Duration = s.config.Testing.LoadTest.DefaultDuration
	}

	// Run load test
	result, err := s.manager.GetLoadTestService().RunLoadTest(r.Context(), &test)
	if err != nil {
		s.logger.WithError(err).Error("Failed to run load test")
		http.Error(w, "Failed to run load test", http.StatusInternalServerError)
		return
	}

	// Return the results
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}

// listLoadTestsHandler lists all load tests
func (s *Server) listLoadTestsHandler(w http.ResponseWriter, r *http.Request) {
	tests, err := s.manager.GetLoadTestService().ListLoadTests(r.Context())
	if err != nil {
		s.logger.WithError(err).Error("Failed to list load tests")
		http.Error(w, "Failed to list load tests", http.StatusInternalServerError)
		return
	}

	// Return the tests
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tests)
}

// getLoadTestHandler gets a load test by ID
func (s *Server) getLoadTestHandler(w http.ResponseWriter, r *http.Request) {
	testID := mux.Vars(r)["id"]
	if testID == "" {
		http.Error(w, "Test ID is required", http.StatusBadRequest)
		return
	}

	test, err := s.manager.GetLoadTestService().GetLoadTest(r.Context(), testID)
	if err != nil {
		s.logger.WithError(err).Error("Failed to get load test")
		http.Error(w, "Failed to get load test", http.StatusInternalServerError)
		return
	}

	if test == nil {
		http.Error(w, "Load test not found", http.StatusNotFound)
		return
	}

	// Return the test
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(test)
}

// getLoadTestReportHandler gets a load test report
func (s *Server) getLoadTestReportHandler(w http.ResponseWriter, r *http.Request) {
	testID := mux.Vars(r)["id"]
	if testID == "" {
		http.Error(w, "Test ID is required", http.StatusBadRequest)
		return
	}

	report, err := s.manager.GetLoadTestService().GenerateReport(r.Context(), testID)
	if err != nil {
		s.logger.WithError(err).Error("Failed to generate load test report")
		http.Error(w, "Failed to generate load test report", http.StatusInternalServerError)
		return
	}

	// Return the report
	w.Header().Set("Content-Type", "text/markdown")
	w.Write([]byte(report))
}

// @Summary     Create and run acceptance tests
// @Description Validates business requirements
// @Tags        acceptance
// @Accept      json
// @Produce     json
// @Param       request body AcceptanceTest true "Acceptance test configuration"
// @Success     200 {object} AcceptanceTestResult
// @Failure     400 {object} ErrorResponse
// @Failure     500 {object} ErrorResponse
// @Router      /acceptance [post]
func (s *Server) createAcceptanceTestHandler(w http.ResponseWriter, r *http.Request) {
	var test models.AcceptanceTest
	if err := json.NewDecoder(r.Body).Decode(&test); err != nil {
		s.logger.WithError(err).Error("Failed to decode acceptance test request")
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Set defaults if not provided
	if test.Timeout == 0 {
		test.Timeout = s.config.Testing.Acceptance.DefaultTimeout
	}

	// Run acceptance test
	result, err := s.manager.GetAcceptanceService().RunAcceptanceTest(r.Context(), &test)
	if err != nil {
		s.logger.WithError(err).Error("Failed to run acceptance test")
		http.Error(w, "Failed to run acceptance test", http.StatusInternalServerError)
		return
	}

	// Return the results
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}

// listAcceptanceTestsHandler lists all acceptance tests
func (s *Server) listAcceptanceTestsHandler(w http.ResponseWriter, r *http.Request) {
	tests, err := s.manager.GetAcceptanceService().ListAcceptanceTests(r.Context())
	if err != nil {
		s.logger.WithError(err).Error("Failed to list acceptance tests")
		http.Error(w, "Failed to list acceptance tests", http.StatusInternalServerError)
		return
	}

	// Return the tests
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tests)
}

// getAcceptanceTestHandler gets an acceptance test by ID
func (s *Server) getAcceptanceTestHandler(w http.ResponseWriter, r *http.Request) {
	testID := mux.Vars(r)["id"]
	if testID == "" {
		http.Error(w, "Test ID is required", http.StatusBadRequest)
		return
	}

	test, err := s.manager.GetAcceptanceService().GetAcceptanceTest(r.Context(), testID)
	if err != nil {
		s.logger.WithError(err).Error("Failed to get acceptance test")
		http.Error(w, "Failed to get acceptance test", http.StatusInternalServerError)
		return
	}

	if test == nil {
		http.Error(w, "Acceptance test not found", http.StatusNotFound)
		return
	}

	// Return the test
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(test)
}

// getAcceptanceTestReportHandler gets an acceptance test report
func (s *Server) getAcceptanceTestReportHandler(w http.ResponseWriter, r *http.Request) {
	testID := mux.Vars(r)["id"]
	if testID == "" {
		http.Error(w, "Test ID is required", http.StatusBadRequest)
		return
	}

	report, err := s.manager.GetAcceptanceService().GenerateReport(r.Context(), testID)
	if err != nil {
		s.logger.WithError(err).Error("Failed to generate acceptance test report")
		http.Error(w, "Failed to generate acceptance test report", http.StatusInternalServerError)
		return
	}

	// Return the report
	w.Header().Set("Content-Type", "text/markdown")
	w.Write([]byte(report))
}

// handleSwaggerUI serves the Swagger UI
func (s *Server) handleSwaggerUI(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.Write([]byte(swaggerUIHTML))
}

// handleOpenAPISpec serves the OpenAPI specification
func (s *Server) handleOpenAPISpec(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(openAPISpec))
}

// Swagger UI HTML template
const swaggerUIHTML = `<!DOCTYPE html>
<html>
<head>
    <title>Driveby API Documentation</title>
    <meta charset="utf-8"/>
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <link rel="stylesheet" type="text/css" href="https://unpkg.com/swagger-ui-dist@4/swagger-ui.css" />
</head>
<body>
    <div id="swagger-ui"></div>
    <script src="https://unpkg.com/swagger-ui-dist@4/swagger-ui-bundle.js"></script>
    <script>
        window.onload = function() {
            SwaggerUIBundle({
                url: "openapi.json",
                dom_id: '#swagger-ui',
                deepLinking: true,
                presets: [
                    SwaggerUIBundle.presets.apis,
                    SwaggerUIBundle.SwaggerUIStandalonePreset
                ],
            });
        }
    </script>
</body>
</html>`

// OpenAPI specification
const openAPISpec = `{
    "openapi": "3.0.0",
    "info": {
        "title": "Driveby Testing API",
        "version": "1.0.0",
        "description": "Documentation-driven API testing service. Validates OpenAPI documentation, runs integration and load tests, and enforces quality gates."
    },
    "servers": [
        {
            "url": "http://localhost:8081/api/v1",
            "description": "Local server"
        }
    ],
    "paths": {
        "/health": {
            "get": {
                "summary": "Health Check",
                "description": "Returns the health status of the API.",
                "responses": {
                    "200": {
                        "description": "API is healthy",
                        "content": {
                            "application/json": {
                                "schema": {
                                    "type": "object",
                                    "properties": {
                                        "status": { "type": "string" },
                                        "version": { "type": "string" },
                                        "timestamp": { "type": "string", "format": "date-time" }
                                    }
                                },
                                "example": {
                                    "status": "healthy",
                                    "version": "1.0.0",
                                    "timestamp": "2024-03-20T10:00:00Z"
                                }
                            }
                        }
                    }
                }
            }
        },
        "/tests": {
            "post": {
                "summary": "Run Full Test Suite",
                "description": "Runs documentation, integration, and load tests against a target API using its OpenAPI spec.",
                "requestBody": {
                    "required": true,
                    "content": {
                        "application/json": {
                            "schema": { "$ref": "#/components/schemas/TestRequest" },
                            "example": {
                                "openapi_spec": "http://perfect-api:8080/openapi.json",
                                "thresholds": {
                                    "documentation": {
                                        "threshold": 0.95
                                    },
                                    "load_test": {
                                        "success_rate": 0.99,
                                        "max_latency": "500ms"
                                    }
                                }
                            }
                        }
                    }
                },
                "responses": {
                    "200": {
                        "description": "Test results",
                        "content": {
                            "application/json": {
                                "schema": { "$ref": "#/components/schemas/TestResponse" },
                                "example": {
                                    "test_id": "test_123456",
                                    "timestamp": "2024-03-20T10:00:00Z",
                                    "results": {
                                        "documentation": {
                                            "compliance_score": 0.98,
                                            "missing_examples": 1,
                                            "undocumented_endpoints": ["/api/v1/health"],
                                            "errors": []
                                        },
                                        "integration": {
                                            "total_tests": 25,
                                            "passed": 24,
                                            "failed": 1,
                                            "failed_endpoints": [
                                                {
                                                    "endpoint": "/api/v1/users",
                                                    "error": "Response schema mismatch"
                                                }
                                            ]
                                        },
                                        "load_test": {
                                            "total_requests": 1000,
                                            "success_rate": 0.995,
                                            "latency_p95": "450ms",
                                            "status_codes": {
                                                "200": 995,
                                                "500": 5
                                            },
                                            "errors": []
                                        }
                                    }
                                }
                            }
                        }
                    }
                }
            }
        },
        "/validation": {
            "post": {
                "summary": "Validate OpenAPI Documentation",
                "description": "Checks endpoint documentation completeness, response descriptions, examples, and error documentation.",
                "requestBody": {
                    "required": true,
                    "content": {
                        "application/json": {
                            "schema": { "$ref": "#/components/schemas/ValidationRequest" },
                            "example": {
                                "openapi_spec": "http://perfect-api:8080/openapi.json"
                            }
                        }
                    }
                },
                "responses": {
                    "200": {
                        "description": "Validation results",
                        "content": {
                            "application/json": {
                                "schema": { "$ref": "#/components/schemas/ValidationResult" },
                                "example": {
                                    "compliance_score": 0.98,
                                    "missing_examples": 1,
                                    "undocumented_endpoints": ["/api/v1/health"],
                                    "errors": [
                                        "Missing response example for POST /api/v1/users"
                                    ]
                                }
                            }
                        }
                    }
                }
            }
        },
        "/loadtest": {
            "post": {
                "summary": "Run Load Test",
                "description": "Runs a load test using Vegeta against the endpoints in the OpenAPI spec.",
                "requestBody": {
                    "required": true,
                    "content": {
                        "application/json": {
                            "schema": { "$ref": "#/components/schemas/LoadTestRequest" },
                            "example": {
                                "openapi_spec": "http://perfect-api:8080/openapi.json",
                                "request_rate": 100,
                                "duration": "30s"
                            }
                        }
                    }
                },
                "responses": {
                    "200": {
                        "description": "Load test results",
                        "content": {
                            "application/json": {
                                "schema": { "$ref": "#/components/schemas/LoadTestResult" },
                                "example": {
                                    "total_requests": 3000,
                                    "success_rate": 0.995,
                                    "latency_p95": "450ms",
                                    "status_codes": {
                                        "200": 2985,
                                        "500": 15
                                    },
                                    "errors": [
                                        "Connection timeout on POST /api/v1/users"
                                    ]
                                }
                            }
                        }
                    }
                }
            }
        },
        "/acceptance": {
            "post": {
                "summary": "Run Acceptance Test",
                "description": "Runs acceptance tests to validate business requirements as described in the OpenAPI spec.",
                "requestBody": {
                    "required": true,
                    "content": {
                        "application/json": {
                            "schema": { "$ref": "#/components/schemas/AcceptanceTestRequest" },
                            "example": {
                                "openapi_spec": "http://perfect-api:8080/openapi.json"
                            }
                        }
                    }
                },
                "responses": {
                    "200": {
                        "description": "Acceptance test results",
                        "content": {
                            "application/json": {
                                "schema": { "$ref": "#/components/schemas/AcceptanceTestResult" },
                                "example": {
                                    "passed": true,
                                    "details": "All business requirements met. User creation, authentication, and data retrieval workflows validated successfully."
                                }
                            }
                        }
                    }
                }
            }
        }
    },
    "components": {
        "schemas": {
            "TestRequest": {
                "type": "object",
                "properties": {
                    "openapi_spec": { 
                        "type": "string", 
                        "description": "URL to the OpenAPI specification",
                        "example": "http://perfect-api:8080/openapi.json"
                    },
                    "thresholds": {
                        "type": "object",
                        "properties": {
                            "documentation": { 
                                "type": "object", 
                                "properties": { 
                                    "threshold": { 
                                        "type": "number",
                                        "example": 0.95
                                    } 
                                } 
                            },
                            "load_test": {
                                "type": "object",
                                "properties": {
                                    "success_rate": { 
                                        "type": "number",
                                        "example": 0.99
                                    },
                                    "max_latency": { 
                                        "type": "string",
                                        "example": "500ms"
                                    }
                                }
                            }
                        }
                    }
                }
            },
            "TestResponse": {
                "type": "object",
                "properties": {
                    "test_id": { 
                        "type": "string",
                        "example": "test_123456"
                    },
                    "timestamp": { 
                        "type": "string", 
                        "format": "date-time",
                        "example": "2024-03-20T10:00:00Z"
                    },
                    "results": {
                        "type": "object",
                        "properties": {
                            "documentation": { "$ref": "#/components/schemas/ValidationResult" },
                            "integration": { "$ref": "#/components/schemas/IntegrationResult" },
                            "load_test": { "$ref": "#/components/schemas/LoadTestResult" }
                        }
                    }
                }
            },
            "ValidationRequest": {
                "type": "object",
                "properties": {
                    "openapi_spec": { 
                        "type": "string",
                        "example": "http://perfect-api:8080/openapi.json"
                    }
                }
            },
            "ValidationResult": {
                "type": "object",
                "properties": {
                    "compliance_score": { 
                        "type": "number", 
                        "description": "Percent of endpoints fully documented",
                        "example": 0.98
                    },
                    "missing_examples": { 
                        "type": "integer",
                        "example": 1
                    },
                    "undocumented_endpoints": { 
                        "type": "array", 
                        "items": { "type": "string" },
                        "example": ["/api/v1/health"]
                    },
                    "errors": { 
                        "type": "array", 
                        "items": { "type": "string" },
                        "example": ["Missing response example for POST /api/v1/users"]
                    }
                }
            },
            "IntegrationResult": {
                "type": "object",
                "properties": {
                    "total_tests": { 
                        "type": "integer",
                        "example": 25
                    },
                    "passed": { 
                        "type": "integer",
                        "example": 24
                    },
                    "failed": { 
                        "type": "integer",
                        "example": 1
                    },
                    "failed_endpoints": {
                        "type": "array",
                        "items": {
                            "type": "object",
                            "properties": {
                                "endpoint": { 
                                    "type": "string",
                                    "example": "/api/v1/users"
                                },
                                "error": { 
                                    "type": "string",
                                    "example": "Response schema mismatch"
                                }
                            }
                        }
                    }
                }
            },
            "LoadTestRequest": {
                "type": "object",
                "properties": {
                    "openapi_spec": { 
                        "type": "string",
                        "example": "http://perfect-api:8080/openapi.json"
                    },
                    "request_rate": { 
                        "type": "integer",
                        "example": 100
                    },
                    "duration": { 
                        "type": "string",
                        "example": "30s"
                    }
                }
            },
            "LoadTestResult": {
                "type": "object",
                "properties": {
                    "total_requests": { 
                        "type": "integer",
                        "example": 3000
                    },
                    "success_rate": { 
                        "type": "number",
                        "example": 0.995
                    },
                    "latency_p95": { 
                        "type": "string",
                        "example": "450ms"
                    },
                    "status_codes": {
                        "type": "object",
                        "additionalProperties": { "type": "integer" },
                        "example": {
                            "200": 2985,
                            "500": 15
                        }
                    },
                    "errors": { 
                        "type": "array", 
                        "items": { "type": "string" },
                        "example": ["Connection timeout on POST /api/v1/users"]
                    }
                }
            },
            "AcceptanceTestRequest": {
                "type": "object",
                "properties": {
                    "openapi_spec": { 
                        "type": "string",
                        "example": "http://perfect-api:8080/openapi.json"
                    }
                }
            },
            "AcceptanceTestResult": {
                "type": "object",
                "properties": {
                    "passed": { 
                        "type": "boolean",
                        "example": true
                    },
                    "details": { 
                        "type": "string",
                        "example": "All business requirements met. User creation, authentication, and data retrieval workflows validated successfully."
                    }
                }
            }
        }
    }
}`
