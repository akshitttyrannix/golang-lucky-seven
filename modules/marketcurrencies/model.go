package marketcurrencies

type MarketCurrency struct {
	MarketCurrencyID string  `bson:"market_currency_id" json:"market_currency_id"`
	MarketID         string  `bson:"market_id" json:"market_id"`
	MarketName       string  `bson:"market_name" json:"market_name"`
	CurrencyCode     string  `bson:"currency_code" json:"currency_code"`
	Min              float32 `bson:"min" json:"min"`
	Max              float32 `bson:"max" json:"max"`
	MaxProfit        float32 `bson:"max_profit" json:"max_profit"`
	Status           string  `bson:"status" json:"status"`
	BetLock          bool    `bson:"bet_lock" json:"bet_lock"`
	IsCurrencyUpdate bool    `bson:"is_currency_update" json:"is_currency_update"`
	Sequence         uint16  `bson:"sequence" json:"sequence"`
	CreatedAt        uint32  `bson:"created_at" json:"created_at"`
	UpdatedAt        uint32  `bson:"updated_at" json:"updated_at"`
	DeletedAt        uint32  `bson:"deleted_at" json:"deleted_at"`
	DeletedBy        string  `bson:"deleted_by" json:"deleted_by"`
	UpdatedBy        string  `bson:"updated_by" json:"updated_by"`
}
