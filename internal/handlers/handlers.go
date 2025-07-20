package handlers

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"jelastic-golang-hello/internal/config"
)

func SetupRoutes(app *fiber.App, cfg *config.Config) {
	setupMiddleware(app)
	setupAPIRoutes(app)
	setupHealthRoute(app, cfg)
}

func setupMiddleware(app *fiber.App) {
	app.Use(logger.New(logger.Config{
		Format: "${time} ${ip} ${method} ${path}?${queryParams} ${status} - ${latency}\n",
	}))
}

func setupAPIRoutes(app *fiber.App) {
	// Simple in-memory data for demonstration
	app.Get("/api/info", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"service":     "jelastic-golang-hello",
			"version":     "1.0.0",
			"timestamp":   time.Now().UTC(),
			"environment": "http-only",
		})
	})

	app.Get("/api/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status":    "healthy",
			"timestamp": time.Now().UTC(),
			"uptime":    time.Since(startTime).String(),
		})
	})

	app.Get("/api/echo", func(c *fiber.Ctx) error {
		message := c.Query("message", "Hello World!")
		return c.JSON(fiber.Map{
			"echo":      message,
			"timestamp": time.Now().UTC(),
		})
	})
}

func setupHealthRoute(app *fiber.App, cfg *config.Config) {
	app.Get("/", func(c *fiber.Ctx) error {
		message := "Built with Love. Run with Ruk Com."
		
		// Add test message from config if available
		if cfg != nil && cfg.App.TestMessage != "" {
			message = message + " " + cfg.App.TestMessage
		}
		
		// Add query parameter if provided
		if queryParam := c.Query("message"); queryParam != "" {
			message = message + " " + queryParam
		}
		
		response := fiber.Map{
			"message": message,
		}
		
		// Add additional info if config is available
		if cfg != nil {
			response["environment"] = cfg.App.Environment
			response["version"] = "1.0.0"
			response["mode"] = "http-only"
		}
		
		return c.JSON(response)
	})
}

var startTime = time.Now()