package markets

import "github.com/gin-gonic/gin"

func Routes(router *gin.RouterGroup) {

	g := router.Group("/markets")

	g.POST("", CreateMarket)
}
