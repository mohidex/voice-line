package main

import (
	"log"
	"time"

	"github.com/mohidex/voice-line/config"
	"github.com/mohidex/voice-line/internal/models"
	"github.com/mohidex/voice-line/internal/repositories"
	"github.com/mohidex/voice-line/internal/server"
	"github.com/mohidex/voice-line/pkg/auth"
	database "github.com/mohidex/voice-line/pkg/db"
)

func main() {

	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Could not load config: %v", err)
	}
	db, err := database.NewPostgresDB(cfg.DB.Host, cfg.DB.User, cfg.DB.Password, cfg.DB.Name, 5432)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	if err = db.GetDB().AutoMigrate(&models.User{}); err != nil {
		log.Fatal("Failed to auto-migrate:", err)
	}

	// Create the server options
	opts := &server.Opt{
		Port:         "8080",
		Environment:  "development",
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	// Initialize the user repository
	userRepo := repositories.NewPostgresUserRepository(db.GetDB())

	// Initialize the authenticator
	authenticator := auth.NewFirebaseAuth(cfg.Firebase.APIKey, cfg.Firebase.BaseURL)

	// Initialize the server
	srv := server.NewServer(opts, userRepo, authenticator)

	// Start the server
	if err := srv.Start(); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
