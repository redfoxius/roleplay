package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/redfoxius/roleplay/services/game-server/internal/config"
	"github.com/redfoxius/roleplay/services/game-server/internal/database"
	"github.com/redfoxius/roleplay/services/game-server/internal/game"
	"github.com/redfoxius/roleplay/services/game-server/internal/middleware"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// Initialize Redis connection
	redis, err := database.NewRedisDB(cfg.RedisURL, "", 0)
	if err != nil {
		log.Fatalf("Failed to connect to Redis: %v", err)
	}

	// Initialize repository
	repo := database.NewRepository(redis)

	// Initialize game server
	server := game.NewGameServer(repo)

	// Initialize game handler
	handler := game.NewHandler(server)

	// Create HTTP server
	r := mux.NewRouter()

	authMiddleware := middleware.NewAuthMiddleware(cfg.AuthServiceURL)
	authMiddleware.Register(r)

	handler.RegisterRoutes(r)

	log.Printf("Game server starting on port %s", cfg.Port)
	if err := http.ListenAndServe(":"+cfg.Port, r); err != nil {
		log.Fatal(err)
	}
}
