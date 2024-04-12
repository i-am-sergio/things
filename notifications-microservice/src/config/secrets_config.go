package config

import (
	"os"

	"github.com/joho/godotenv"
)

func LoadSecrets() (string, string, error) {
	err := godotenv.Load()

	port := os.Getenv("PORT")
	mongoURI := os.Getenv("MONGO_URI")

	return port, mongoURI, err
}
