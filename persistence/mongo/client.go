package mongo

import (
	"context"
	"log"

	"github.com/mongodb/mongo-go-driver/mongo"
)

// Connect is Shared connection
func Connect() *mongo.Database {
	client, err := mongo.NewClient("mongodb://admin:admin@127.0.0.1:27017")

	if err != nil {
		log.Fatal(err)
	}
	err = client.Connect(context.TODO())
	if err != nil {
		log.Fatal(err)
	}

	db := client.Database("demo")

	return db
}
