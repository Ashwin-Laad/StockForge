package main

import (
	"Mod/database"
	"Mod/routes"
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {



	//set port to 8000
	Port := "8000"

	//initialise redis
   database.ConnectRedis();
	//redis client created

	//*gin.Engine an engine instance with logger and recovery preinstalled
	router := gin.Default()

	// //add the middleware for authentication
	// router.Use(middleware.Authorization())
	router.Use(cors.New(cors.Config{
		AllowAllOrigins: true,
		AllowMethods: []string{
			"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS",
		},
		AllowHeaders: []string{
			"Origin",
			"Content-Type",
			"Accept",
			"Token",
		},
	}))

	//pass the router to routes to check for the "basic routes first"
	routes.UserRouter(router)

	//start the server at the port
	log.Fatal(router.Run(":" + Port))
}
