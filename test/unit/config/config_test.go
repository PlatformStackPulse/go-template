package config_test

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/PlatformStackPulse/go-template/internal/config"
)

func TestLoadDefaults(t *testing.T) {
	_ = os.Unsetenv("DEBUG")
	_ = os.Unsetenv("APP_NAME")
	_ = os.Unsetenv("APP_VERSION")

	cfg := config.Load()

	assert.False(t, cfg.Debug)
	assert.Equal(t, "go-template", cfg.AppName)
	assert.Equal(t, "dev", cfg.Version)
}

func TestLoadFromEnv(t *testing.T) {
	_ = os.Setenv("DEBUG", "true")
	_ = os.Setenv("APP_NAME", "my-app")
	_ = os.Setenv("APP_VERSION", "v1.0.0")
	defer func() {
		_ = os.Unsetenv("DEBUG")
		_ = os.Unsetenv("APP_NAME")
		_ = os.Unsetenv("APP_VERSION")
	}()

	cfg := config.Load()

	assert.True(t, cfg.Debug)
	assert.Equal(t, "my-app", cfg.AppName)
	assert.Equal(t, "v1.0.0", cfg.Version)
}

func TestInvalidBoolEnvFallsBack(t *testing.T) {
	_ = os.Setenv("DEBUG", "notabool")
	defer func() { _ = os.Unsetenv("DEBUG") }()

	cfg := config.Load()
	assert.False(t, cfg.Debug)
}
