package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	FRONTEND_URL         string
	BACKEND_PORT         string
	GOOGLE_CLIENT_ID     string
	GOOGLE_CLIENT_SECRET string
	GOOGLE_REDIRECT      string
	MYSQL_URI            string
}

func Load() (*Config, error) {
	err := godotenv.Load()

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	FRONTEND_URL := os.Getenv("FRONTEND_URL")
	BACKEND_PORT := os.Getenv("BACKEND_PORT")
	GOOGLE_CLIENT_ID := os.Getenv("GOOGLE_CLIENT_ID")
	GOOGLE_CLIENT_SECRET := os.Getenv("GOOGLE_CLIENT_SECRET")
	GOOGLE_REDIRECT := os.Getenv("GOOGLE_REDIRECT")
	MYSQL_URI := os.Getenv("MYSQL_URI")

	return &Config{
		FRONTEND_URL:         FRONTEND_URL,
		BACKEND_PORT:         BACKEND_PORT,
		GOOGLE_CLIENT_ID:     GOOGLE_CLIENT_ID,
		GOOGLE_CLIENT_SECRET: GOOGLE_CLIENT_SECRET,
		GOOGLE_REDIRECT:      GOOGLE_REDIRECT,
		MYSQL_URI:            MYSQL_URI,
	}, nil
}
