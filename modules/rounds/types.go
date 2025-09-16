package rounds

type Round struct {
	RoundID       string `bson:"round_id" json:"round_id"`
	DisplayID     string `bson:"display_id" json:"display_id"`
	GameID        string `bson:"game_id" json:"game_id"`
	GameName      string `bson:"game_name" json:"game_name"`
	Result        any    `bson:"result" json:"result"`
	State         string `bson:"state" json:"state"`
	StartTime     int64  `bson:"start_time" json:"start_time"`
	EndTime       int64  `bson:"end_time" json:"end_time"`
	Odds          any    `bson:"odds" json:"odds"`
	HasBet        bool   `bson:"has_bet" json:"has_bet"`
	BetPlaceCount int16  `bson:"bet_place_count" json:"bet_place_count"`
	BetAmount     int16  `bson:"bet_amount" json:"bet_amount"`
	Profit        int16  `bson:"profit" json:"profit"`
	CreatedAt     int64  `bson:"created_at" json:"created_at"`
	UpdatedBy     string `bson:"updated_by" json:"updated_by"`
	UpdatedAt     int64  `bson:"updated_at" json:"updated_at"`
	DeletedAt     int64  `bson:"deleted_at" json:"deleted_at"`
	DeletedBy     string `bson:"deleted_by" json:"deleted_by"`
	IsDeleted     bool   `bson:"is_deleted" json:"is_deleted"`
	Markets       any    `bson:"markets" json:"markets"`
}
