package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	APIKey string
	Version string
}

func LoadConfig() (*Config, error) {
	if err := godotenv.Load(); err != nil {
		return nil, fmt.Errorf("failed to load .env file: %v", err)
	}

	apiKey := os.Getenv("api_key")

	version := os.Getenv("riot_version")

	return &Config{
		APIKey: apiKey,
		Version: version,
	}, nil
}