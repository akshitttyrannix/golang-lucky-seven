package main

import (
	"gamesanct.com/lucky-seven/config/database"
	"gamesanct.com/lucky-seven/config/env"
	"github.com/gin-gonic/gin"
)

func main() {
	uri, dbName, port := env.GetEnv()
	database.Connect(uri, dbName)

	server := gin.Default()

	server.Run(":" + port)
}
