package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"roleplay/internal/database"
	"roleplay/internal/game"
)

func main() {
	// Initialize Redis connection
	redisAddr := os.Getenv("REDIS_ADDR")
	if redisAddr == "" {
		redisAddr = "localhost:6379"
	}

	redisPassword := os.Getenv("REDIS_PASSWORD")
	redisDB := 0 // Use default database

	db, err := database.NewRedisDB(redisAddr, redisPassword, redisDB)
	if err != nil {
		log.Fatalf("Failed to connect to Redis: %v", err)
	}
	defer db.Close()

	// Initialize repository
	repo := database.NewRepository(db)

	// Initialize game server
	server := game.NewGameServer(repo)

	// Initialize HTTP handler
	handler := game.NewHandler(server)

	// Create HTTP server
	mux := http.NewServeMux()
	handler.RegisterRoutes(mux)

	// Start HTTP server
	go func() {
		port := os.Getenv("PORT")
		if port == "" {
			port = "8080"
		}

		log.Printf("Starting server on port %s", port)
		if err := http.ListenAndServe(":"+port, mux); err != nil {
			log.Fatalf("Failed to start server: %v", err)
		}
	}()

	// Handle graceful shutdown
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	<-sigChan
	log.Println("Shutting down server...")

	// Save game state before shutdown
	if err := server.SaveGameState(); err != nil {
		log.Printf("Failed to save game state: %v", err)
	}

	log.Println("Server shutdown complete")
}
