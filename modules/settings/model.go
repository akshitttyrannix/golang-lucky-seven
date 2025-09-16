package settings

type Setting struct {
	SettingID    string `bson:"setting_id" json:"setting_id"`
	Status       string `bson:"status" json:"status"`
	BettingTimer uint16 `bson:"betting_timer" json:"betting_timer"`
	BonusTimer   uint16 `bson:"bonus_timer" json:"bonus_timer"`
	ResultTimer  uint16 `bson:"result_timer" json:"result_timer"`
	PauseTimer   uint16 `bson:"pause_timer" json:"pause_timer"`
	CreatedAt    uint32 `bson:"created_at" json:"created_at"`
	UpdatedAt    uint32 `bson:"updated_at" json:"updated_at"`
}
