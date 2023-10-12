package config

import (
	"log"

	"github.com/joho/godotenv"
)

// LoadEnv loads the environment variables from .env file
func LoadEnv() error {
	err := godotenv.Load(".env")
	if err != nil {
		log.Println("Failed to load .env file")
	}

	return nil
}
