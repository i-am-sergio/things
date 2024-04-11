package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadSecrets() (string, string) {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	port := os.Getenv("PORT")
	mongoURI := os.Getenv("MONGO_URI")

	return port, mongoURI
}
