package controllers

import (
	"Mod/database"
	"Mod/models"
	"errors"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
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

// verifies the ids before adding to datatbase
func (app *Application) AddProductToCart() gin.HandlerFunc {

	return func(ctx *gin.Context) {

		fmt.Println("Inside ")
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
		productID := ctx.Query("iD")

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

		userID := ctx.Query("id")

		if userID == "" {

			ctx.AbortWithError(
				http.StatusBadRequest,
				errors.New("user ID is missing"),
			)

			return
		}

		ObjectUserId, err := primitive.ObjectIDFromHex(userID)
		if err != nil {
			ctx.AbortWithError(http.StatusBadRequest, errors.New("invalid user ID"))
			return
		}
		//----------------------------------------------------
		// Validate the Product ID
		//----------------------------------------------------

		// MongoDB stores document IDs as ObjectIDs, not plain strings.
		//
		// However, values received from an HTTP request are always strings.
		//
		// Example URL:
		//
		//	/add-product?id=687ab3d8d0c9e7f5d8c21a34
		//
		// Here, productID is:
		//
		//	"687ab3d8d0c9e7f5d8c21a34"
		//
		// Before querying MongoDB, this hexadecimal string must be converted
		// into MongoDB's ObjectID type.
		//
		// ObjectIDFromHex performs two tasks:
		//
		//   1. Checks whether the string is a valid 24-character hexadecimal ObjectID.
		//   2. Converts it into primitive.ObjectID.
		//
		// Returns:
		//
		//	productObjectID : primitive.ObjectID
		//	err             : error
		//
		// If the string is malformed, err will not be nil.
		productObjectID, err := primitive.ObjectIDFromHex(productID)

		if err != nil {

			// Never use println() in production code.
			// AbortWithError sends the error to Gin's error handling system
			// and stops processing this request.
			ctx.AbortWithError(
				http.StatusBadRequest,
				errors.New("invalid product ID"),
			)

			return
		}

		//----------------------------------------------------
		// Database logic will go here.
		//----------------------------------------------------

		//can acess app.pro and app.user because it is binded to the struct

		fmt.Println("Going to database ")
		err = database.AddProductToCart(ctx, app.userCollection, app.prodCollection, productObjectID, ObjectUserId)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, err.Error())
			return
		}
		fmt.Println("crossed database ")
		ctx.JSON(http.StatusAccepted, "added product to cart")
		// Typical flow:
		//
		// 1. Find product using productID.
		// 2. Find user using userID.
		// 3. Append product to user's cart.
		// 4. Update MongoDB.

	}
}
func (app *Application) RemoveProductFromCart() gin.HandlerFunc {

	return func(ctx *gin.Context) {
		//get the id of product
		ProductID := ctx.Query("id")
		if ProductID == "" {
			ctx.AbortWithError(
				http.StatusBadRequest,
				errors.New("product ID is missing"),
			)

			return
		}
		UserId := ctx.Query("user_id")
		if UserId == "" {
			ctx.AbortWithError(
				http.StatusBadRequest,
				errors.New("product ID is missing"),
			)

			return
		}
		//check if product string is malformed or not
		//project id is acual object id stored in mongo db

		productObjectID, err := primitive.ObjectIDFromHex(ProductID)

		if err != nil {

			// Never use println() in production code.
			// AbortWithError sends the error to Gin's error handling system
			// and stops processing this request.
			ctx.AbortWithError(
				http.StatusBadRequest,
				errors.New("invalid product ID"),
			)

			return
		}

		userObjectID, err := primitive.ObjectIDFromHex(UserId)
		if err != nil {
			ctx.AbortWithError(http.StatusBadRequest, errors.New("invalid user ID"))
			return
		}
		//ctx is created by gin
		//add auto passed to handler func

		err = database.RemoveProductFromCart(ctx, app.userCollection, app.prodCollection, productObjectID, userObjectID)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, err)
			return
		}

		ctx.JSON(http.StatusAccepted, "Removed product from cart")

	}

}

func (app *Application) BuyItemFromCart() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		//get the user id
		UserId := ctx.Query("id")
		if UserId == "" {
			ctx.AbortWithError(
				http.StatusBadRequest,
				errors.New("User Id is missing"),
			)
			return
		}
		UserID, err := primitive.ObjectIDFromHex(UserId)
		if err != nil {
			ctx.AbortWithError(http.StatusBadRequest, errors.New("invalid user ID"))
			return
		}
		err = database.BuyProductFromCart(ctx, app.userCollection, app.prodCollection, UserID)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, err.Error())
			return
		}
		//run a goroutine for concurrent logging
		go func() {
			file, err := os.OpenFile("User.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, err.Error())
				return
			}
			defer file.Close()

			file.WriteString(fmt.Sprintf("userId %s bought a Product", UserId))
		}()
		ctx.JSON(http.StatusAccepted, "Item bought from cart")

	}
}

// /users/searchProducts

func (app *Application) SearchProducts() gin.HandlerFunc {

	// Returns every product stored in the Products collection.
	return func(ctx *gin.Context) {

		// products will hold every document retrieved from MongoDB.
		var products []models.Product

		// Empty filter.
		//
		// bson.D{{}} tells MongoDB:
		// "Return every document in this collection."
		cursor, err := app.prodCollection.Find(ctx, bson.D{{}})
		if err != nil {

			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})

			return
		}

		// Always close the cursor after use.
		defer cursor.Close(ctx)

		// Decode every document from the cursor into the products slice.
		err = cursor.All(ctx, &products)
		if err != nil {

			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})

			return
		}

		// Gin automatically marshals the slice into JSON.
		ctx.JSON(http.StatusOK, products)
	}
}
func (app *Application) SearchProductsByName() gin.HandlerFunc {

	return func(ctx *gin.Context) {

		// Holds all matching products.
		products := make([]models.Product, 0)

		// Read the search query from the URL.
		//
		// Example:
		// GET /search?name=laptop
		query := ctx.Query("name")

		if query == "" {

			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": "query parameter 'name' is required",
			})

			return
		}

		// Case-insensitive search on product_name.
		filter := bson.M{
			"product_name": bson.M{
				"$regex":   query,
				"$options": "i",
			},
		}

		cursor, err := app.prodCollection.Find(ctx, filter)

		if err != nil {

			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})

			return
		}

		defer cursor.Close(ctx)

		err = cursor.All(ctx, &products)

		if err != nil {

			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})

			return
		}

		ctx.JSON(http.StatusOK, products)
	}
}
