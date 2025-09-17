package markets

import (
	"gamesanct.com/lucky-seven/common/error"
	"gamesanct.com/lucky-seven/common/messages"
	"gamesanct.com/lucky-seven/common/success"
	"github.com/gin-gonic/gin"
)

func CreateMarket(ctx *gin.Context) {
	var dto CreateMarketDTO
	if err := ctx.ShouldBindJSON(&dto); err != nil {
		error.BadRequest(ctx, err)
		return
	}

	market := CreateMarketEntity(dto)

	if err := Create(market); err != nil {
		error.SomethingWentWrong(ctx, err)
		return
	}

	success.Success(ctx, messages.MARKET_CREATED, market)
}
