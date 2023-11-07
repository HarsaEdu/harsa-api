package configs

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type AppConfig struct {
	MySQL MySQLConfig
}

type MySQLConfig struct {
	Username string
	Password string
	Host     string
	Port     string
	Database string
}

func LoadConfig() (*AppConfig, error) {
	_, err := os.Stat(".env")
	if err == nil {
		err := godotenv.Load()
		if err != nil {
			return nil, fmt.Errorf("failed to load environment variables from .env file: %w", err)
		}
	}

	return &AppConfig{
		MySQL: MySQLConfig{
			Username: os.Getenv("DB_USERNAME"),
			Password: os.Getenv("DB_PASSWORD"),
			Host:     os.Getenv("DB_HOST"),
			Port:     os.Getenv("DB_PORT"),
			Database: os.Getenv("DB_NAME"),
		},
	}, nil
}