package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DatabaseURL  string
	DatabaseName string
}

func New(prod bool) *Config {

	if prod {
		// Find .env file
		err := godotenv.Load("config/.env")
		if err != nil {
			log.Fatalf("Error loading .env file: %s", err)
		}
	}
	databaseURL := os.Getenv("DATABASE_URL")
	DatabaseName := os.Getenv("DATABASE_NAME")

	return &Config{
		DatabaseURL:  databaseURL,
		DatabaseName: DatabaseName,
	}
}
