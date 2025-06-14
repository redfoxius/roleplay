package config

import (
	"errors"
	"os"
	"strconv"
)

type Config struct {
	Port               string
	JWTSecret          string
	RedisURL           string
	Debug              bool
	CorsAllowedOrigins []string
}

func Load() *Config {
	return &Config{
		Port:               getEnv("PORT", "8081"),
		JWTSecret:          getEnv("JWT_SECRET", "your-secret-key"),
		RedisURL:           getEnv("REDIS_URL", "redis:6379"),
		Debug:              getEnvBool("DEBUG", false),
		CorsAllowedOrigins: getEnvSlice("CORS_ALLOWED_ORIGINS", []string{"http://localhost:3000"}),
	}
}

func (c *Config) Validate() error {
	if c.Port == "" {
		return errors.New("port is required")
	}

	if c.JWTSecret == "" {
		return errors.New("JWT secret is required")
	}

	if c.RedisURL == "" {
		return errors.New("Redis URL is required")
	}

	if len(c.CorsAllowedOrigins) == 0 {
		return errors.New("at least one CORS allowed origin is required")
	}

	return nil
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

func getEnvBool(key string, defaultValue bool) bool {
	if value, exists := os.LookupEnv(key); exists {
		boolValue, err := strconv.ParseBool(value)
		if err == nil {
			return boolValue
		}
	}
	return defaultValue
}

func getEnvSlice(key string, defaultValue []string) []string {
	if value, exists := os.LookupEnv(key); exists {
		return []string{value} // For now, just split by comma if needed
	}
	return defaultValue
}
