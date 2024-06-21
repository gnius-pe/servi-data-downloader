package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	ServerPort,
	MongoURI,
	DatabaseName,
	CollectionNameP string
)

func LoadConfig() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	ServerPort = os.Getenv("SERVER_PORT")
	MongoURI = os.Getenv("MONGO_URI")
	DatabaseName = os.Getenv("DATABASE_NAME")
	CollectionNameP = os.Getenv("COLLECTION_NAME_P")
}
