package repositories

import (
	"context"
	"fmt"
	"sync"

	"github.com/gnius-pe/servi-data-downloader/configs"
	"github.com/gnius-pe/servi-data-downloader/models"
	"github.com/gnius-pe/servi-data-downloader/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	collection *mongo.Collection
	once       sync.Once
)

func GetCollection() *mongo.Collection {
	once.Do(func() {
		collection = utils.GetCollection(configs.DatabaseName, configs.CollectionNameP)
		fmt.Println("entroooooo")
	})
	return collection
}

func GetPatientByID(id string) (*models.Patient, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	var patient models.Patient
	err = GetCollection().FindOne(context.TODO(), bson.M{"_id": objectID}).Decode(&patient)
	if err != nil {
		return nil, err
	}
	return &patient, nil
}
