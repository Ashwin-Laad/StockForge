package controllers

import (
	"Mod/database"
	"Mod/models"
	"fmt"
	"go/token"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func HashPassword(s string) string

func verifyPassword(userPass string, actualPass string) (bool, string)

func signUp() gin.HandlerFunc {
	//context here contains all info about the request object and allows for response
	return func(ctx *gin.Context) {

		//REMEMBER--> for incoming json cannot deal directly with it
		//we will have to bind the incoming json to a struct of Same type

		var user models.User
		// func (c *gin.Context) BindJSON(obj any) error--bound to the Context
		err := ctx.BindJSON(&user)
		if err != nil {
			fmt.Println("Error in Signing UP")
			return
		}

		validateErr := validator.New().Struct(user)
		if validateErr != nil {
			ctx.JSON(400, gin.H{"error": "Some info is wrong"})
			return
		}
		UserCollection := database.GetCollection(database.Client, "User")

		count, err := UserCollection.CountDocuments(ctx, bson.M{"email": user.Email})
		if err != nil {
			fmt.Println("Error in finding email")
			return
		}
		if count >= 1 {
			ctx.JSON(400, gin.H{"count": "more than 1"})
		}
		//all good means new user here
		user.Password = HashPassword(user.Password)

		//now for rest of the feilds server will have to fill it
		Token, Refresh_Token := token.GenerateToken(user.name, user.email)
		Created_At := time.Now()
		Updated_At := time.Now()
		UserID := primitive.NewObjectID()
		Usercart := make([]models.ProductUser, 0)
		Address_details := make([]models.Address, 0)
		OrderStatus := make([]models.Order, 0)

		//insert the user into database
		_, insertErr := UserCollection.InsertOne(ctx, user)
		if insertErr != nil {
			fmt.Println("Error in inserting email")
			return
		}

		//if everything was working
		ctx.JSON(200, "Successfully signed in")
	}

}
func Login() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		//REMEMBER--> for incoming json cannot deal directly with it
		//we will have to bind the incoming json to a struct of Same type

		var user models.User
		// func (c *gin.Context) BindJSON(obj any) error--bound to the Context
		err := ctx.BindJSON(&user)
		//even though extra feils exits its okay
		if err != nil {
			fmt.Println("Error in Logging UP")
			return
		}

		//gets me the table or collection via the client object 
	
		UserCollection := database.GetCollection(database.Client, "User")

		var foundUser models.User
		//decode is passed the address of found user
	    UserCollection.FindOne(ctx, bson.M{"email": user.Email}).Decode(&foundUser)
		if err != nil {
			fmt.Println("Error in finding email login")
			return
		}
		//verify password
		verify,message:=verifyPassword(user.Password, foundUser.Password)
		if verify!=true{
			ctx.JSON(400, "Wrong password"+message)
		}

		//verified generate token
		Token, Refresh_Token := token.GenerateToken(foundUser.First_Name, foundUser.email)

		c.JSON(200,"Login Success")

	}

}

func addProducts() gin.HandlerFunc {}

func SearchProduct() gin.HandlerFunc {}

func SeeProducts() gin.HandlerFunc {}
