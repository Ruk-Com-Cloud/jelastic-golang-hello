package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	Server ServerConfig `mapstructure:"server"`
	App    AppConfig    `mapstructure:"app"`
}

type ServerConfig struct {
	Port string `mapstructure:"port"`
	Host string `mapstructure:"host"`
}


type AppConfig struct {
	TestMessage string `mapstructure:"test_message"`
	Environment string `mapstructure:"environment"`
}

func Load() (*Config, error) {
	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()

	// Set defaults
	setDefaults()

	// Read config file if it exists
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			fmt.Println("No .env file found, using environment variables and defaults")
		} else {
			return nil, fmt.Errorf("error reading config file: %w", err)
		}
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		return nil, fmt.Errorf("unable to decode config: %w", err)
	}

	// Override with environment variable mappings
	mapEnvironmentVariables(&config)

	return &config, nil
}

func setDefaults() {
	// Server defaults
	viper.SetDefault("server.port", "3000")
	viper.SetDefault("server.host", "0.0.0.0")

	// App defaults
	viper.SetDefault("app.test_message", "")
	viper.SetDefault("app.environment", "development")
}

func mapEnvironmentVariables(config *Config) {
	// Map legacy environment variables to new structure
	if port := viper.GetString("PORT"); port != "" {
		config.Server.Port = port
	}
	if host := viper.GetString("HOST"); host != "" {
		config.Server.Host = host
	}

	if testMsg := viper.GetString("TEST_MSG"); testMsg != "" {
		config.App.TestMessage = testMsg
	}
	if env := viper.GetString("ENVIRONMENT"); env != "" {
		config.App.Environment = env
	}
}

func (c *Config) Validate() error {
	if c.Server.Port == "" {
		return fmt.Errorf("server port is required")
	}

	return nil
}

func (c *Config) Print() {
	fmt.Printf("Configuration loaded:\n")
	fmt.Printf("  Server:\n")
	fmt.Printf("    Host: %s\n", c.Server.Host)
	fmt.Printf("    Port: %s\n", c.Server.Port)
	fmt.Printf("  App:\n")
	fmt.Printf("    Environment: %s\n", c.App.Environment)
	if c.App.TestMessage != "" {
		fmt.Printf("    Test Message: %s\n", c.App.TestMessage)
	} else {
		fmt.Printf("    Test Message: (not set)\n")
	}
}

