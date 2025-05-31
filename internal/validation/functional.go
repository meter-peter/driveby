package validation

import (
	"context"
	"fmt"
	"time"

	"github.com/meter-peter/driveby/internal/openapi"
)

// RunFunctional runs endpoint functional tests (P006)
func RunFunctional(ctx context.Context, config ValidatorConfig) (*ValidationReport, error) {
	loader := openapi.NewLoader()
	if err := loader.LoadFromFileOrURL(config.SpecPath); err != nil {
		return nil, fmt.Errorf("failed to load OpenAPI spec: %w", err)
	}
	doc := loader.GetDocument()
	if doc == nil {
		return nil, fmt.Errorf("failed to get OpenAPI document")
	}

	validator := NewValidator()
	validator.SetBaseURL(config.BaseURL)
	endpointResult, err := validator.ValidateEndpoints(ctx, doc, config.BaseURL)
	if err != nil {
		return nil, fmt.Errorf("endpoint functional testing failed: %w", err)
	}

	allSuccess := true
	var failedEndpoints []string
	for _, epVal := range endpointResult.Endpoints {
		if epVal.Status != "success" {
			allSuccess = false
			failedEndpoints = append(failedEndpoints, fmt.Sprintf("%s %s (Status: %s, Code: %d)", epVal.Method, epVal.Path, epVal.Status, epVal.StatusCode))
		}
	}

	principleResult := PrincipleResult{
		Principle: Principle{
			ID:   "P006",
			Name: "Endpoint Functional Testing",
		},
		Passed:  allSuccess,
		Details: endpointResult.Endpoints,
	}
	if allSuccess {
		principleResult.Message = "All documented endpoints are reachable and return documented status codes."
	} else {
		principleResult.Message = fmt.Sprintf("Some endpoints failed functional tests. Failed: %d/%d", len(failedEndpoints), len(endpointResult.Endpoints))
		principleResult.Details = map[string]interface{}{"failed_endpoints": failedEndpoints, "all_results": endpointResult.Endpoints}
	}

	report := &ValidationReport{
		Version:     config.Version,
		Environment: config.Environment,
		Timestamp:   time.Now(),
		Principles:  []PrincipleResult{principleResult},
	}
	return report, nil
}
