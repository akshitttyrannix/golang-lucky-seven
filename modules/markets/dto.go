package markets

type RunnerGet struct {
	RunnerName        string  `json:"runner_name" binding:"required"`
	DisplayRunnerName string  `json:"display_runner_name" binding:"required"`
	Odds              float32 `json:"odds" binding:"required"`
}

type CreateMarketDTO struct {
	MarketName string      `json:"market_name" binding:"required"`
	Runners    []RunnerGet `json:"runners" binding:"required"`
	Min        float32     `json:"min" binding:"required"`
	Max        float32     `json:"max" binding:"required"`
	MaxProfit  float32     `json:"max_profit" binding:"required"`
}
