package rounds

import (
	"context"
	"log"

	"gamesanct.com/lucky-seven/common/functions"
	"gamesanct.com/lucky-seven/config/database"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var roundsCollection *mongo.Collection

func getRoundsCollection() *mongo.Collection {
	if roundsCollection == nil {
		if database.DB == nil {
			log.Fatal("database not initialized: call database.Connect before using rounds repository")
		}
		roundsCollection = database.DB.Collection("rounds")
	}
	return roundsCollection
}

func Create(round *Round) error {
	coll := getRoundsCollection()
	if _, err := coll.InsertOne(context.Background(), round); err != nil {
		return err
	}

	return nil

}

func Update(round *Round) error {
	coll := getRoundsCollection()
	if _, err := coll.UpdateOne(context.Background(), bson.M{"round_id": round.RoundID}, bson.M{"$set": round}); err != nil {
		return err
	}

	return nil
}

func GetTodaysRoundCount() int64 {
	coll := getRoundsCollection()

	filter := bson.M{
		"created_at": bson.M{"$gte": functions.GetStartOfDay(), "$lte": functions.GetEndOfDay()},
	}

	projection := bson.M{"display_id": 1}
	opts := options.FindOne().SetProjection(projection).SetSort(bson.M{"created_at": -1})

	var latestRound Round
	err := coll.FindOne(context.Background(), filter, opts).Decode(&latestRound)
	if err != nil {
		return 0
	}

	// Extract the last 5 characters from display_id and convert to int64
	displayIDSuffix := latestRound.DisplayID[len(latestRound.DisplayID)-5:]
	count := int64(0)

	// Parse the suffix to get the count
	for _, char := range displayIDSuffix {
		if char >= '0' && char <= '9' {
			count = count*10 + int64(char-'0')
		}
	}

	return count
}
