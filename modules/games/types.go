package games

const (
	GAME_STATE_INIT    = "init"
	GAME_STATE_BETTING = "betting"
	GAME_STATE_BONUS   = "bonus"
	GAME_STATE_RESULT  = "result"
	GAME_STATE_PAUSE   = "pause"
	GAME_STATE_END     = "end"
)

type ResultRunner struct {
	RunnerName string `json:"runner_name"`
	IsWin      bool   `json:"is_win"`
}

type ResultMarket struct {
	MarketName string         `json:"market_name"`
	Runners    []ResultRunner `json:"runners"`
}

type Result struct {
	Card   string         `json:"card"`
	Winner []ResultMarket `json:"winner,omitempty"`
}
