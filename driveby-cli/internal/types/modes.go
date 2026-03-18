package types

// ValidationMode defines the level of validation to perform
type ValidationMode string

const (
	ValidationModeStrict   ValidationMode = "strict"    // Comprehensive validation
	ValidationModeMinimal  ValidationMode = "minimal"   // Basic validation for test generation
	ValidationModeTestOnly ValidationMode = "test-only" // Skip validation, run tests only
	ValidationModeFlexible ValidationMode = "flexible"  // Allow tests even with some validation failures
)

// TestMode defines the type of testing to perform
type TestMode string

const (
	TestModeNone        TestMode = "none"        // No testing
	TestModeFunctional  TestMode = "functional"  // Only functional testing
	TestModePerformance TestMode = "performance" // Only performance testing
	TestModeAll         TestMode = "all"         // Both functional and performance testing
)

// TestStatus represents the overall status of tests
type TestStatus string

const (
	TestStatusPassed     TestStatus = "passed"
	TestStatusFailed     TestStatus = "failed"
	TestStatusWarning    TestStatus = "warning"
	TestStatusSkipped    TestStatus = "skipped"
	TestStatusIncomplete TestStatus = "incomplete"
)

// ImpactLevel represents the severity of test impact
type ImpactLevel string

const (
	ImpactLevelNone     ImpactLevel = "none"
	ImpactLevelLow      ImpactLevel = "low"
	ImpactLevelMedium   ImpactLevel = "medium"
	ImpactLevelHigh     ImpactLevel = "high"
	ImpactLevelCritical ImpactLevel = "critical"
)

// Exit codes for the CLI
const (
	ExitSuccess          = 0 // Successful execution
	ExitValidationFailed = 1 // Tests ran but failed validation
	ExitExecutionError   = 2 // Error executing tests
	ExitInvalidArgs      = 3 // Invalid command line arguments
)
