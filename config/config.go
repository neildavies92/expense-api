package config

import (
	"fmt"
	"log/slog"
	"os"
)

type Config struct {
	Port     string
	Database DatabaseConfig
}

type DatabaseConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
	SSLMode  string
}

func Load() (*Config, error) {
	slog.Info("loading configuration")

	cfg := &Config{
		Port: getEnvOrDefault("PORT", "8080"),
		Database: DatabaseConfig{
			Host:     getEnvOrDefault("DB_HOST", "localhost"),
			Port:     getEnvOrDefault("DB_PORT", "5432"),
			User:     getEnvOrDefault("DB_USER", "postgres"),
			Password: getEnvOrDefault("DB_PASSWORD", "postgres"),
			DBName:   getEnvOrDefault("DB_NAME", "expense-api"),
			SSLMode:  getEnvOrDefault("DB_SSLMODE", "disable"),
		},
	}

	slog.Info("configuration loaded successfully",
		"port", cfg.Port,
		"db_host", cfg.Database.Host,
		"db_port", cfg.Database.Port,
		"db_name", cfg.Database.DBName,
		"db_ssl_mode", cfg.Database.SSLMode,
	)

	return cfg, nil
}

func getEnvOrDefault(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		slog.Debug("environment variable found",
			"key", key,
			"value", maskSensitiveValue(key, value),
		)
		return value
	}
	slog.Debug("environment variable not found, using default",
		"key", key,
		"default_value", maskSensitiveValue(key, defaultValue),
	)
	return defaultValue
}

// maskSensitiveValue masks sensitive configuration values in logs
func maskSensitiveValue(key, value string) string {
	sensitiveKeys := []string{"DB_PASSWORD", "password"}
	for _, sensitiveKey := range sensitiveKeys {
		if key == sensitiveKey {
			return "***"
		}
	}
	return value
}

func (c *DatabaseConfig) ConnectionString() string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		c.Host, c.Port, c.User, c.Password, c.DBName, c.SSLMode)
}
