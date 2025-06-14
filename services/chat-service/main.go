package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/redfoxius/roleplay/services/chat-service/internal/chat"
)

func main() {
	r := mux.NewRouter()

	chatHandler := chat.NewHandler()
	chatHandler.Register(r)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8082"
	}

	log.Printf("Chat service starting on port %s", port)
	if err := http.ListenAndServe(":"+port, r); err != nil {
		log.Fatal(err)
	}
}
