package settings

type Setting struct {
	SettingID    string `json:"setting_id"`
	Status       string `json:"status"`
	BettingTimer int16  `json:"betting_timer"`
	BonusTimer   int16  `json:"bonus_timer"`
	ResultTimer  int16  `json:"result_timer"`
	PauseTimer   int16  `json:"pause_timer"`
	CreatedAt    string `json:"created_at"`
	UpdatedAt    string `json:"updated_at"`
}
