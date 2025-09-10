package settings

type Setting struct {
	SettingID    string `bson:"setting_id" json:"setting_id"`
	Status       string `bson:"status" json:"status"`
	BettingTimer int16  `bson:"betting_timer" json:"betting_timer"`
	BonusTimer   int16  `bson:"bonus_timer" json:"bonus_timer"`
	ResultTimer  int16  `bson:"result_timer" json:"result_timer"`
	PauseTimer   int16  `bson:"pause_timer" json:"pause_timer"`
	CreatedAt    int64  `bson:"created_at" json:"created_at"`
	UpdatedAt    int64  `bson:"updated_at" json:"updated_at"`
}
