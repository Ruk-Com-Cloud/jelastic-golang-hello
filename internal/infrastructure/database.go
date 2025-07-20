package infrastructure

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"jelastic-golang-hello/internal/config"
)

func NewDatabaseFromConfig(cfg *config.DatabaseConfig) (*gorm.DB, error) {
	dsn := cfg.GetDSN()

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	if err := db.AutoMigrate(&UserModel{}); err != nil {
		return nil, fmt.Errorf("failed to migrate database: %w", err)
	}

	return db, nil
}

// Legacy function for backward compatibility
func NewDatabase(host, user, password, dbname, port, sslmode string) (*gorm.DB, error) {
	cfg := &config.DatabaseConfig{
		Host:     host,
		User:     user,
		Password: password,
		DBName:   dbname,
		Port:     port,
		SSLMode:  sslmode,
	}
	return NewDatabaseFromConfig(cfg)
}