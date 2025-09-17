package marketcurrencies

import (
	"context"
	"log"

	"gamesanct.com/lucky-seven/config/database"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var marketCurrenciesCollection *mongo.Collection

func getMarketCurrenciesCollection() *mongo.Collection {
	if marketCurrenciesCollection == nil {
		if database.DB == nil {
			log.Println("database not initialized: call database.Connect before using market currencies repository")
			return nil
		}
		marketCurrenciesCollection = database.DB.Collection("market_currencies")
	}
	return marketCurrenciesCollection
}

func Create(round *MarketCurrency) error {
	coll := getMarketCurrenciesCollection()
	if _, err := coll.InsertOne(context.Background(), round); err != nil {
		return err
	}

	return nil
}

func Count() (uint16, error) {
	coll := getMarketCurrenciesCollection()

	count, err := coll.CountDocuments(context.Background(), bson.M{"is_deleted": false})
	if err != nil {
		return 0, err
	}

	return uint16(count), nil
}
