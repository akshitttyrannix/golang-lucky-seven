package runners

import (
	"context"
	"log"

	"gamesanct.com/lucky-seven/config/database"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var runnersCollection *mongo.Collection

func getRunnersCollection() *mongo.Collection {
	if runnersCollection == nil {
		if database.DB == nil {
			log.Println("database not initialized: call database.Connect before using runners repository")
			return nil
		}
		runnersCollection = database.DB.Collection("runners")
	}
	return runnersCollection
}

func Create(round *Runner) error {
	coll := getRunnersCollection()
	if _, err := coll.InsertOne(context.Background(), round); err != nil {
		return err
	}

	return nil
}

func CreateMany(runners []Runner) error {
	coll := getRunnersCollection()

	docs := make([]any, len(runners))
	for i := range runners {
		docs[i] = runners[i]
	}

	if _, err := coll.InsertMany(context.Background(), docs); err != nil {
		return err
	}

	return nil
}

func Count() (uint16, error) {
	coll := getRunnersCollection()

	count, err := coll.CountDocuments(context.Background(), bson.M{"is_deleted": false})
	if err != nil {
		return 0, err
	}

	return uint16(count), nil
}
