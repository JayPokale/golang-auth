package config

import (
	"context"
	"errors"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var db *mongo.Database

func GetCollection(val string) *mongo.Collection {
	return db.Collection(val)
}

func ConnectDB() error {
	uri := os.Getenv("MONGODB_URI")
	if uri == "" {
		return errors.New("MONGODB_URI is not set in environment")
	}

	// Connect to mongodb
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		return errors.New("Database connection failed")
	}

	db = client.Database("GOLANG")

	return nil
}

func CloseDB() error {
	return db.Client().Disconnect(context.TODO())
}
