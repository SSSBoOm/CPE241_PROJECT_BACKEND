package config

import (
	"os"

	"github.com/SSSBoOm/github.com/SSSBoOm/CPE241_Project_Backend/domain"
	"github.com/joho/godotenv"
)

func LoadEnv() *domain.Env {
	err := godotenv.Load(".env")

	if err != nil {
		panic(err)
	}

	GOOGLE_CLIENT_ID := os.Getenv("GOOGLE_CLIENT_ID")
	GOOGLE_CLIENT_SECRET := os.Getenv("GOOGLE_CLIENT_SECRET")
	GOOGLE_REDIRECT := os.Getenv("GOOGLE_REDIRECT")
	MYSQL_URI := os.Getenv("MYSQL_URI")

	return &domain.Env{
		GOOGLE_CLIENT_ID:     GOOGLE_CLIENT_ID,
		GOOGLE_CLIENT_SECRET: GOOGLE_CLIENT_SECRET,
		GOOGLE_REDIRECT:      GOOGLE_REDIRECT,
		MYSQL_URI:            MYSQL_URI,
	}
}
