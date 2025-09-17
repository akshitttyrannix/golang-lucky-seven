package markets

import (
	"log"

	"gamesanct.com/lucky-seven/common/error"
	"gamesanct.com/lucky-seven/common/messages"
	"gamesanct.com/lucky-seven/common/success"
	"gamesanct.com/lucky-seven/modules/marketcurrencies"
	"gamesanct.com/lucky-seven/modules/runners"
	"github.com/gin-gonic/gin"
)

func CreateMarket(ctx *gin.Context) {
	var dto CreateMarketDTO
	if err := ctx.ShouldBindJSON(&dto); err != nil {
		error.BadRequest(ctx, err)
		return
	}

	count, err := Count()
	if err != nil {
		log.Println("Error counting markets:", err)
		return
	}

	marketEntity := CreateMarketEntity(dto, count)

	runnersEntity := CreateRunnersEntity(marketEntity.MarketID, dto)

	marketCurrencyEntity := CreateMarketCurrencyEntity(marketEntity.MarketID, dto)

	if err := Create(marketEntity); err != nil {
		error.SomethingWentWrong(ctx, err)
		return
	}

	if err := runners.CreateMany(runnersEntity); err != nil {
		error.SomethingWentWrong(ctx, err)
		return
	}

	if err := marketcurrencies.Create(marketCurrencyEntity); err != nil {
		error.SomethingWentWrong(ctx, err)
		return
	}

	success.Success(ctx, messages.MARKET_CREATED, marketEntity)
}
