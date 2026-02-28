package config

import (
	"fmt"
	"os"
	"strconv"
)

// Config holds app settings loaded from environment variables.
type Config struct {
	AppEnv  string // App environment (e.g. dev, prod)
	AppPort int    // HTTP port to listen on

	DatabaseURL string // Database connection string (required)

	JWTSecret           string // Secret key for signing JWTs (required)
	JWTExpiresInSeconds int    // JWT lifetime in seconds
}

// Load reads config from env vars and applies safe defaults.
// It returns an error if required values are missing.
func Load() (*Config, error) {
	cfg := &Config{}

	// Use defaults for non-critical values.
	cfg.AppEnv = getEnv("APP_ENV", "dev")
	cfg.AppPort = getEnvInt("APP_PORT", 8080)

	// Prefer a single DATABASE_URL (works well with Docker and deployments).
	cfg.DatabaseURL = os.Getenv("DATABASE_URL")
	if cfg.DatabaseURL == "" {
		return nil, fmt.Errorf("DATABASE_URL is required")
	}

	// JWT settings are required for auth to work.
	cfg.JWTSecret = os.Getenv("JWT_SECRET")
	if cfg.JWTSecret == "" {
		return nil, fmt.Errorf("JWT_SECRET is required")
	}

	// Default token expiry: 1 hour.
	cfg.JWTExpiresInSeconds = getEnvInt("JWT_EXPIRES_IN_SECONDS", 3600)

	return cfg, nil
}

// getEnv returns env value or a default if missing.
func getEnv(key, def string) string {
	v := os.Getenv(key)
	if v == "" {
		return def
	}
	return v
}

// getEnvInt returns env value as int or a default if missing/invalid.
func getEnvInt(key string, def int) int {
	v := os.Getenv(key)
	if v == "" {
		return def
	}
	n, err := strconv.Atoi(v)
	if err != nil {
		return def
	}
	return n
}
