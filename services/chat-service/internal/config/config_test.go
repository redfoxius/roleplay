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

	if cfg.Port != "8082" {
		t.Errorf("Load() Port = %v, want %v", cfg.Port, "8082")
	}

	if cfg.RedisURL != "redis:6379" {
		t.Errorf("Load() RedisURL = %v, want %v", cfg.RedisURL, "redis:6379")
	}

	if cfg.WSMaxConnections != 1000 {
		t.Errorf("Load() WSMaxConnections = %v, want %v", cfg.WSMaxConnections, 1000)
	}

	if cfg.WSMessageSizeLimit != 4096 {
		t.Errorf("Load() WSMessageSizeLimit = %v, want %v", cfg.WSMessageSizeLimit, 4096)
	}

	// Test custom values
	os.Setenv("PORT", "9090")
	os.Setenv("REDIS_URL", "localhost:6379")
	os.Setenv("WS_MAX_CONNECTIONS", "500")
	os.Setenv("WS_MESSAGE_SIZE_LIMIT", "8192")

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

	if cfg.WSMaxConnections != 500 {
		t.Errorf("Load() WSMaxConnections = %v, want %v", cfg.WSMaxConnections, 500)
	}

	if cfg.WSMessageSizeLimit != 8192 {
		t.Errorf("Load() WSMessageSizeLimit = %v, want %v", cfg.WSMessageSizeLimit, 8192)
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
				Port:               "8082",
				RedisURL:           "redis:6379",
				CorsAllowedOrigins: []string{"http://localhost:3000"},
				WSMaxConnections:   1000,
				WSMessageSizeLimit: 4096,
			},
			wantErr: false,
		},
		{
			name: "missing port",
			cfg: &Config{
				RedisURL:           "redis:6379",
				CorsAllowedOrigins: []string{"http://localhost:3000"},
				WSMaxConnections:   1000,
				WSMessageSizeLimit: 4096,
			},
			wantErr: true,
		},
		{
			name: "invalid redis url",
			cfg: &Config{
				Port:               "8082",
				RedisURL:           "invalid-url",
				CorsAllowedOrigins: []string{"http://localhost:3000"},
				WSMaxConnections:   1000,
				WSMessageSizeLimit: 4096,
			},
			wantErr: true,
		},
		{
			name: "empty cors origins",
			cfg: &Config{
				Port:               "8082",
				RedisURL:           "redis:6379",
				CorsAllowedOrigins: []string{},
				WSMaxConnections:   1000,
				WSMessageSizeLimit: 4096,
			},
			wantErr: true,
		},
		{
			name: "invalid ws max connections",
			cfg: &Config{
				Port:               "8082",
				RedisURL:           "redis:6379",
				CorsAllowedOrigins: []string{"http://localhost:3000"},
				WSMaxConnections:   0,
				WSMessageSizeLimit: 4096,
			},
			wantErr: true,
		},
		{
			name: "invalid ws message size limit",
			cfg: &Config{
				Port:               "8082",
				RedisURL:           "redis:6379",
				CorsAllowedOrigins: []string{"http://localhost:3000"},
				WSMaxConnections:   1000,
				WSMessageSizeLimit: 0,
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
