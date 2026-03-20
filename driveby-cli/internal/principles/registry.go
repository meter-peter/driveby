package principles

import "github.com/meter-peter/driveby/driveby-cli/internal/types"

// Registry holds all registered principle checkers.
type Registry struct {
	checkers []PrincipleChecker
}

// NewRegistry creates a registry with all built-in principle checkers.
func NewRegistry() *Registry {
	return &Registry{
		checkers: []PrincipleChecker{
			&P001Compliance{},
			&P002Documentation{},
			&P003Errors{},
			&P004Schema{},
			&P005Security{},
			&P008Versioning{},
			&P009TestReadiness{},
		},
	}
}

// All returns all registered checkers.
func (r *Registry) All() []PrincipleChecker {
	return r.checkers
}

// Get returns a checker by principle ID, or nil if not found.
func (r *Registry) Get(id string) PrincipleChecker {
	for _, c := range r.checkers {
		if c.ID() == id {
			return c
		}
	}
	return nil
}

// ForMode returns checkers appropriate for a given validation mode.
func (r *Registry) ForMode(mode types.ValidationMode) []PrincipleChecker {
	switch mode {
	case types.ValidationModeTestOnly:
		return nil
	case types.ValidationModeMinimal:
		return []PrincipleChecker{r.Get("P001")}
	case types.ValidationModeTestReady:
		return []PrincipleChecker{r.Get("P001"), r.Get("P004"), r.Get("P009")}
	case types.ValidationModeStrict:
		return []PrincipleChecker{
			r.Get("P001"),
			r.Get("P002"),
			r.Get("P003"),
			r.Get("P004"),
			r.Get("P005"),
			r.Get("P008"),
		}
	default:
		return []PrincipleChecker{r.Get("P001")}
	}
}
