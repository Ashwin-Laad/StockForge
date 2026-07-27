package controllers

import (
	"Mod/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// func DeleteAddress() gin.HandlerFunc {
// 	return func(ctx *gin.Context) {
// 		//get the user query
// 		userID := ctx.Query("id")
// 		if userID == "" {
// 			ctx.JSON(http.StatusBadRequest, gin.H{"error": "No user Id "})
// 			ctx.Abort() //stops future handlers
// 			return
// 		}

// 		address := make([]models.Address, 0)
// 		userHexID, err := primitive.ObjectIDFromHex(userID)
// 		// Step 1: Find the document whose _id matches userID.
// 		filter := bson.M{
// 			"_id": userHexID,
// 		}

// 		// Step 2: Build the update document.
// 		// "$set" tells MongoDB to update only the specified fields
// 		// instead of replacing the entire document.
// 		update := bson.M{
// 			"$set": bson.M{
// 				// Change the "address" field to the new value.
// 				"address": address,
// 			},
// 		}

// 		// Step 3: Execute the update.

// 		_, err = app.userCollection.UpdateOne(ctx, filter, update)
// 		if err != nil {
// 			return
// 		}

// 		ctx.JSON(http.StatusAccepted, gin.H{"status": "User address is deleted"})
// 	}
// }

func (app *Application) AddAddress() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userID := ctx.Query("id")
		if userID == "" {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "No user Id "})
			ctx.Abort() //stops future handlers
			return
		}
		var address models.Address //empty struct

		//Bind json auto decodes and inserts the value in the structs
		//bind object by address
		err := ctx.BindJSON(&address)

		address.AddressID = primitive.NewObjectID()
		if err != nil {
			ctx.JSON(http.StatusBadRequest, "No proper address passed")
			return
		}

		//address is binded

		//find the user and insert to address
		userHexID, err := primitive.ObjectIDFromHex(userID)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": "Invalid user ID",
			})
			return
		}

		filter := bson.M{"_id": userHexID}

		//update
		update := bson.M{
			"$push": bson.M{
				"address": address,
			},
		}

		_, error := app.userCollection.UpdateOne(ctx, filter, update)
		if error != nil {
			ctx.JSON(http.StatusBadRequest, "Error inserting error")
			return
		}

		ctx.JSON(http.StatusAccepted, "Inserted Address")
	}
}
