package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"jelastic-golang-hello/internal/config"
	"jelastic-golang-hello/internal/handlers"
)

func main() {
	// Load configuration
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	// Validate configuration
	if err := cfg.Validate(); err != nil {
		log.Fatalf("Invalid configuration: %v", err)
	}

	// Print configuration
	cfg.Print()

	// Initialize Fiber app
	app := fiber.New()

	// Setup routes
	handlers.SetupRoutes(app, cfg)

	// Start server
	fmt.Printf("Starting HTTP-only server on %s\n", cfg.Server.GetAddress())
	log.Fatal(app.Listen(":" + cfg.Server.Port))
}