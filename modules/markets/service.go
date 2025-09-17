package markets

import (
	"time"

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
