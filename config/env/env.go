package env

import (
	"log"
	"os"

	"gamesanct.com/lucky-seven/common"
	"github.com/joho/godotenv"
)

func GetEnv() (string, string, string) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dbUrl := os.Getenv(common.DOT_ENV_URL)
	dbName := os.Getenv(common.DOT_ENV_DB_NAME)
	port := os.Getenv(common.DOT_ENV_PORT)
	nodeEnv := os.Getenv(common.DOT_ENV_NODE_ENV)

	if dbUrl == "" {
		log.Fatal(common.DOT_ENV_URL + " environment variable is not set")
	}

	if dbName == "" {
		log.Fatal(common.DOT_ENV_DB_NAME + " environment variable is not set")
	}

	if port == "" {
		log.Fatal(common.DOT_ENV_PORT + " environment variable is not set")
	}

	if nodeEnv == "" {
		log.Fatal(common.DOT_ENV_NODE_ENV + " environment variable is not set")
	}

	return dbUrl, dbName, port
}
