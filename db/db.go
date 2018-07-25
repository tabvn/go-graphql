package db

import (
	"github.com/mongodb/mongo-go-driver/mongo"
	"log"
	"context"
	"go-graphql/config"
)

var database *mongo.Database = nil

func Connect() (*mongo.Database, error) {

	if database != nil {
		return database, nil
	}

	client, err := mongo.NewClient(config.DatabaseURL)
	if err != nil {
		log.Fatal(err)
	}
	err = client.Connect(context.TODO())
	if err != nil {
		log.Fatal(err)

		return nil, err
	}

	database = client.Database(config.DatabaseName)

	return database, nil

}

func Collection(name string) (*mongo.Collection) {
	if database == nil {
		Connect()
	}
	
	return database.Collection(name)
}
