package config

import (
	"os"
	"testing"
)

func TestLoad(t *testing.T) {
	// Test default values
	cfg, err := Load()
	if err != nil {
		t.Fatalf("Load() error = %v", err)
	}

	if cfg.Port != "8080" {
		t.Errorf("Load() Port = %v, want %v", cfg.Port, "8080")
	}

	if cfg.RedisURL != "redis:6379" {
		t.Errorf("Load() RedisURL = %v, want %v", cfg.RedisURL, "redis:6379")
	}

	if cfg.AuthServiceURL != "http://auth-service:8081" {
		t.Errorf("Load() AuthServiceURL = %v, want %v", cfg.AuthServiceURL, "http://auth-service:8081")
	}

	if cfg.ChatServiceURL != "http://chat-service:8082" {
		t.Errorf("Load() ChatServiceURL = %v, want %v", cfg.ChatServiceURL, "http://chat-service:8082")
	}

	if cfg.MaxPlayers != 100 {
		t.Errorf("Load() MaxPlayers = %v, want %v", cfg.MaxPlayers, 100)
	}

	if cfg.WorldSize != 1000 {
		t.Errorf("Load() WorldSize = %v, want %v", cfg.WorldSize, 1000)
	}

	if cfg.TickRate != 60 {
		t.Errorf("Load() TickRate = %v, want %v", cfg.TickRate, 60)
	}

	// Test custom values
	os.Setenv("PORT", "9090")
	os.Setenv("REDIS_URL", "localhost:6379")
	os.Setenv("AUTH_SERVICE_URL", "http://localhost:8081")
	os.Setenv("CHAT_SERVICE_URL", "http://localhost:8082")
	os.Setenv("MAX_PLAYERS", "50")
	os.Setenv("WORLD_SIZE", "500")
	os.Setenv("TICK_RATE", "30")

	cfg, err = Load()
	if err != nil {
		t.Fatalf("Load() error = %v", err)
	}

	if cfg.Port != "9090" {
		t.Errorf("Load() Port = %v, want %v", cfg.Port, "9090")
	}

	if cfg.RedisURL != "localhost:6379" {
		t.Errorf("Load() RedisURL = %v, want %v", cfg.RedisURL, "localhost:6379")
	}

	if cfg.AuthServiceURL != "http://localhost:8081" {
		t.Errorf("Load() AuthServiceURL = %v, want %v", cfg.AuthServiceURL, "http://localhost:8081")
	}

	if cfg.ChatServiceURL != "http://localhost:8082" {
		t.Errorf("Load() ChatServiceURL = %v, want %v", cfg.ChatServiceURL, "http://localhost:8082")
	}

	if cfg.MaxPlayers != 50 {
		t.Errorf("Load() MaxPlayers = %v, want %v", cfg.MaxPlayers, 50)
	}

	if cfg.WorldSize != 500 {
		t.Errorf("Load() WorldSize = %v, want %v", cfg.WorldSize, 500)
	}

	if cfg.TickRate != 30 {
		t.Errorf("Load() TickRate = %v, want %v", cfg.TickRate, 30)
	}
}

func TestValidate(t *testing.T) {
	tests := []struct {
		name    string
		cfg     *Config
		wantErr bool
	}{
		{
			name: "valid config",
			cfg: &Config{
				Port:               "8080",
				RedisURL:           "redis:6379",
				AuthServiceURL:     "http://auth-service:8081",
				ChatServiceURL:     "http://chat-service:8082",
				CorsAllowedOrigins: []string{"http://localhost:3000"},
				MaxPlayers:         100,
				WorldSize:          1000,
				TickRate:           20,
			},
			wantErr: false,
		},
		{
			name: "missing port",
			cfg: &Config{
				RedisURL:           "redis:6379",
				AuthServiceURL:     "http://auth-service:8081",
				ChatServiceURL:     "http://chat-service:8082",
				CorsAllowedOrigins: []string{"http://localhost:3000"},
				MaxPlayers:         100,
				WorldSize:          1000,
				TickRate:           20,
			},
			wantErr: true,
		},
		{
			name: "invalid redis url",
			cfg: &Config{
				Port:               "8080",
				RedisURL:           "invalid-url",
				AuthServiceURL:     "http://auth-service:8081",
				ChatServiceURL:     "http://chat-service:8082",
				CorsAllowedOrigins: []string{"http://localhost:3000"},
				MaxPlayers:         100,
				WorldSize:          1000,
				TickRate:           20,
			},
			wantErr: true,
		},
		{
			name: "empty cors origins",
			cfg: &Config{
				Port:               "8080",
				RedisURL:           "redis:6379",
				AuthServiceURL:     "http://auth-service:8081",
				ChatServiceURL:     "http://chat-service:8082",
				CorsAllowedOrigins: []string{},
				MaxPlayers:         100,
				WorldSize:          1000,
				TickRate:           20,
			},
			wantErr: true,
		},
		{
			name: "invalid max players",
			cfg: &Config{
				Port:               "8080",
				RedisURL:           "redis:6379",
				AuthServiceURL:     "http://auth-service:8081",
				ChatServiceURL:     "http://chat-service:8082",
				CorsAllowedOrigins: []string{"http://localhost:3000"},
				MaxPlayers:         0,
				WorldSize:          1000,
				TickRate:           20,
			},
			wantErr: true,
		},
		{
			name: "invalid world size",
			cfg: &Config{
				Port:               "8080",
				RedisURL:           "redis:6379",
				AuthServiceURL:     "http://auth-service:8081",
				ChatServiceURL:     "http://chat-service:8082",
				CorsAllowedOrigins: []string{"http://localhost:3000"},
				MaxPlayers:         100,
				WorldSize:          0,
				TickRate:           20,
			},
			wantErr: true,
		},
		{
			name: "invalid tick rate",
			cfg: &Config{
				Port:               "8080",
				RedisURL:           "redis:6379",
				AuthServiceURL:     "http://auth-service:8081",
				ChatServiceURL:     "http://chat-service:8082",
				CorsAllowedOrigins: []string{"http://localhost:3000"},
				MaxPlayers:         100,
				WorldSize:          1000,
				TickRate:           0,
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.cfg.validate()
			if (err != nil) != tt.wantErr {
				t.Errorf("Config.validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
