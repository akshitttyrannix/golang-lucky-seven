package markets

import (
	"log"
	"time"

	"gamesanct.com/lucky-seven/common/helper"
	"gamesanct.com/lucky-seven/modules/marketcurrencies"
	"gamesanct.com/lucky-seven/modules/runners"
	"github.com/google/uuid"
)

func CreateMarketEntity(dto CreateMarketDTO, count uint16) *Market {
	return &Market{
		MarketID:   uuid.New().String(),
		MarketName: dto.MarketName,
		Status:     MARKET_STATUS_ACTIVE,
		Min:        dto.Min,
		Max:        dto.Max,
		MaxProfit:  dto.MaxProfit,
		Sequence:   count + 1,
		CreatedAt:  uint32(time.Now().Unix()),
		UpdatedAt:  uint32(time.Now().Unix()),
	}
}

func CreateRunnersEntity(marketID string, dto CreateMarketDTO) []runners.Runner {
	count, err := Count()
	if err != nil {
		log.Println("Error counting runners:", err)
		return nil
	}

	runnerEntities := []runners.Runner{}
	for _, runnerInput := range dto.Runners {

		odds, err := helper.StringToDecimal128(runnerInput.Odds)
		if err != nil {
			log.Println("Error converting odds to decimal128:", err)
			return nil
		}

		runnerEntities = append(runnerEntities, runners.Runner{
			RunnerID:          uuid.New().String(),
			RunnerName:        runnerInput.RunnerName,
			DisplayRunnerName: runnerInput.DisplayRunnerName,
			Odds:              odds,
			MarketID:          marketID,
			MarketName:        dto.MarketName,
			Sequence:          count + 1,
			Status:            runners.RUNNER_STATUS_ACTIVE,
			CreatedAt:         uint32(time.Now().Unix()),
			UpdatedAt:         uint32(time.Now().Unix()),
		})
	}

	return runnerEntities
}

func CreateMarketCurrencyEntity(marketID string, dto CreateMarketDTO) *marketcurrencies.MarketCurrency {
	count, err := marketcurrencies.Count()
	if err != nil {
		log.Println("Error counting market currencies:", err)
		return nil
	}

	return &marketcurrencies.MarketCurrency{
		MarketCurrencyID: uuid.New().String(),
		MarketID:         marketID,
		MarketName:       dto.MarketName,
		CurrencyCode:     marketcurrencies.CURRENCY_CODE_DEFAULT,
		Min:              dto.Min,
		Max:              dto.Max,
		MaxProfit:        dto.MaxProfit,
		Status:           marketcurrencies.MARKET_CURRENCY_STATUS_ACTIVE,
		BetLock:          false,
		IsCurrencyUpdate: false,
		Sequence:         count + 1,
		CreatedAt:        uint32(time.Now().Unix()),
		UpdatedAt:        uint32(time.Now().Unix()),
	}
}
