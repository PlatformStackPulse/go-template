package feature_test

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/PlatformStackPulse/go-template/internal/feature"
)

func TestFeatureManager(t *testing.T) {
	originalEnv := os.Getenv("FEATURE_FLAGS")
	defer func() {
		if originalEnv != "" {
			_ = os.Setenv("FEATURE_FLAGS", originalEnv)
		} else {
			_ = os.Unsetenv("FEATURE_FLAGS")
		}
	}()

	tests := []struct {
		name     string
		env      string
		flag     feature.Flag
		expected bool
	}{
		{name: "enabled feature", env: "feature_hello=true", flag: feature.FeatureHello, expected: true},
		{name: "disabled feature", env: "feature_hello=false", flag: feature.FeatureHello, expected: false},
		{name: "multiple features", env: "feature_hello=true,feature_new_cli=false", flag: feature.FeatureHello, expected: true},
		{name: "nonexistent feature", env: "feature_hello=true", flag: feature.FeatureMetrics, expected: false},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			_ = os.Setenv("FEATURE_FLAGS", tc.env)
			fm := feature.NewManager()
			assert.Equal(t, tc.expected, fm.IsEnabled(tc.flag))
		})
	}
}

func TestFeatureManagerSet(t *testing.T) {
	_ = os.Unsetenv("FEATURE_FLAGS")
	fm := feature.NewManager()

	fm.Set(feature.FeatureHello, true)
	assert.True(t, fm.IsEnabled(feature.FeatureHello))

	fm.Set(feature.FeatureHello, false)
	assert.False(t, fm.IsEnabled(feature.FeatureHello))
}
