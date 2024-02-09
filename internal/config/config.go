package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	GOOGLE_CLIENT_ID     string
	GOOGLE_CLIENT_SECRET string
	GOOGLE_REDIRECT      string
	MYSQL_URI            string
}

func Load() (*Config, error) {
	err := godotenv.Load(".env")

	if err != nil {
		panic(err)
	}

	GOOGLE_CLIENT_ID := os.Getenv("GOOGLE_CLIENT_ID")
	GOOGLE_CLIENT_SECRET := os.Getenv("GOOGLE_CLIENT_SECRET")
	GOOGLE_REDIRECT := os.Getenv("GOOGLE_REDIRECT")
	MYSQL_URI := os.Getenv("MYSQL_URI")

	return &Config{
		GOOGLE_CLIENT_ID:     GOOGLE_CLIENT_ID,
		GOOGLE_CLIENT_SECRET: GOOGLE_CLIENT_SECRET,
		GOOGLE_REDIRECT:      GOOGLE_REDIRECT,
		MYSQL_URI:            MYSQL_URI,
	}, nil
}
