package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type AppConfig struct {
	PORT        string
	MONGODB_URL string
}

func Load() (*AppConfig, error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
		return nil, err
	}

	cfg := &AppConfig{
		PORT:        os.Getenv("PORT"),
		MONGODB_URL: os.Getenv("MONGODB_URL"),
	}

	return cfg, nil
}

type Response struct {
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}
