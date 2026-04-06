package health

import (
	"context"
)

// Status represents health status
type Status string

const (
	StatusHealthy   Status = "healthy"
	StatusUnhealthy Status = "unhealthy"
	StatusDegraded  Status = "degraded"
)

// Check performs a health check
type Check interface {
	Name() string
	Check(ctx context.Context) error
}

// Handler manages health checks
type Handler struct {
	checks map[string]Check
}

// NewHandler creates a new health handler
func NewHandler() *Handler {
	return &Handler{
		checks: make(map[string]Check),
	}
}

// Register registers a health check
func (h *Handler) Register(check Check) {
	h.checks[check.Name()] = check
}

// CheckAll performs all health checks
func (h *Handler) CheckAll(ctx context.Context) map[string]error {
	results := make(map[string]error)
	for name, check := range h.checks {
		results[name] = check.Check(ctx)
	}
	return results
}

// IsHealthy checks if all checks pass
func (h *Handler) IsHealthy(ctx context.Context) bool {
	for _, err := range h.CheckAll(ctx) {
		if err != nil {
			return false
		}
	}
	return true
}
