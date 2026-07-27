package routes

import (
	"Mod/controllers"
	"Mod/database"
	"Mod/middleware"
	"fmt"

	"github.com/gin-gonic/gin"
)

// UserRouter registers all API routes.
func UserRouter(router *gin.Engine) {

	// Get MongoDB collections
	userCollection := database.GetCollection(database.Client, "User")
	productCollection := database.GetCollection(database.Client, "Products")

	// Create one Application instance that holds the collections
	app := controllers.CreateApplication(
		userCollection,
		productCollection,
	)

	//----------------------------------------------------
	// Authentication (Public Routes)
	//----------------------------------------------------
	//
	// These endpoints are accessible without a JWT token.
	// A new user cannot already possess a token, so signup
	// and login must remain public.
	//

	router.POST("/user/signup", controllers.SignUp())
	router.POST("/user/login", controllers.Login())

	//----------------------------------------------------
	// Protected Routes
	//----------------------------------------------------
	//
	// Every route registered under this group will first
	// execute the Authorization middleware.
	//
	// Request
	//     │
	//     ▼
	// Authorization()
	//     │
	//     ▼
	// Token Valid?
	//     │
	//  Yes │ No
	//     │
	//     ▼
	// Controller
	//

	fmt.Println("UnCroosed middleware")
	protected := router.Group("/")
	protected.Use(middleware.Authorization())

	fmt.Println("Croosed middleware")
	//----------------------------------------------------
	// Cart
	//----------------------------------------------------

	protected.POST("/user/add-product", app.AddProductToCart())
	protected.DELETE("/user/remove-product", app.RemoveProductFromCart())
	protected.POST("/user/buy", app.BuyItemFromCart())

	//----------------------------------------------------
	// Address
	//----------------------------------------------------

	protected.POST("/user/address", app.AddAddress())

	//----------------------------------------------------
	// Cart Total
	//----------------------------------------------------

	protected.GET("/user/cart/total", app.GetTotalFromUserCart())

	//----------------------------------------------------
	// Products
	//----------------------------------------------------
	//
	// Product browsing is public because users generally
	// don't need to be logged in just to view products.
	// If you want only logged-in users to view products,
	// simply move these two routes into the protected group.
	//

	router.POST("/products", app.AddProduct())

	router.GET("/products", app.SearchProducts())
	router.GET("/products/search", app.SearchProductsByName())

	router.GET("/instantbuy", app.InstantBuy())


	router.GET("/get-cart",app.GetTotalFromUserCart())
}
