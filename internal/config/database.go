package config

import (
	"fmt"
	"log"
	"os"

	"BlogManagment/internal/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// DatabaseConfig holds database configuration
type DatabaseConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
}

// NewDatabaseConfig creates a new database configuration from environment variables
func NewDatabaseConfig() *DatabaseConfig {
	return &DatabaseConfig{
		Host:     getEnv("DB_HOST", "localhost"),
		Port:     getEnv("DB_PORT", "5432"),
		User:     getEnv("DB_USER", "postgres"),
		Password: getEnv("DB_PASSWORD", "password"),
		DBName:   getEnv("DB_NAME", "blog_management"),
	}
}

// Connect establishes a database connection
func (c *DatabaseConfig) Connect() (*gorm.DB, error) {
	// Prefer DATABASE_URL if set
	databaseURL := os.Getenv("DATABASE_URL")
	var dsn string
	if databaseURL != "" {
		dsn = databaseURL
	} else {
		dsn = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=UTC",
			c.Host, c.User, c.Password, c.DBName, c.Port)
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	// Auto migrate the schema
	if err := db.AutoMigrate(&models.Blog{}); err != nil {
		return nil, fmt.Errorf("failed to migrate database: %w", err)
	}

	log.Println("Database connected and migrated successfully")
	return db, nil
}

// getEnv gets an environment variable or returns a default value
func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
