package feature

import (
	"os"
	"strings"
)

// Manager manages feature flags
type Manager struct {
	flags map[Flag]bool
}

// NewManager creates a new feature manager
func NewManager() *Manager {
	return &Manager{
		flags: loadFromEnv(),
	}
}

// loadFromEnv loads feature flags from environment variables
// Format: FEATURE_FLAGS=feature_hello=true,feature_new_cli=false
func loadFromEnv() map[Flag]bool {
	flags := make(map[Flag]bool)

	envValue := os.Getenv("FEATURE_FLAGS")
	if envValue == "" {
		return flags
	}

	pairs := strings.Split(envValue, ",")
	for _, p := range pairs {
		kv := strings.Split(p, "=")
		if len(kv) != 2 {
			continue
		}

		key := Flag(strings.TrimSpace(kv[0]))
		val := strings.TrimSpace(kv[1]) == "true"
		flags[key] = val
	}

	return flags
}

// IsEnabled checks if a feature flag is enabled
func (m *Manager) IsEnabled(flag Flag) bool {
	v, ok := m.flags[flag]
	return ok && v
}

// Set sets a flag (useful for testing)
func (m *Manager) Set(flag Flag, enabled bool) {
	m.flags[flag] = enabled
}

// GetAll returns all flags
func (m *Manager) GetAll() map[Flag]bool {
	return m.flags
}
