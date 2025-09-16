package markets

type RunnerGet struct {
	RunnerName        string `json:"runner_name" binding:"required"`
	DisplayRunnerName string `json:"display_runner_name" binding:"required"`
	Odds              uint16 `json:"odds" binding:"required"`
}

type CreateMarketDTO struct {
	MarketName string      `json:"market_name" binding:"required"`
	Runners    []RunnerGet `json:"runners" binding:"required"`
	Min        uint16      `json:"min" binding:"required"`
	Max        uint16      `json:"max" binding:"required"`
	MaxProfit  uint16      `json:"max_profit" binding:"required"`
}
