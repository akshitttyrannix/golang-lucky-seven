package settings

import "github.com/gin-gonic/gin"

func Routes(router *gin.RouterGroup) {

	g := router.Group("/settings")

	g.POST("", CreateSetting)
	g.PUT("/start", StartGame)
	g.PUT("/stop", StopGame)
}
