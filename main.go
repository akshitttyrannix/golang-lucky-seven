package main

import (
	"gamesanct.com/lucky-seven/config/database"
	"gamesanct.com/lucky-seven/config/env"
	"gamesanct.com/lucky-seven/modules/settings"
	"github.com/gin-gonic/gin"
)

func main() {
	uri, dbName, port := env.GetEnv()
	database.Connect(uri, dbName)

	server := gin.Default()

	prefix := server.Group("/api/v1")
	settings.Routes(prefix)

	server.Run(":" + port)
}
