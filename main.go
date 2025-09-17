package main

import (
	"gamesanct.com/lucky-seven/config/database"
	"gamesanct.com/lucky-seven/config/env"
	"gamesanct.com/lucky-seven/modules/games"
	"gamesanct.com/lucky-seven/modules/markets"
	"gamesanct.com/lucky-seven/modules/settings"
	"github.com/gin-gonic/gin"
)

func main() {
	uri, dbName, port := env.GetEnv()
	database.Connect(uri, dbName)

	server := gin.Default()

	games.InitGame()

	prefix := server.Group("/api/v1")
	settings.Routes(prefix)
	markets.Routes(prefix)

	server.Run(":" + port)
}
