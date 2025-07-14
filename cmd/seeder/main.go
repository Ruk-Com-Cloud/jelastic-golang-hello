package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"jelastic-golang-hello/internal/config"
	"jelastic-golang-hello/internal/infrastructure"
	"jelastic-golang-hello/internal/seeder"
)

func main() {
	var (
		action     = flag.String("action", "seed", "Action to perform: seed, rollback, list")
		seederName = flag.String("seeder", "", "Specific seeder to run (optional)")
		help       = flag.Bool("help", false, "Show help message")
	)
	flag.Parse()

	if *help {
		showHelp()
		return
	}

	// Load configuration
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	// Initialize database
	db, err := infrastructure.NewDatabaseFromConfig(&cfg.Database)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Setup seeder manager
	manager := seeder.SetupSeeders(db)
	ctx := context.Background()

	// Execute action
	switch strings.ToLower(*action) {
	case "seed":
		if *seederName != "" {
			err = manager.RunSeeder(ctx, *seederName)
		} else {
			err = manager.RunAll(ctx)
		}
	case "rollback":
		if *seederName != "" {
			fmt.Println("Rollback of specific seeders is not supported yet")
			os.Exit(1)
		} else {
			err = manager.RollbackAll(ctx)
		}
	case "list":
		listSeeders(manager)
		return
	default:
		fmt.Printf("Unknown action: %s\n", *action)
		showHelp()
		os.Exit(1)
	}

	if err != nil {
		log.Fatalf("Seeder action failed: %v", err)
	}
}

func showHelp() {
	fmt.Println("Database Seeder Tool")
	fmt.Println("Usage: go run cmd/seeder/main.go [options]")
	fmt.Println()
	fmt.Println("Options:")
	fmt.Println("  -action string")
	fmt.Println("        Action to perform: seed, rollback, list (default \"seed\")")
	fmt.Println("  -seeder string")
	fmt.Println("        Specific seeder to run (optional)")
	fmt.Println("  -help")
	fmt.Println("        Show this help message")
	fmt.Println()
	fmt.Println("Examples:")
	fmt.Println("  go run cmd/seeder/main.go -action=seed")
	fmt.Println("  go run cmd/seeder/main.go -action=seed -seeder=UserSeeder")
	fmt.Println("  go run cmd/seeder/main.go -action=rollback")
	fmt.Println("  go run cmd/seeder/main.go -action=list")
	fmt.Println()
	fmt.Println("Environment Variables:")
	fmt.Println("  Set database configuration via environment variables or .env file")
	fmt.Println("  See .env.example for available options")
}

func listSeeders(manager *seeder.SeederManager) {
	seeders := manager.ListSeeders()
	fmt.Println("Available Seeders:")
	for i, name := range seeders {
		fmt.Printf("  %d. %s\n", i+1, name)
	}
}