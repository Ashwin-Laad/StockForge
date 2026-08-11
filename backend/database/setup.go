package database

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectToDb() *mongo.Client {

	// uri stores the address of the MongoDB server.

	// uri := "mongodb://localhost:27017"

	uri := "mongodb://mongodb:27017"

	client, err := mongo.Connect(
		context.TODO(),
		options.Client().ApplyURI(uri),
	)

	// If Connect() encounters an error while creating
	// the client, return nil to indicate failure.
	if err != nil {
		return nil
	}

	pingErr := client.Ping(context.TODO(), nil)

	if pingErr != nil {
		fmt.Println("DB problem")
	}

	fmt.Println("DB worked")

	return client
}

// Client is a global MongoDB client shared throughout
// the application.
//once client is called auto creates the connection

var Client *mongo.Client = ConnectToDb()

func GetCollection(client *mongo.Client, collectionName string) *mongo.Collection {

	var collection *mongo.Collection = client.Database("Inventory").Collection(collectionName)

	return collection
}
