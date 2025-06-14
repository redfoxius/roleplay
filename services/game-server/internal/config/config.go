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
	AuthServiceURL     string
	ChatServiceURL     string
	Debug              bool
	CorsAllowedOrigins []string
	MaxPlayers         int
	WorldSize          int
	TickRate           int
	SaveInterval       time.Duration
	MaxPartySize       int
	MaxInventorySize   int
	MaxChatHistory     int
	MaxQuestLog        int
	MaxFriends         int
}

func Load() (*Config, error) {
	cfg := &Config{
		Port:               getEnv("PORT", "8080"),
		RedisURL:           getEnv("REDIS_URL", "redis:6379"),
		AuthServiceURL:     getEnv("AUTH_SERVICE_URL", "http://auth-service:8081"),
		ChatServiceURL:     getEnv("CHAT_SERVICE_URL", "http://chat-service:8082"),
		Debug:              getEnvBool("DEBUG", false),
		CorsAllowedOrigins: getEnvSlice("CORS_ALLOWED_ORIGINS", []string{"http://localhost:3000"}),
		MaxPlayers:         getEnvInt("MAX_PLAYERS", 100),
		WorldSize:          getEnvInt("WORLD_SIZE", 1000),
		TickRate:           getEnvInt("TICK_RATE", 60),
		SaveInterval:       getEnvDuration("SAVE_INTERVAL", 5*time.Minute),
		MaxPartySize:       getEnvInt("MAX_PARTY_SIZE", 5),
		MaxInventorySize:   getEnvInt("MAX_INVENTORY_SIZE", 50),
		MaxChatHistory:     getEnvInt("MAX_CHAT_HISTORY", 100),
		MaxQuestLog:        getEnvInt("MAX_QUEST_LOG", 20),
		MaxFriends:         getEnvInt("MAX_FRIENDS", 100),
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

	if _, err := url.Parse(c.AuthServiceURL); err != nil {
		return fmt.Errorf("invalid AUTH_SERVICE_URL: %v", err)
	}

	if _, err := url.Parse(c.ChatServiceURL); err != nil {
		return fmt.Errorf("invalid CHAT_SERVICE_URL: %v", err)
	}

	if len(c.CorsAllowedOrigins) == 0 {
		return fmt.Errorf("CORS_ALLOWED_ORIGINS must contain at least one origin")
	}

	if c.MaxPlayers <= 0 {
		return fmt.Errorf("MAX_PLAYERS must be positive")
	}

	if c.WorldSize <= 0 {
		return fmt.Errorf("WORLD_SIZE must be positive")
	}

	if c.TickRate <= 0 {
		return fmt.Errorf("TICK_RATE must be positive")
	}

	if c.SaveInterval <= 0 {
		return fmt.Errorf("SAVE_INTERVAL must be positive")
	}

	if c.MaxPartySize <= 0 {
		return fmt.Errorf("MAX_PARTY_SIZE must be positive")
	}

	if c.MaxInventorySize <= 0 {
		return fmt.Errorf("MAX_INVENTORY_SIZE must be positive")
	}

	if c.MaxChatHistory <= 0 {
		return fmt.Errorf("MAX_CHAT_HISTORY must be positive")
	}

	if c.MaxQuestLog <= 0 {
		return fmt.Errorf("MAX_QUEST_LOG must be positive")
	}

	if c.MaxFriends <= 0 {
		return fmt.Errorf("MAX_FRIENDS must be positive")
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
