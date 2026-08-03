package database

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// ConnectToDb establishes a connection to the MongoDB server
// and returns a pointer to the connected MongoDB client.
//
// Purpose:
// --------
// Every application that wants to communicate with MongoDB
// must first create a client. Think of the client as the
// application's gateway to the database.
//
// Once connected, this client can:
//
//   - Open databases
//   - Access collections
//   - Insert documents
//   - Read documents
//   - Update documents
//   - Delete documents
//
// Returns:
//
//	*mongo.Client
//	    A pointer to the connected MongoDB client.
//
//	    If the connection fails, nil is returned.
//
// Why return a pointer?
// ---------------------
// mongo.Client is a large object responsible for maintaining
// connection pools and communicating with MongoDB.
//
// Returning a pointer means:
//
//   - The object is not copied.
//   - Every package shares the same client.
//   - Database connections remain reusable throughout
//     the application.
func ConnectToDb() *mongo.Client {

	// uri stores the address of the MongoDB server.
	//
	// mongodb://
	//     Specifies that we are using the MongoDB protocol.
	//
	// localhost
	//     Means MongoDB is running on this computer.
	//
	// 27017
	//     Default MongoDB port.
	//
	// If MongoDB were hosted online, this URI would instead
	// point to MongoDB Atlas or another remote server.

	// uri := "mongodb://localhost:27017"

	uri := "mongodb://mongodb:27017"

	// ApplyURI creates a ClientOptions object.
	//
	// ClientOptions contains configuration describing
	// how the client should connect.
	//
	// Here we specify only the database address,
	// although many other options (authentication,
	// TLS, timeouts, etc.) can also be configured.
	client, err := mongo.Connect(
		context.TODO(),
		options.Client().ApplyURI(uri),
	)

	// If Connect() encounters an error while creating
	// the client, return nil to indicate failure.
	if err != nil {
		return nil
	}

	// Ping checks whether the application can actually
	// communicate with MongoDB.
	//
	// Why call Ping?
	// --------------
	// Creating the client object alone does not guarantee
	// the database is reachable.
	//
	// Ping sends a small request to MongoDB and waits
	// for a response.
	//
	// context.TODO() supplies a Context object.
	//
	// Context carries request-scoped information such as:
	//
	//   - cancellation signals
	//   - deadlines
	//   - timeouts
	//
	// TODO() is commonly used when no specific context
	// has been created yet.
	//
	// nil means MongoDB may choose any available server
	// to respond to the ping.
	pingErr := client.Ping(context.TODO(), nil)

	if pingErr != nil {
		fmt.Println("DB problem")
	}

	fmt.Println("DB worked")

	// Return the connected MongoDB client so other parts
	// of the application can use it.
	return client
}

// Client is a global MongoDB client shared throughout
// the application.
//
// Since it begins with a capital letter, it is exported,
// meaning other packages can access it as:
//
// database.Client
//
// ConnectToDb() is executed once while this package is
// being initialized, so the client is ready before main()
// begins executing.
var Client *mongo.Client = ConnectToDb()

// Collection represents one MongoDB collection.
//
// (The driver already defines this type, so this commented
// section is only for reference.)
//
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

// GetCollection returns a pointer to a MongoDB collection.
//
// Parameters:
// -----------
//
// client
//
//	The connected MongoDB client.
//
// collectionName
//
//	The name of the collection to access.
//
// Returns:
// --------
//
// *mongo.Collection
//
//	A pointer representing the requested collection.
//
// This function does NOT create the collection.
// It simply returns a handle that lets us perform
// operations such as:
//
//   - InsertOne()
//   - Find()
//   - FindOne()
//   - UpdateOne()
//   - DeleteOne()
func GetCollection(client *mongo.Client, collectionName string) *mongo.Collection {

	// Database("Inventory")
	// selects the Inventory database.
	//
	// Collection(collectionName)
	// selects one collection inside that database.
	//
	// Example:
	//
	// collectionName = "Users"
	//
	// Inventory
	//     └── Users
	//
	// The returned object is simply a handle used to
	// perform database operations.
	var collection *mongo.Collection = client.Database("Inventory").Collection(collectionName)

	return collection
}
