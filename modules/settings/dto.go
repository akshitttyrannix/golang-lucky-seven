package settings

type CreateSettingDTO struct {
	BettingTimer int16  `json:"betting_timer" binding:"required,gte=0,lte=100"`
	BonusTimer   int16  `json:"bonus_timer" binding:"required,gte=0,lte=100"`
	ResultTimer  int16  `json:"result_timer" binding:"required,gte=0,lte=100"`
	PauseTimer   int16  `json:"pause_timer" binding:"required,gte=0,lte=100"`
	Status       string `json:"status" binding:"required,oneof=active inactive"`
}
