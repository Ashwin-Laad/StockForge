package controllers

import (
	"Mod/database"
	"Mod/models"
	"Mod/token"

	"net/http"
	"time"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

// HashPassword converts a plain-text password into a secure hashed password.
//
// NOTE:
// The implementation is assumed to exist elsewhere.
// In production this would typically use bcrypt.
func HashPassword(password string) string

// VerifyPassword compares the password entered by the user with the
// hashed password stored in the database.
//
// Returns:
//
//	true  -> passwords match
//	false -> passwords do not match
//
// The string return value usually contains an error message.
func VerifyPassword(userPassword string, hashedPassword string) (bool, string)

// SignUp returns a Gin handler responsible for registering a new user.
//
// Why return gin.HandlerFunc?
// ---------------------------
// Gin expects every route to receive a handler function.
//
// Instead of executing immediately, this function returns another
// function that Gin stores internally.
//
// Later, whenever a request arrives at:
//
//	POST /user/signup
//
// Gin executes the returned function.
func SignUp() gin.HandlerFunc {

	return func(ctx *gin.Context) {

		//--------------------------------------------------------
		// Step 1 : Read incoming JSON
		//--------------------------------------------------------

		// user will hold the JSON received from the client.
		//
		// Example JSON:
		//
		// {
		//   "first_name":"Ashwin",
		//   "email":"abc@gmail.com"
		// }
		var user models.User

		// BindJSON:
		//
		// Reads the HTTP request body.
		// Converts JSON into the User struct.
		//
		// The '&' passes the address of the struct because
		// BindJSON needs permission to modify it.
		if err := ctx.BindJSON(&user); err != nil {

			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": "Invalid JSON",
			})
			return
		}

		//--------------------------------------------------------
		// Step 2 : Validate fields
		//--------------------------------------------------------

		validate := validator.New()

		if err := validate.Struct(user); err != nil {

			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})

			return
		}

		//--------------------------------------------------------
		// Step 3 : Access User collection
		//--------------------------------------------------------

		userCollection := database.GetCollection(database.Client, "User")

		//--------------------------------------------------------
		// Step 4 : Check if email already exists
		//--------------------------------------------------------

		count, err := userCollection.CountDocuments(
			ctx,
			bson.M{
				"email": user.Email,
			},
		)

		if err != nil {

			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": "Database query failed",
			})

			return
		}

		if count > 0 {

			ctx.JSON(http.StatusConflict, gin.H{
				"error": "Email already exists",
			})

			return
		}

		//--------------------------------------------------------
		// Step 5 : Hash password
		//--------------------------------------------------------

		user.Password = HashPassword(user.Password)

		//--------------------------------------------------------
		// Step 6 : Generate server-side fields
		//--------------------------------------------------------

		token, refreshToken := token.GenerateToken(
			user.FirstName,
			user.Email,
		)

		user.Token = token
		user.RefreshToken = refreshToken

		user.CreatedAt = time.Now()
		user.UpdatedAt = time.Now()

		user.UserID = primitive.NewObjectID().Hex()

		user.UserCart = make([]models.ProductUser, 0)

		user.AddressDetails = make([]models.Address, 0)

		user.OrderStatus = make([]models.Order, 0)

		//--------------------------------------------------------
		// Step 7 : Insert into MongoDB
		//--------------------------------------------------------

		_, err = userCollection.InsertOne(ctx, user)

		if err != nil {

			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": "Failed to create user",
			})

			return
		}

		//--------------------------------------------------------
		// Step 8 : Success response
		//--------------------------------------------------------

		ctx.JSON(http.StatusCreated, gin.H{
			"message": "Successfully signed up",
		})
	}
}

// Login authenticates an existing user.
func Login() gin.HandlerFunc {

	return func(ctx *gin.Context) {

		//--------------------------------------------------------
		// Step 1 : Read JSON
		//--------------------------------------------------------

		var user models.User

		if err := ctx.BindJSON(&user); err != nil {

			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": "Invalid JSON",
			})

			return
		}

		//--------------------------------------------------------
		// Step 2 : Get collection
		//--------------------------------------------------------

		userCollection := database.GetCollection(database.Client, "User")

		//--------------------------------------------------------
		// Step 3 : Find user by email
		//--------------------------------------------------------

		var foundUser models.User

		err := userCollection.FindOne(
			ctx,
			bson.M{
				"email": user.Email,
			},
		).Decode(&foundUser)

		if err != nil {

			if err == mongo.ErrNoDocuments {

				ctx.JSON(http.StatusUnauthorized, gin.H{
					"error": "Invalid email",
				})

				return
			}

			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": "Database error",
			})

			return
		}

		//--------------------------------------------------------
		// Step 4 : Verify password
		//--------------------------------------------------------

		isValid, message := VerifyPassword(
			user.Password,
			foundUser.Password,
		)

		if !isValid {

			ctx.JSON(http.StatusUnauthorized, gin.H{
				"error": message,
			})

			return
		}

		//--------------------------------------------------------
		// Step 5 : Generate new tokens
		//--------------------------------------------------------

		token, refreshToken := token.GenerateToken(
			foundUser.FirstName,
			foundUser.Email,
		)

		foundUser.Token = token
		foundUser.RefreshToken = refreshToken

		//--------------------------------------------------------
		// Step 6 : Send response
		//--------------------------------------------------------

		ctx.JSON(http.StatusOK, gin.H{
			"message": "Login successful",
			"token": token,
		})
	}
}

func AddProducts() gin.HandlerFunc {
	return nil
}

func SearchProduct() gin.HandlerFunc {
	return nil
}

func SeeProducts() gin.HandlerFunc {
	return nil
}