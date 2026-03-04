package config

import (
	"os"
	"github.com/joho/godotenv"
)

type Config struct {
	DbUrl   string
	Port    string
	BaseUrl string
}

func Load() (*Config, error) {
	_ = godotenv.Load()

	dbUrl := os.Getenv("DATABASE_URL")
	if dbUrl == "" {
		return nil, nil
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	baseUrl := os.Getenv("BASE_URL")
	if baseUrl == "" {
		baseUrl = "http://localhost:" + port
	}

	return &Config{
		Port:    port,
		BaseUrl: baseUrl,
		DbUrl:   dbUrl,
	}, nil
}
