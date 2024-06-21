package utils

import (
	"context"
	"fmt"

	"github.com/gnius-pe/servi-data-downloader/configs"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

func ConnectDB() error {
	var err error
	client, err = mongo.Connect(context.TODO(), options.Client().ApplyURI(configs.MongoURI))
	if err != nil {
		return fmt.Errorf("error connecting to MongoDB: %v", err)
	}

	// Ping MongoDB to ensure connection is established
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		return fmt.Errorf("error pinging MongoDB: %v", err)
	}

	fmt.Println("Connected to MongoDB!")
	return nil
}

func GetCollection(database, collection string) *mongo.Collection {
	if client == nil {
		// Manejar el caso donde client es nil (por ejemplo, intentar reconectar o devolver un error)
		panic("MongoDB client is nil")
	}
	return client.Database(database).Collection(collection)
}
