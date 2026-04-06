package feature

// Flag represents a feature flag
type Flag string

const (
	// Example feature flags
	FeatureHello    Flag = "feature_hello"
	FeatureNewCLI   Flag = "feature_new_cli"
	FeatureMetrics  Flag = "feature_metrics"
	FeatureDebugLog Flag = "feature_debug_log"
)
