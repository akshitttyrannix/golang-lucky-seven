package runners

import "go.mongodb.org/mongo-driver/bson/primitive"

type Runner struct {
	RunnerID          string               `bson:"runner_id" json:"runner_id"`
	RunnerName        string               `bson:"runner_name" json:"runner_name"`
	DisplayRunnerName string               `bson:"display_runner_name" json:"display_runner_name"`
	MarketID          string               `bson:"market_id" json:"market_id"`
	MarketName        string               `bson:"market_name" json:"market_name"`
	Odds              primitive.Decimal128 `bson:"odds" json:"odds"`
	Sequence          uint16               `bson:"sequence" json:"sequence"`
	Status            string               `bson:"status" json:"status"`
	IsDeleted         bool                 `bson:"is_deleted" json:"is_deleted"`
	CreatedAt         uint32               `bson:"created_at" json:"created_at"`
	UpdatedAt         uint32               `bson:"updated_at" json:"updated_at"`
	DeletedAt         uint32               `bson:"deleted_at" json:"deleted_at"`
	DeletedBy         string               `bson:"deleted_by" json:"deleted_by"`
	UpdatedBy         string               `bson:"updated_by" json:"updated_by"`
}
