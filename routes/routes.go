package routes

import (
	"github.com/gin-gonic/gin"
	// "Mod/controllers"
)

// UserRouter registers every route related to users and products.
//
// Purpose:
// --------
// Gin itself does not know which function should execute when a request
// arrives. This function "teaches" Gin by mapping URL paths and HTTP
// methods to handler functions.
//
// Think of it like a receptionist:
//
// Client Request
//        │
//        ▼
//  UserRouter (checks URL + HTTP method)
//        │
//        ▼
// Appropriate Controller Function
//
// This function is usually called once during application startup
// from main.go.
func UserRouter(router *gin.Engine) {

	// router is a pointer to the Gin Engine.
	//
	// Type:
	// *gin.Engine
	//
	// The Gin Engine is the heart of a Gin application.
	// It stores:
	// - All registered routes
	// - Middleware
	// - Configuration
	// - HTTP server settings
	//
	// A pointer (*) is used because we want to modify the same router
	// instance that was created in main.go instead of making a copy.

	// Register a POST endpoint for user signup.
	//
	// HTTP Method:
	// POST
	//
	// URL:
	// /user/signUp
	//
	// Why POST?
	// Creating a new user changes data on the server,
	// so POST is the appropriate HTTP method.
	//
	// Normally the second argument would be a controller function like:
	//
	// controllers.SignUp()
	//
	// or
	//
	// controllers.SignUp
	//
	// depending on how the handler is written.
	router.POST("/user/signUp")

	// Register a POST endpoint for user login.
	//
	// The client sends credentials (email/password),
	// and the controller verifies them before generating
	// authentication tokens.
	router.POST("/user/login")

	// Register a POST endpoint to add products.
	//
	// Since products are being created or modified,
	// POST is used.
	router.POST("/user/add-products")

	// Register a GET endpoint for searching products.
	//
	// GET requests are used when the client only wants
	// to retrieve data without modifying anything.
	//
	// Example:
	// GET /searchProduct?name=iphone
	router.GET("/searchProduct")

	// Register a GET endpoint for viewing product details.
	//
	// Example:
	// GET /product-view?id=12345
	router.GET("/product-view")
}