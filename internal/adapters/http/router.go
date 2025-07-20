package http

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"jelastic-golang-hello/internal/config"
)

func SetupRoutes(app *fiber.App, userHandler *UserHandler) {
	setupMiddleware(app)
	setupAPIRoutes(app, userHandler)
	setupHealthRoute(app, nil)
}

func SetupRoutesWithConfig(app *fiber.App, userHandler *UserHandler, cfg *config.Config) {
	setupMiddleware(app)
	setupAPIRoutes(app, userHandler)
	setupHealthRoute(app, cfg)
}

func setupMiddleware(app *fiber.App) {
	app.Use(logger.New(logger.Config{
		Format: "${time} ${ip} ${method} ${path}?${queryParams} ${status} - ${latency}\n",
	}))
}

func setupAPIRoutes(app *fiber.App, userHandler *UserHandler) {
	// API Routes
	app.Post("/users", userHandler.CreateUser)
	app.Get("/users", userHandler.GetUsers)
	app.Get("/users/:id", userHandler.GetUser)
	app.Put("/users/:id", userHandler.UpdateUser)
	app.Delete("/users/:id", userHandler.DeleteUser)
}

func setupHealthRoute(app *fiber.App, cfg *config.Config) {
	app.Get("/", func(c *fiber.Ctx) error {
		message := "Built with Love. Run with Ruk Com.: "
		
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
		}
		
		return c.JSON(response)
	})
}