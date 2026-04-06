package feature_test

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/PlatformStackPulse/go-template/internal/feature"
)

func TestFeatureManager(t *testing.T) {
	// Save original env
	originalEnv := os.Getenv("FEATURE_FLAGS")
	defer func() {
		if originalEnv != "" {
			os.Setenv("FEATURE_FLAGS", originalEnv)
		} else {
			os.Unsetenv("FEATURE_FLAGS")
		}
	}()

	tests := []struct {
		name     string
		env      string
		flag     feature.Flag
		expected bool
	}{
		{
			name:     "enabled feature",
			env:      "feature_hello=true",
			flag:     feature.FeatureHello,
			expected: true,
		},
		{
			name:     "disabled feature",
			env:      "feature_hello=false",
			flag:     feature.FeatureHello,
			expected: false,
		},
		{
			name:     "multiple features",
			env:      "feature_hello=true,feature_new_cli=false",
			flag:     feature.FeatureHello,
			expected: true,
		},
		{
			name:     "nonexistent feature",
			env:      "feature_hello=true",
			flag:     feature.FeatureMetrics,
			expected: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			os.Setenv("FEATURE_FLAGS", tc.env)
			fm := feature.NewManager()
			result := fm.IsEnabled(tc.flag)
			assert.Equal(t, tc.expected, result)
		})
	}
}

func TestFeatureManagerSet(t *testing.T) {
	os.Unsetenv("FEATURE_FLAGS")
	fm := feature.NewManager()

	fm.Set(feature.FeatureHello, true)
	assert.True(t, fm.IsEnabled(feature.FeatureHello))

	fm.Set(feature.FeatureHello, false)
	assert.False(t, fm.IsEnabled(feature.FeatureHello))
}
