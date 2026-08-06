package database

import (
	"Mod/models"
	"errors"

	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"

	// "go.mongodb.org/mongo-driver/v2/bson"
	// "go.mongodb.org/mongo-driver/v2/mongo"

	"go.mongodb.org/mongo-driver/bson"

	"go.mongodb.org/mongo-driver/mongo"
)

//called from the handler to actually perform the database functions

func AddProductToCart(ctx *gin.Context, userCollection *mongo.Collection, prodCollection *mongo.Collection, productID primitive.ObjectID, userID primitive.ObjectID) error {

	cursor, err := prodCollection.Find(ctx, bson.M{"_id": productID})

	if err != nil {

		return err
	}

	//Cart specific to user
	var ProductCart []models.ProductUser
	err = cursor.All(ctx, &ProductCart)
	if err != nil {

		return err
	}

	filter := bson.M{
		"_id": userID,
	}
	update := bson.M{
		"$push": bson.M{
			"user_cart": bson.M{
				"$each": ProductCart,
			},
		},
	}
	_, err = userCollection.UpdateOne(ctx, filter, update)
	if err != nil {

		return err
	}

	return nil

}
func RemoveProductFromCart(
	ctx *gin.Context,
	userCollection *mongo.Collection,
	prodCollection *mongo.Collection,
	productID primitive.ObjectID,
	userID primitive.ObjectID,
) error {

	// Find the user document.

	filter := bson.M{
		"_id": userID,
	}

	//pull removes an element from the array

	update := bson.M{
		"$pull": bson.M{
			"user_cart": bson.M{
				"_id": productID,
			},
		},
	}

	// Update the user's cart.

	result, err := userCollection.UpdateOne(
		ctx,
		filter,
		update,
	)

	if err != nil {
		return err
	}

	// No user found.

	if result.MatchedCount == 0 {
		return errors.New("user not found")
	}

	// User exists but the product was not present
	// in the cart.

	if result.ModifiedCount == 0 {
		return errors.New("product not found in cart")
	}
	//remove product from cache

	return nil
}
func BuyProductFromCart(ctx *gin.Context,
	userCollection *mongo.Collection,
	prodCollection *mongo.Collection,
	userID primitive.ObjectID,
) error {

	var getCartItem models.User
	var order models.Order

	order.OrderID = primitive.NewObjectID()
	order.OrderedAt = time.Now()
	order.OrderCart = make([]models.ProductUser, 0)
	order.PaymentMethod.Digital = true

	// Unwind the user's cart
	//unwind allows for aggregation operations

	// {
	//     "_id": 101,
	//     "user_cart": [
	//         { "product_name": "Laptop", "product_price": 50000 },
	//         { "product_name": "Mouse",  "product_price": 1000  },
	//         { "product_name": "Keyboard","product_price": 3000 }
	//     ]
	// }
	//
	// MongoDB cannot directly calculate the sum of prices inside an array.

	// After $unwind:
	//
	// Document 1 -> Laptop
	// Document 2 -> Mouse
	// Document 3 -> Keyboard
	//
	// Each product now behaves like its own MongoDB document,

	unwind := bson.D{
		{
			Key: "$unwind",
			Value: bson.D{
				{
					Key:   "path",
					Value: "$user_cart",
				},
			},
		},
	}

	//Group the documents and calculate the total price

	//
	// After $unwind, every product exists as a separate document.
	//
	// We now regroup those documents using the user's "_id"
	// and calculate the total price of every product in the cart.

	// "$sum" adds the value of the "product_price"
	// field from every unwound document.
	//

	//
	// {
	//     "_id": 101,
	//     "total": 54000
	// }
	//
	group := bson.D{
		{
			Key: "$group",
			Value: bson.D{
				{
					// Group all products belonging to the same user.
					Key:   "_id",
					Value: "$_id",
				},
				{
					// Name of the newly created field.
					Key: "total",

					// Calculate the sum of every product's price.
					Value: bson.D{
						{
							Key: "$sum",

							// Read the product_price field
							// from every unwound cart item.
							Value: "$user_cart.product_price",
						},
					},
				},
			},
		},
	}

	// Execute the aggregation pipeline

	//
	// Aggregate() executes every stage in order.
	//
	// Stage 1:
	//     $unwind
	//
	// Stage 2:
	//     $group
	//
	// MongoDB passes the output of one stage
	// directly into the next stage.
	//
	// Pipeline:
	//
	// User Document
	//      |
	//      ▼
	//  $unwind
	//      |
	//      ▼
	// One document per cart item
	//      |
	//      ▼
	//   $group
	//      |
	//      ▼
	// Total cart price
	//
	cursor, err := userCollection.Aggregate(
		ctx,
		mongo.Pipeline{
			unwind,
			group,
		},
	)

	if err != nil {
		return err
	}

	var getcartItems []bson.M //an object of key, value mongodb return a cursor of document so have to contain in slice

	err = cursor.All(ctx, &getcartItems)
	if err != nil {
		return err
	}

	var total_price int

	//index , userobj
	for _, userC := range getcartItems {
		total_price = int(userC["total"].(int32))

	}
	order.Price = total_price
	filter := bson.D{
		{Key: "_id", Value: userID},
	}

	update := bson.D{
		{
			Key: "$push",
			Value: bson.D{
				{Key: "orders", Value: order},
			},
		},
	}
	_, err = userCollection.UpdateMany(ctx, filter, update)
	if err != nil {
		return err
	}

	filter = bson.D{
		{Key: "_id", Value: userID},
	}
	err = userCollection.FindOne(ctx, filter).Decode(&getCartItem)
	if err != nil {
		return err
	}

	filter2 := bson.D{
		{Key: "_id", Value: userID},
	}

	update2 := bson.M{
		"$push": bson.M{
			"orders.$[].order_list": bson.M{
				"$each": getCartItem.UserCart,
			},
		},
	}

	_, err = userCollection.UpdateOne(ctx, filter2, update2)
	if err != nil {
		return err
	}

	user_cartEmpty := make([]models.ProductUser, 0)

	//get me the document with the matcing Id from the user table

	filter3 := bson.D{
		{Key: "_id", Value: userID},
	}

	update3 := bson.D{
		{Key: "$set", Value: bson.D{{Key: "user_cart", Value: user_cartEmpty}}},
	}
	//now update the collection
	//change the cart to be empty
	_, err = userCollection.UpdateOne(ctx, filter3, update3)
	if err != nil {
		return err
	}
	var productIDs []primitive.ObjectID

	for _, product := range getCartItem.UserCart {
		productIDs = append(productIDs, product.ProductID)
	}

	productFilter := bson.M{
		"_id": bson.M{
			"$in": productIDs,
		},
	}

	_, err = prodCollection.DeleteMany(
		ctx,
		productFilter,
	)

	if err != nil {
		return err
	}
	// err = RedisClient.Del(ctx, "products").Err()
	// if err != nil {
	// 	return err
	// }
	return nil

}

