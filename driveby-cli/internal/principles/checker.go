package principles

import (
	"context"

	"github.com/meter-peter/driveby/driveby-cli/internal/spec"
	"github.com/meter-peter/driveby/driveby-cli/internal/types"
)

// PrincipleChecker validates a single DDT principle against an API spec.
type PrincipleChecker interface {
	ID() string
	Check(ctx context.Context, apiSpec spec.APISpec, mode types.ValidationMode) types.PrincipleResult
}
