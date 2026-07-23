package controllers

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

// Application stores resources that every controller needs.
//
// Why create this struct?
// -----------------------
// Instead of creating or fetching the database collections inside every
// controller function, we store them once in an Application object.
//
// This makes the code:
//
//   - easier to organize
//   - easier to test
//   - easier to extend later
//
// Every method attached to Application can directly access these
// collections using:
//
//	app.userCollection
//	app.prodCollection
type Application struct {

	// userCollection is a handle to the MongoDB User collection.
	//
	// A collection represents one group of documents
	// inside the database.
	userCollection *mongo.Collection

	// prodCollection points to the Products collection.
	prodCollection *mongo.Collection
}

// CreateApplication creates and initializes an Application object.
//
// Parameters:
//
//	userCollection
//	    MongoDB User collection.
//
//	prodCollection
//	    MongoDB Product collection.
//
// Returns:
//
//	*Application
//
// Returning a pointer avoids copying the struct and lets every
// handler use the same Application instance.
func CreateApplication(
	userCollection *mongo.Collection,
	prodCollection *mongo.Collection,
) *Application {

	return &Application{
		userCollection: userCollection,
		prodCollection: prodCollection,
	}
}

// AddProductToCart returns a Gin handler that adds one product
// to a user's shopping cart.
//
// Route Example:
//
//	POST /user/add-product?id=123&userId=456
//
// The method has a receiver:
//
//	(app *Application)
//
// which means every handler can directly access:
//
//	app.userCollection
//	app.prodCollection
func (app *Application) AddProductToCart() gin.HandlerFunc {

	return func(ctx *gin.Context) {

		//----------------------------------------------------
		// Read Product ID
		//----------------------------------------------------

		// Query() reads a value from the URL query string.
		//
		// Example:
		//
		// /add-product?id=123
		//
		// returns "123".
		productID := ctx.Query("id")

		if productID == "" {

			ctx.AbortWithError(
				http.StatusBadRequest,
				errors.New("product ID is missing"),
			)

			return
		}

		//----------------------------------------------------
		// Read User ID
		//----------------------------------------------------

		userID := ctx.Query("userId")

		if userID == "" {

			ctx.AbortWithError(
				http.StatusBadRequest,
				errors.New("user ID is missing"),
			)

			return
		}

		//----------------------------------------------------
		// Database logic will go here.
		//----------------------------------------------------

		// Typical flow:
		//
		// 1. Find product using productID.
		// 2. Find user using userID.
		// 3. Append product to user's cart.
		// 4. Update MongoDB.
	}
}