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
	now := uint32(time.Now().Unix())
	return &Market{
		MarketID:   uuid.NewString(),
		MarketName: dto.MarketName,
		Status:     MARKET_STATUS_ACTIVE,
		Min:        dto.Min,
		Max:        dto.Max,
		MaxProfit:  dto.MaxProfit,
		Sequence:   count + 1,
		CreatedAt:  now,
		UpdatedAt:  now,
	}
}

func CreateRunnersEntity(marketID string, dto CreateMarketDTO) []runners.Runner {
	baseCount, err := runners.Count()
	if err != nil {
		log.Println("Error counting runners:", err)
		return nil
	}

	now := uint32(time.Now().Unix())
	runnerEntities := make([]runners.Runner, 0, len(dto.Runners))

	for i, runnerInput := range dto.Runners {
		odds, err := helper.StringToDecimal128(runnerInput.Odds)
		if err != nil {
			log.Println("Error converting odds to decimal128:", err)
			return nil
		}

		runnerEntities = append(runnerEntities, runners.Runner{
			RunnerID:          uuid.NewString(),
			RunnerName:        runnerInput.RunnerName,
			DisplayRunnerName: runnerInput.DisplayRunnerName,
			Odds:              odds,
			MarketID:          marketID,
			MarketName:        dto.MarketName,
			Sequence:          baseCount + uint16(i) + 1,
			Status:            runners.RUNNER_STATUS_ACTIVE,
			CreatedAt:         now,
			UpdatedAt:         now,
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

	now := uint32(time.Now().Unix())
	return &marketcurrencies.MarketCurrency{
		MarketCurrencyID: uuid.NewString(),
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
		CreatedAt:        now,
		UpdatedAt:        now,
	}
}
