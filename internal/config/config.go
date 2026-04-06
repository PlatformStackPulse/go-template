package config

import (
	"os"
	"strconv"
)

// Config holds application configuration
type Config struct {
	Debug   bool
	AppName string
	Version string
}

// Load loads configuration from environment variables
func Load() *Config {
	return &Config{
		Debug:   getBoolEnv("DEBUG", false),
		AppName: getEnv("APP_NAME", "go-template"),
		Version: getEnv("APP_VERSION", "dev"),
	}
}

// Helper functions for environment variables
func getEnv(key, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultVal
}

func getBoolEnv(key string, defaultVal bool) bool {
	if value, exists := os.LookupEnv(key); exists {
		if b, err := strconv.ParseBool(value); err == nil {
			return b
		}
	}
	return defaultVal
}
