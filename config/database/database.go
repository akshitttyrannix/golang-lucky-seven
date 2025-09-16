package database

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DB *mongo.Database

func Connect(uri string, dbName string) {
	clientOptions := options.Client().ApplyURI(uri)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Println("Failed to create MongoDB client:", err)
		return
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Println("Failed to ping MongoDB:", err)
		return
	}

	DB = client.Database(dbName)
	log.Println("✅ Connected to MongoDB")
}
