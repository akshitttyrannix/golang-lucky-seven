package markets

import (
	"time"

	"gamesanct.com/lucky-seven/common/error"
	"gamesanct.com/lucky-seven/common/messages"
	"gamesanct.com/lucky-seven/common/success"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func CreateMarketEntity(dto CreateMarketDTO) *Market {
	return &Market{
		MarketID:   uuid.New().String(),
		MarketName: dto.MarketName,
		Status:     MARKET_STATUS_ACTIVE,
		Min:        dto.Min,
		Max:        dto.Max,
		MaxProfit:  dto.MaxProfit,
		CreatedAt:  uint32(time.Now().Unix()),
		UpdatedAt:  uint32(time.Now().Unix()),
	}
}

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
