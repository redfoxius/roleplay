package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/redfoxius/roleplay/services/auth-service/internal/auth"
	"github.com/redfoxius/roleplay/services/auth-service/internal/config"
)

func main() {
	cfg := config.Load()

	r := mux.NewRouter()

	authHandler := auth.NewHandler(cfg.JWTSecret)
	authHandler.Register(r)

	log.Printf("Auth service starting on port %s", cfg.Port)
	if err := http.ListenAndServe(":"+cfg.Port, r); err != nil {
		log.Fatal(err)
	}
}
