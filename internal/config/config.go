package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Server   ServerConfig
	Database DatabaseConfig
}

type ServerConfig struct {
	Port      string
	Env       string
	JwtSecret string
}

type DatabaseConfig struct {
	DBUrl string
}

func Load() (*Config, error) {
	if err := godotenv.Load(); err != nil {
		return nil, err
	}

	return &Config{
		Server: ServerConfig{
			Port:      getEnv("PORT", "8080"),
			Env:       getEnv("ENV", "development"),
			JwtSecret: getEnv("JWT_SECRET", "99322be7-6f9c-4de3-af5c-f024539ab0b2"),
		},
		Database: DatabaseConfig{
			DBUrl: getEnv("DB_URL", ""),
		},
	}, nil
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
