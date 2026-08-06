package controllers

import (
	"Mod/database"
	"Mod/models"
	"Mod/token"

	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) string {

	// Parameters:
	//
	// []byte(password)
	//     bcrypt works with bytes instead of strings.
	//
	// bcrypt.DefaultCost
	//     Determines how much computation bcrypt performs.
	//     Higher cost = more secure but slower.
	//
	// Returns:
	//
	// hashedPassword []byte
	// err            error
	hashedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(password),
		bcrypt.DefaultCost,
	)

	// If hashing fails, return an empty string.
	if err != nil {
		return ""
	}

	// Convert []byte back into a string before storing it.
	return string(hashedPassword)
}

// VerifyPassword compares the password entered by the user
// with the hashed password stored in MongoDB.

func VerifyPassword(
	userPassword string,
	hashedPassword string,
) (bool, string) {

	// CompareHashAndPassword hashes the user's entered password
	// internally and compares it against the stored hash.
	//
	// Parameters:
	//
	// []byte(hashedPassword)
	//     Password hash stored in MongoDB.
	//
	// []byte(userPassword)
	//     Password entered during login.
	err := bcrypt.CompareHashAndPassword(
		[]byte(hashedPassword),
		[]byte(userPassword),
	)

	// Passwords do not match.
	if err != nil {
		return false, err.Error()
	}

	// Password is correct.
	return true, ""
}

// SignUp returns a Gin handler responsible for registering a new user.

func SignUp() gin.HandlerFunc {

	return func(ctx *gin.Context) {

		// Read incoming JSON
		// user will hold the JSON received from the client.

		var user models.User

		// BindJSON:
		//
		// Reads the HTTP request body.
		// Converts JSON into the User struct.

		if err := ctx.BindJSON(&user); err != nil {

			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": "Invalid JSON",
			})
			return
		}

		//Validate fields

		validate := validator.New()
		//validates the fields with the annotations provided

		if err := validate.Struct(user); err != nil {

			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})

			return
		}

		//Access User collection

		userCollection := database.GetCollection(database.Client, "User")

		//  Check if email already exists
		count, err := userCollection.CountDocuments(
			ctx,
			bson.M{
				"email": user.Email,
			},
		)

		if err != nil {

			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})

			return
		}

		if count > 0 {

			ctx.JSON(http.StatusConflict, gin.H{
				"error": "Email already exists",
			})

			return
		}

		//  Hash password

		user.Password = HashPassword(user.Password)

		//  Generate server-side fields
		//start with token creation

		token, refreshToken, err := token.GenerateToken(
			user.FirstName,
			user.Email,
		)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": "Failed to generate token",
			})
			return
		}
		user.Token = token
		user.RefreshToken = refreshToken

		user.CreatedAt = time.Now()
		user.UpdatedAt = time.Now()

		user.UserID = primitive.NewObjectID().Hex() //human readable form of the hex id (stored as binary)

		user.UserCart = make([]models.ProductUser, 0)

		user.AddressDetails = make([]models.Address, 0)

		user.OrderStatus = make([]models.Order, 0)

		//  Insert into MongoDB

		_, err = userCollection.InsertOne(ctx, user)

		if err != nil {

			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})

			return
		}

		//  Success response

		ctx.JSON(http.StatusCreated, gin.H{
			"message": "Successfully signed up",
		})
	}
}

// Login authenticates an existing user.
func Login() gin.HandlerFunc {

	return func(ctx *gin.Context) {

		//  Read JSON

		var user models.User

		if err := ctx.BindJSON(&user); err != nil {

			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": "Invalid JSON",
			})

			return
		}

		//Get collection

		userCollection := database.GetCollection(database.Client, "User")

		//  Find user by email

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

		//  Verify password

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

		// Generate new tokens

		token, refreshToken, err := token.GenerateToken(
			foundUser.FirstName,
			foundUser.Email,
		)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": "Failed to generate token",
			})
			return
		}
		foundUser.Token = token
		foundUser.RefreshToken = refreshToken

		//  Send response

		ctx.JSON(http.StatusOK, gin.H{
			"message": "Login successful",
			"token":   token,
			"user_id": foundUser.ID,
		})
	}
}

func (app *Application) AddProduct() gin.HandlerFunc {

	return func(ctx *gin.Context) {

		productCollection := database.GetCollection(database.Client, "Products")

		var product models.Product

		// Read JSON

		if err := ctx.BindJSON(&product); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": "Invalid product data",
			})
			return
		}

		// Generate MongoDB ObjectID

		product.ProductID = primitive.NewObjectID()

		// Insert product

		_, err := productCollection.InsertOne(ctx, product)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": "Failed to add product",
			})
			return
		}

		ctx.JSON(http.StatusCreated, gin.H{
			"message": "Product added successfully",
		})
	}
}

// func (app *Application) BuyDirectly() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		UserQueryID := c.Query("userid")
// 		if UserQueryID == "" {
// 			fmt.Println("UserID is empty")
// 			c.JSON(http.StatusBadRequest, "Id not given")
// 			return
// 		}
// 		ProductQueryID := c.Query("pid")
// 		if ProductQueryID == "" {
// 			fmt.Println("Product_ID id is empty")
// 			c.JSON(http.StatusBadRequest, "product_id is empty")
// 		}
// 		productID, err := primitive.ObjectIDFromHex(ProductQueryID)
// 		if err != nil {
// 			fmt.Println(err)
// 			c.AbortWithStatus(http.StatusInternalServerError)
// 			return
// 		}

// 		err = database.InstantBuyer(c, app.prodCollection, app.userCollection, productID, UserQueryID)
// 		if err != nil {
// 			c.IndentedJSON(http.StatusInternalServerError, err)
// 		}
// 		c.IndentedJSON(200, "Successully placed the order")
// 	}
// }

// gives the total and user cart
func (app *Application) GetTotalFromUserCart() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		userID := ctx.Query("id") //string format
		if userID == "" {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": "No user ID",
			})

			return
		}

		//converts object id in hexadecimal form to object id in mongo
		UserHexId, err := primitive.ObjectIDFromHex(userID)

		// // primitive.E represents one field of a MongoDB document. A single key pair value .M is a map

		// _, err:= app.userCollection.FindOne(ctx,bson.D{primitive.E{Key: "_id", Value: UserHexId}})

		if err != nil {
			return
		}

		var user models.User
		userCollection := database.GetCollection(database.Client, "User")
		err = userCollection.FindOne(
			ctx,
			bson.M{"_id": UserHexId},
		).Decode(&user)

		if err != nil {
			return
		}

		total := 0

		for _, product := range user.UserCart {
			total += int(product.ProductPrice)
		}

		ctx.JSON(http.StatusOK, gin.H{
			"total": total,
			"cart":  user.UserCart,
		})
	}
}
