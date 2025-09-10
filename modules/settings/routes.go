package settings

import "github.com/gin-gonic/gin"

func Routes(router *gin.RouterGroup) {
	router.POST("/settings", CreateSetting)
}
