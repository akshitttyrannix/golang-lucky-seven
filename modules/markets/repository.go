package markets

import (
	"context"
	"log"

	"gamesanct.com/lucky-seven/config/database"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var marketsCollection *mongo.Collection

func getMarketsCollection() *mongo.Collection {
	if marketsCollection == nil {
		if database.DB == nil {
			log.Println("database not initialized: call database.Connect before using markets repository")
			return nil
		}
		marketsCollection = database.DB.Collection("markets")
	}
	return marketsCollection
}

func Create(round *Market) error {
	coll := getMarketsCollection()
	if _, err := coll.InsertOne(context.Background(), round); err != nil {
		return err
	}

	return nil
}

func Count() (uint16, error) {
	coll := getMarketsCollection()

	count, err := coll.CountDocuments(context.Background(), bson.M{"is_deleted": false})
	if err != nil {
		return 0, err
	}

	return uint16(count), nil
}
