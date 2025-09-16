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
		log.Println("Error loading .env file")
		return "", "", ""
	}

	dbUrl := os.Getenv(common.DOT_ENV_URL)
	dbName := os.Getenv(common.DOT_ENV_DB_NAME)
	port := os.Getenv(common.DOT_ENV_PORT)
	nodeEnv := os.Getenv(common.DOT_ENV_NODE_ENV)

	if dbUrl == "" {
		log.Println(common.DOT_ENV_URL + " environment variable is not set")
		return "", "", ""
	}

	if dbName == "" {
		log.Println(common.DOT_ENV_DB_NAME + " environment variable is not set")
		return "", "", ""
	}

	if port == "" {
		log.Println(common.DOT_ENV_PORT + " environment variable is not set")
		return "", "", ""
	}

	if nodeEnv == "" {
		log.Println(common.DOT_ENV_NODE_ENV + " environment variable is not set")
		return "", "", ""
	}

	return dbUrl, dbName, port
}
