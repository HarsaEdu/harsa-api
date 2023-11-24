package configs

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type AppConfig struct {
	MySQL      MySQLConfig
	Cloudinary CloudinaryConfig
	OpenAI     OpenAI
}

type MySQLConfig struct {
	Username string
	Password string
	Host     string
	Port     string
	Database string
}

type CloudinaryConfig struct {
	CloudName string
	ApiKey    string
	ApiSecret string
}

type OpenAI struct {
	ApiKey string
	AssistantID string
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
		Cloudinary: CloudinaryConfig{
			CloudName: os.Getenv("CLOUDINARY_CLOUD_NAME"),
			ApiKey:    os.Getenv("CLOUDINARY_API_KEY"),
			ApiSecret: os.Getenv("CLOUDINARY_API_SECRET"),
		},
		OpenAI: OpenAI{
			ApiKey: os.Getenv("OPENAI_API_KEY"),
			AssistantID: os.Getenv("OPENAI_ASSISTANT_ID"),
		},
	}, nil
}
