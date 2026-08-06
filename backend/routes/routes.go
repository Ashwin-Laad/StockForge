package routes

import (
	"Mod/controllers"
	"Mod/database"
	"Mod/middleware"
	

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

	//public routes
	// These endpoints are accessible without a JWT token.
	

	router.POST("/user/signup", controllers.SignUp())
	router.POST("/user/login", controllers.Login())


	// Every route registered under this group will first
	// execute the Authorization middleware.
	


	protected := router.Group("/")
	protected.Use(middleware.Authorization())


	// Cart

	//binded these function to app as database interaction is needed

	protected.POST("/user/add-product", app.AddProductToCart())
	protected.DELETE("/user/remove-product", app.RemoveProductFromCart())
	protected.POST("/user/buy", app.BuyItemFromCart())

	
	// Address


	protected.POST("/user/address", app.AddAddress())

	// Cart Total
	

	protected.GET("/user/cart/total", app.GetTotalFromUserCart())

	
	// Products
	
	// Product browsing is public 
	//no token is required
	router.POST("/products", app.AddProduct())

	router.GET("/products", app.GetProducts())
	router.GET("/products/search", app.SearchProductsByName())

	// router.GET("/instantbuy", app.InstantBuy())

	router.GET("/get-cart", app.GetTotalFromUserCart())

	router.GET("/prev-order", database.GetOrder())
}
