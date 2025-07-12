package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/spf13/viper"
)

func main() {
	// Configure Viper
	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()

	// Set defaults
	viper.SetDefault("PORT", "3000")
	viper.SetDefault("TEST_MSG", "")

	// Read config file if it exists
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			fmt.Println("No .env file found, using environment variables and defaults")
		} else {
			log.Fatalf("Error reading config file: %v", err)
		}
	}

	// Validate and display configuration
	port := viper.GetString("PORT")
	testMsg := viper.GetString("TEST_MSG")

	fmt.Printf("Configuration loaded:\n")
	fmt.Printf("  PORT: %s\n", port)
	if testMsg != "" {
		fmt.Printf("  TEST_MSG: %s\n", testMsg)
	} else {
		fmt.Printf("  TEST_MSG: (not set)\n")
	}

	app := fiber.New()

	app.Use(logger.New(logger.Config{
		Format: "${time} ${ip} ${method} ${path}?${queryParams} ${status} - ${latency}\n",
	}))

	app.Get("/", func(c *fiber.Ctx) error {
		message := "Built with Love. Run with Ruk Com.: "
		if testMsg := viper.GetString("TEST_MSG"); testMsg != "" {
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
