package markets

import "github.com/gin-gonic/gin"

func Routes(router *gin.RouterGroup) {
	router.POST("/markets", CreateMarket)
}
