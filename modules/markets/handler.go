package markets

import (
	"log"

	"gamesanct.com/lucky-seven/common/customerror"
	"gamesanct.com/lucky-seven/common/messages"
	"gamesanct.com/lucky-seven/common/success"
	"gamesanct.com/lucky-seven/modules/marketcurrencies"
	"gamesanct.com/lucky-seven/modules/runners"
	"github.com/gin-gonic/gin"
	"golang.org/x/sync/errgroup"
)

func CreateMarket(ctx *gin.Context) {
	var dto CreateMarketDTO
	if err := ctx.ShouldBindJSON(&dto); err != nil {
		customerror.BadRequest(ctx, err)
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

	var g errgroup.Group

	g.Go(func() error {
		return Create(marketEntity)
	})

	g.Go(func() error {
		return runners.CreateMany(runnersEntity)
	})

	g.Go(func() error {
		return marketcurrencies.Create(marketCurrencyEntity)
	})

	if err := g.Wait(); err != nil {
		customerror.SomethingWentWrong(ctx, err)
		return
	}

	success.Success(ctx, messages.MARKET_CREATED, marketEntity)
}
