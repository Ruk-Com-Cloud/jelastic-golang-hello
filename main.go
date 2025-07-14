package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"jelastic-golang-hello/internal/adapters/http"
	"jelastic-golang-hello/internal/application"
	"jelastic-golang-hello/internal/config"
	"jelastic-golang-hello/internal/infrastructure"
	"jelastic-golang-hello/internal/seeder"
)

func main() {
	// Parse command line flags
	var (
		runSeeders = flag.Bool("seed", false, "Run database seeders on startup")
		seedOnly   = flag.Bool("seed-only", false, "Run seeders and exit (don't start server)")
	)
	flag.Parse()

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

	// Initialize database
	db, err := infrastructure.NewDatabaseFromConfig(&cfg.Database)
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	fmt.Println("Database connected and migrated successfully")

	// Run seeders if requested
	if *runSeeders || *seedOnly {
		fmt.Println("Running database seeders...")
		manager := seeder.SetupSeeders(db)
		ctx := context.Background()
		
		if err := manager.RunAll(ctx); err != nil {
			log.Fatalf("Failed to run seeders: %v", err)
		}
		
		if *seedOnly {
			fmt.Println("Seeders completed. Exiting...")
			os.Exit(0)
		}
	}

	// Initialize repositories
	userRepo := infrastructure.NewPostgresUserRepository(db)

	// Initialize services
	userService := application.NewUserService(userRepo)

	// Initialize handlers
	userHandler := http.NewUserHandler(userService)

	// Initialize Fiber app
	app := fiber.New()

	// Setup routes with config
	http.SetupRoutesWithConfig(app, userHandler, cfg)

	// Start server
	fmt.Printf("Starting server on %s\n", cfg.Server.GetAddress())
	app.Listen(":" + cfg.Server.Port)
}
