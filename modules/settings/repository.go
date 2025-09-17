package settings

import (
	"context"
	"log"

	"gamesanct.com/lucky-seven/config/database"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var settingsCollection *mongo.Collection

func getSettingsCollection() *mongo.Collection {
	if settingsCollection == nil {
		if database.DB == nil {
			log.Println("database not initialized: call database.Connect before using settings repository")
			return nil
		}
		settingsCollection = database.DB.Collection("settings")
	}
	return settingsCollection
}

func Create(setting *Setting) error {
	coll := getSettingsCollection()
	if _, err := coll.InsertOne(context.Background(), setting); err != nil {
		return err
	}

	return nil
}

func GetSettingByID() (*Setting, error) {
	coll := getSettingsCollection()
	var setting Setting
	if err := coll.FindOne(context.Background(), bson.M{"setting_id": SETTING_ID}).Decode(&setting); err != nil {
		return nil, err
	}
	return &setting, nil
}

func UpdateFields(settingID string, fields bson.M) error {
	coll := getSettingsCollection()
	if _, err := coll.UpdateOne(context.Background(), bson.M{"setting_id": settingID}, bson.M{"$set": fields}); err != nil {
		return err
	}
	return nil
}
