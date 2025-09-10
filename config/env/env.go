package env

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func GetEnv() (string, string, string) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	uri := os.Getenv("DATABASE_URL")
	dbName := os.Getenv("DATABASE_NAME")
	port := os.Getenv("PORT")

	if uri == "" || dbName == "" || port == "" {
		log.Fatal("DATABASE_URL, DATABASE_NAME or PORT environment variable is not set")
	}

	return uri, dbName, port
}
