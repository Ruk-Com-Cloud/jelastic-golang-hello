package main

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	app := fiber.New()

	app.Use(logger.New(logger.Config{
		Format: "${time} ${ip} ${method} ${path}?${queryParams} ${status} - ${latency}\n",
	}))

	app.Get("/", func(c *fiber.Ctx) error {
		message := "Built with Love. Run with Ruk Com.: "
		if testMsg := os.Getenv("TEST_MSG"); testMsg != "" {
			message = message + " " + testMsg
		}
		if queryParam := c.Query("message"); queryParam != "" {
			message = message + " " + queryParam
		}
		return c.JSON(fiber.Map{
			"message": message,
		})
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	app.Listen(":" + port)
}
