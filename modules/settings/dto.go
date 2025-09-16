package settings

type CreateSettingDTO struct {
	BettingTimer uint16 `json:"betting_timer" binding:"required,gte=0,lte=100"`
	BonusTimer   uint16 `json:"bonus_timer" binding:"required,gte=0,lte=100"`
	ResultTimer  uint16 `json:"result_timer" binding:"required,gte=0,lte=100"`
	PauseTimer   uint16 `json:"pause_timer" binding:"required,gte=0,lte=100"`
	Status       string `json:"status" binding:"required,oneof=active inactive"`
}
