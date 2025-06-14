package config

import (
	"os"
	"testing"
)

func TestLoad(t *testing.T) {
	// Test default values
	cfg := Load()

	if cfg.Port != "8081" {
		t.Errorf("Load() Port = %v, want %v", cfg.Port, "8081")
	}

	if cfg.JWTSecret != "your-secret-key" {
		t.Errorf("Load() JWTSecret = %v, want %v", cfg.JWTSecret, "your-secret-key")
	}

	if cfg.RedisURL != "redis:6379" {
		t.Errorf("Load() RedisURL = %v, want %v", cfg.RedisURL, "redis:6379")
	}

	// Test custom values
	os.Setenv("PORT", "9090")
	os.Setenv("JWT_SECRET", "test-secret")
	os.Setenv("REDIS_URL", "localhost:6379")

	cfg = Load()

	if cfg.Port != "9090" {
		t.Errorf("Load() Port = %v, want %v", cfg.Port, "9090")
	}

	if cfg.JWTSecret != "test-secret" {
		t.Errorf("Load() JWTSecret = %v, want %v", cfg.JWTSecret, "test-secret")
	}

	if cfg.RedisURL != "localhost:6379" {
		t.Errorf("Load() RedisURL = %v, want %v", cfg.RedisURL, "localhost:6379")
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
				Port:               "8081",
				JWTSecret:          "test-secret",
				RedisURL:           "redis:6379",
				CorsAllowedOrigins: []string{"http://localhost:3000"},
			},
			wantErr: false,
		},
		{
			name: "missing port",
			cfg: &Config{
				JWTSecret:          "test-secret",
				RedisURL:           "redis:6379",
				CorsAllowedOrigins: []string{"http://localhost:3000"},
			},
			wantErr: true,
		},
		{
			name: "invalid redis url",
			cfg: &Config{
				Port:               "8081",
				JWTSecret:          "test-secret",
				RedisURL:           "invalid-url",
				CorsAllowedOrigins: []string{"http://localhost:3000"},
			},
			wantErr: true,
		},
		{
			name: "empty cors origins",
			cfg: &Config{
				Port:               "8081",
				JWTSecret:          "test-secret",
				RedisURL:           "redis:6379",
				CorsAllowedOrigins: []string{},
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.cfg.Validate()
			if (err != nil) != tt.wantErr {
				t.Errorf("Config.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
