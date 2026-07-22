package database

import (
	"errors"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

//define the errors--> custom errors

var(
	CantFindProduct=errors.New("")
	CantUpdateUser=errors.New("")
)
//create an application struct that merges the user and the product collection
type Application interface{
	UserCollection *mongo.Collection
	ProdCollection *mongo.Collection
}

//initialise application
func CreateApplication(UserCollection, ProdCollection * mongo.Collection)(*Application){
	return &Application{   
		
	UserCollection:UserCollection,
	ProdCollection:ProdCollection
}
}
//ADDs the product to the cart
func AddProductToCart(ctx *gin.Context)(gin.HandlerFunc){

	ProductID=ctx.query("id")
	if ProductID==""{
		//abort need the product id
		ctx.AbortWithError(505,errors.New("The product id is missing"))
		return 
	}


}
func RemoveProductToCart(){}
func BuyProductToCart(){}
func InstantBuy(){}