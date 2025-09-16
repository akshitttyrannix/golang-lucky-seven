package markets

type Market struct {
	MarketID         string `bson:"market_id" json:"market_id"`
	MarketName       string `bson:"market_name" json:"market_name"`
	Status           string `bson:"status" json:"status"`
	MarketType       string `bson:"market_type" json:"market_type"`
	Min              int    `bson:"min" json:"min"`
	Max              int    `bson:"max" json:"max"`
	MaxProfit        int    `bson:"max_profit" json:"max_profit"`
	BetLock          bool   `bson:"bet_lock" json:"bet_lock"`
	IsGameBetLock    bool   `bson:"is_game_bet_lock" json:"is_game_bet_lock"`
	IsPackageBetLock bool   `bson:"is_package_bet_lock" json:"is_package_bet_lock"`
	GameStatus       string `bson:"game_status" json:"game_status"`
	PackageStatus    string `bson:"package_status" json:"package_status"`
	BonusOdds        []int  `bson:"bonus_odds" json:"bonus_odds"`
	BonusStatus      string `bson:"bonus_status" json:"bonus_status"`
	Sequence         int    `bson:"sequence" json:"sequence"`
	MaxBonusRunners  int    `bson:"max_bonus_runners" json:"max_bonus_runners"`
	IsDeleted        bool   `bson:"is_deleted" json:"is_deleted"`
	CreatedAt        int64  `bson:"created_at" json:"created_at"`
	UpdatedAt        int64  `bson:"updated_at" json:"updated_at"`
	DeletedAt        int64  `bson:"deleted_at" json:"deleted_at"`
	DeletedBy        string `bson:"deleted_by" json:"deleted_by"`
	UpdatedBy        string `bson:"updated_by" json:"updated_by"`
}
