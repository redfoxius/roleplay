package config

import (
	"fmt"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"
)

type Config struct {
	Port               string
	RedisURL           string
	Debug              bool
	CorsAllowedOrigins []string
	WSMaxConnections   int
	WSMessageSizeLimit int
	MessageTTL         time.Duration
	MaxMessageLength   int
	PingInterval       time.Duration
	PongWait           time.Duration
	WriteWait          time.Duration
}

func Load() (*Config, error) {
	cfg := &Config{
		Port:               getEnv("PORT", "8082"),
		RedisURL:           getEnv("REDIS_URL", "redis:6379"),
		Debug:              getEnvBool("DEBUG", false),
		CorsAllowedOrigins: getEnvSlice("CORS_ALLOWED_ORIGINS", []string{"http://localhost:3000"}),
		WSMaxConnections:   getEnvInt("WS_MAX_CONNECTIONS", 1000),
		WSMessageSizeLimit: getEnvInt("WS_MESSAGE_SIZE_LIMIT", 4096),
		MessageTTL:         getEnvDuration("MESSAGE_TTL", 24*time.Hour),
		MaxMessageLength:   getEnvInt("MAX_MESSAGE_LENGTH", 1000),
		PingInterval:       getEnvDuration("PING_INTERVAL", 30*time.Second),
		PongWait:           getEnvDuration("PONG_WAIT", 60*time.Second),
		WriteWait:          getEnvDuration("WRITE_WAIT", 10*time.Second),
	}

	if err := cfg.validate(); err != nil {
		return nil, err
	}

	return cfg, nil
}

func (c *Config) validate() error {
	if c.Port == "" {
		return fmt.Errorf("PORT is required")
	}

	if _, err := url.Parse(c.RedisURL); err != nil {
		return fmt.Errorf("invalid REDIS_URL: %v", err)
	}

	if len(c.CorsAllowedOrigins) == 0 {
		return fmt.Errorf("CORS_ALLOWED_ORIGINS must contain at least one origin")
	}

	if c.WSMaxConnections <= 0 {
		return fmt.Errorf("WS_MAX_CONNECTIONS must be positive")
	}

	if c.WSMessageSizeLimit <= 0 {
		return fmt.Errorf("WS_MESSAGE_SIZE_LIMIT must be positive")
	}

	if c.MessageTTL <= 0 {
		return fmt.Errorf("MESSAGE_TTL must be positive")
	}

	if c.MaxMessageLength <= 0 {
		return fmt.Errorf("MAX_MESSAGE_LENGTH must be positive")
	}

	if c.PingInterval <= 0 {
		return fmt.Errorf("PING_INTERVAL must be positive")
	}

	if c.PongWait <= c.PingInterval {
		return fmt.Errorf("PONG_WAIT must be greater than PING_INTERVAL")
	}

	if c.WriteWait <= 0 {
		return fmt.Errorf("WRITE_WAIT must be positive")
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

func getEnvInt(key string, defaultValue int) int {
	if value, exists := os.LookupEnv(key); exists {
		intValue, err := strconv.Atoi(value)
		if err == nil {
			return intValue
		}
	}
	return defaultValue
}

func getEnvDuration(key string, defaultValue time.Duration) time.Duration {
	if value, exists := os.LookupEnv(key); exists {
		duration, err := time.ParseDuration(value)
		if err == nil {
			return duration
		}
	}
	return defaultValue
}

func getEnvSlice(key string, defaultValue []string) []string {
	if value, exists := os.LookupEnv(key); exists {
		return strings.Split(value, ",")
	}
	return defaultValue
}
