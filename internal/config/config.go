package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	Server   ServerConfig   `mapstructure:"server"`
	Database DatabaseConfig `mapstructure:"database"`
	App      AppConfig      `mapstructure:"app"`
}

type ServerConfig struct {
	Port string `mapstructure:"port"`
	Host string `mapstructure:"host"`
}

type DatabaseConfig struct {
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	DBName   string `mapstructure:"dbname"`
	SSLMode  string `mapstructure:"sslmode"`
	URL      string `mapstructure:"url"`
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

	// Database defaults
	viper.SetDefault("database.host", "localhost")
	viper.SetDefault("database.port", "5432")
	viper.SetDefault("database.user", "postgres")
	viper.SetDefault("database.password", "password")
	viper.SetDefault("database.dbname", "testdb")
	viper.SetDefault("database.sslmode", "disable")
	viper.SetDefault("database.url", "")

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

	if dbHost := viper.GetString("DB_HOST"); dbHost != "" {
		config.Database.Host = dbHost
	}
	if dbPort := viper.GetString("DB_PORT"); dbPort != "" {
		config.Database.Port = dbPort
	}
	if dbUser := viper.GetString("DB_USER"); dbUser != "" {
		config.Database.User = dbUser
	}
	if dbPassword := viper.GetString("DB_PASSWORD"); dbPassword != "" {
		config.Database.Password = dbPassword
	}
	if dbName := viper.GetString("DB_NAME"); dbName != "" {
		config.Database.DBName = dbName
	}
	if sslMode := viper.GetString("DB_SSLMODE"); sslMode != "" {
		config.Database.SSLMode = sslMode
	}
	if dbURL := viper.GetString("DATABASE_URL"); dbURL != "" {
		config.Database.URL = dbURL
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

	if c.Database.URL == "" {
		if c.Database.Host == "" {
			return fmt.Errorf("database host is required")
		}
		if c.Database.User == "" {
			return fmt.Errorf("database user is required")
		}
		if c.Database.DBName == "" {
			return fmt.Errorf("database name is required")
		}
	}

	return nil
}

func (c *Config) Print() {
	fmt.Printf("Configuration loaded:\n")
	fmt.Printf("  Server:\n")
	fmt.Printf("    Host: %s\n", c.Server.Host)
	fmt.Printf("    Port: %s\n", c.Server.Port)
	fmt.Printf("  Database:\n")
	if c.Database.URL != "" {
		fmt.Printf("    URL: %s\n", maskPassword(c.Database.URL))
	} else {
		fmt.Printf("    Host: %s\n", c.Database.Host)
		fmt.Printf("    Port: %s\n", c.Database.Port)
		fmt.Printf("    User: %s\n", c.Database.User)
		fmt.Printf("    Database: %s\n", c.Database.DBName)
		fmt.Printf("    SSL Mode: %s\n", c.Database.SSLMode)
	}
	fmt.Printf("  App:\n")
	fmt.Printf("    Environment: %s\n", c.App.Environment)
	if c.App.TestMessage != "" {
		fmt.Printf("    Test Message: %s\n", c.App.TestMessage)
	} else {
		fmt.Printf("    Test Message: (not set)\n")
	}
}

func maskPassword(url string) string {
	// Simple password masking for logs
	return "postgres://user:***@host:port/db"
}
