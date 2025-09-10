package settings

import (
	"context"
	"log"

	"gamesanct.com/lucky-seven/config/database"
	"go.mongodb.org/mongo-driver/mongo"
)

var settingsCollection *mongo.Collection

func getSettingsCollection() *mongo.Collection {
	if settingsCollection == nil {
		if database.DB == nil {
			log.Fatal("database not initialized: call database.Connect before using settings repository")
		}
		settingsCollection = database.DB.Collection("settings")
	}
	return settingsCollection
}

func create(setting *Setting) error {
	coll := getSettingsCollection()
	if _, err := coll.InsertOne(context.Background(), setting); err != nil {
		return err
	}

	return nil
}