// func InstantBuyer(ctx *gin.Context, prodCollection, userCollection *mongo.Collection, productID primitive.ObjectID, UserID string) error {

// 	id, err := primitive.ObjectIDFromHex(UserID)
// 	if err != nil {
// 		log.Println(err)
// 		return err
// 	}
// 	var product_details models.ProductUser
// 	var orders_detail models.Order

// 	orders_detail.OrderID = primitive.NewObjectID()
// 	orders_detail.OrderedAt = time.Now()
// 	orders_detail.OrderCart = make([]models.ProductUser, 0)
// 	orders_detail.PaymentMethod.COD = true

// 	err = prodCollection.FindOne(ctx, bson.D{primitive.E{Key: "_id", Value: productID}}).Decode(&product_details)
// 	if err != nil {
// 		log.Println(err)
// 	}

// 	orders_detail.Price = product_details.ProductPrice

// 	filter := bson.D{primitive.E{Key: "_id", Value: id}}
// 	update := bson.D{{Key: "$push", Value: bson.D{primitive.E{Key: "orders", Value: orders_detail}}}}
// 	_, err = userCollection.UpdateOne(ctx, filter, update)

// 	if err != nil {
// 		log.Println(err)
// 	}

// 	filter2 := bson.D{primitive.E{Key: "_id", Value: id}}
// 	update2 := bson.M{"$push": bson.M{"orders.$[].order_cart": product_details}}
// 	_, err = userCollection.UpdateOne(ctx, filter2, update2)
// 	if err != nil {
// 		log.Println(err)
// 	}
// 	update3 := bson.M{"$push": bson.M{"user_cart": product_details}}
// 	_, err = userCollection.UpdateOne(ctx, filter2, update3)
// 	if err != nil {
// 		log.Println(err)
// 	}
// 	return nil
// }

func GetOrder() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userCollection := GetCollection(Client, "user")
		userId := ctx.Query("user_id")
		filter2 := bson.D{primitive.E{Key: "_id", Value: userId}}

		var foundUser models.User
		err := userCollection.FindOne(ctx, filter2).Decode(&foundUser)
		if err != nil {
			ctx.JSON(500, "Cannot get order collection")
			return
		}

		UserOrder := foundUser.OrderStatus

		ctx.JSON(http.StatusAccepted, UserOrder)

	}
}
