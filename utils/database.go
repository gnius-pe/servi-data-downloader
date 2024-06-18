package utils

import (
	"context"
	"fmt"

	"github.com/gnius-pe/servi-data-downloader/configs"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

func ConnectDB() {
	fmt.Println("Entroooo DB")
	var err error
	client, err = mongo.Connect(context.TODO(), options.Client().ApplyURI(configs.MongoURI))
	if err != nil {
		fmt.Println("err", err)
		panic(err)

	}
	fmt.Println("Acabo DB")
}

func GetCollection(database, collection string) *mongo.Collection {
	return client.Database(database).Collection(collection)
}
