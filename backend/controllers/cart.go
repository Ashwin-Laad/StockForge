package controllers

import (
	"Mod/database"
	"Mod/models"
	"time"

	// "context"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// Instead of creating or fetching the database collections inside every
// controller function, we store them once in an Application object.

// Every method attached to Application can directly access these
// collections using:
//
//	app.userCollection
//	app.prodCollection
type Application struct {

	// A collection represents one group of documents
	// inside the database.
	userCollection *mongo.Collection

	// prodCollection points to the Products collection.
	prodCollection *mongo.Collection
}

func CreateApplication(
	userCollection *mongo.Collection,
	prodCollection *mongo.Collection,
) *Application {
	//pointer to avoid copying
	return &Application{
		userCollection: userCollection,
		prodCollection: prodCollection,
	}
}

//	(app *Application)
//
// which means every handler can directly access:
//
//	app.userCollection
//	app.prodCollection

// verifies the ids before adding to datatbase
func (app *Application) AddProductToCart() gin.HandlerFunc {

	return func(ctx *gin.Context) {

		// Read Product ID

		// Query() reads a value from the URL query string.
		// ? reads it
		productID := ctx.Query("iD")

		if productID == "" {

			ctx.AbortWithError(
				http.StatusBadRequest,
				errors.New("product ID is missing"),
			)

			return
		}

		// Read User ID

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

		// Validate the Product ID

		productObjectID, err := primitive.ObjectIDFromHex(productID)

		if err != nil {

			ctx.AbortWithError(
				http.StatusBadRequest,
				errors.New("invalid product ID"),
			)

			return
		}

		err = database.AddProductToCart(ctx, app.userCollection, app.prodCollection, productObjectID, ObjectUserId)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, err.Error())
			return
		}

		ctx.JSON(http.StatusAccepted, "added product to cart")

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

// /users/GetProducts

func (app *Application) GetProducts() gin.HandlerFunc {

	// Returns every product stored in the Products collection.
	return func(ctx *gin.Context) {
		// start := time.Now()

		// products will hold every document retrieved from MongoDB.
		var products []models.Product

		//before checking mongodb check for cached data

		println("Redis client", database.RedisClient)

		ctxt, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()
		cachedData, err := database.RedisClient.Get(ctxt, "products").Result()
		// cachedData, err := database.RedisGet("products")

		fmt.Println("REDIS DATA:", cachedData)
		fmt.Println("REDIS ERROR:", err)
		if err == nil && cachedData != "" {
			err := json.Unmarshal([]byte(cachedData), &products)
			if err != nil {
				return
			}

			ctx.JSON(http.StatusOK, products)
			return

		}
		// Empty filter.
		//
		// bson.D{{}} tells MongoDB:
		// "Return every document in this collection."

		print("In process to fetch from mongodb")
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
		// if err != nil {

		// 	ctx.JSON(http.StatusInternalServerError, gin.H{
		// 		"error": err.Error(),
		// 	})

		// 	return
		// }
		if err != nil {
			fmt.Println("MONGO FIND ERROR:", err)

			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}
		data, err := json.Marshal(products)
		if err != nil {

			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})

			return
		}

		err = database.RedisClient.Set(ctx, "products", data, 5*time.Minute).Err()
		if err != nil {
			fmt.Println("REDIS SET ERROR:", err)
		} else {
			fmt.Println("REDIS SET SUCCESS")
		}
		// err = database.RedisSet("products", string(data), 300)
		// if err != nil {
		// 	ctx.JSON(http.StatusInternalServerError, gin.H{
		// 		"error": err.Error(),
		// 	})

		// 	return
		// }
		// elapsed := time.Since(start)

		// println("Time: ", elapsed)
		// Gin automatically marshals the slice into JSON.
		ctx.JSON(http.StatusOK, products)

	}
}
func (app *Application) SearchProductsByName() gin.HandlerFunc {

	return func(ctx *gin.Context) {

		// Holds all matching products.
		products := make([]models.Product, 0)

		// Read the search query from the URL.

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
