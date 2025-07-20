package main

import (
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
)

func main() {
	// Load .env file if it exists
	godotenv.Load()

	// Validate required environment variables
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
		fmt.Println("Warning: PORT not set, using default:", port)
	} else {
		fmt.Println("Using PORT:", port)
	}

	testMsg := os.Getenv("TEST_MSG")
	if testMsg == "" {
		fmt.Println("Info: TEST_MSG not set, will use default message")
	} else {
		fmt.Println("Using TEST_MSG:", testMsg)
	}

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

	app.Listen(":" + port)
}
