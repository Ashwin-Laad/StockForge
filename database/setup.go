package database

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

// MongoClient object, which creates the connection to and performs operations on the MongoDB deployment.
func ConnectToDb() *mongo.Client {

	URI := "mongodb://localhost:27017"
	//setting up options for connection to our database


	client, err := mongo.Connect(options.Client().ApplyURI(URI))
	if err != nil {
		return nil
	}

	//testing if the database is working or not
	error:=client.Ping(context.TODO(),nil)
	if error!=nil{
		fmt.Println("DB problem")
	}
	fmt.Println("DB worked")
	return client
}

//create a variable to export it to our other functions
//client so it connects to MongoDb
var Client *mongo.Client=ConnectToDb()

// type Collection struct {
//     client         *Client
//     db             *Database
//     name           string
//     readConcern    *readconcern.ReadConcern
//     writeConcern   *writeconcern.WriteConcern
//     readPreference *readpref.ReadPref
//     readSelector   description.ServerSelector
//     writeSelector  description.ServerSelector
//     bsonOpts       *options.BSONOptions
//     registry       *bson.Registry
// }
//user data function
//collection imported from the mongo package

func GetCollection(client *mongo.Client, collectionName string)(*mongo.Collection){
var collection *mongo.Collection=client.Database("Inventory").Collection(collectionName)
return collection
}