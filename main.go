package main

import (
	"Mod/middleware"
	"Mod/routes"
	"log"
	"github.com/gin-gonic/gin"
)

func main(){

	//set port to 8080
	Port:="8080"

	//*gin.Engine an engine instance with logger and recovery preinstalled
	router:= gin.Default()

	//pass the router to routes to check for the "basic routes first"
	routes.UserRouter(router)
	//add the middleware for authentication
	router.Use(middleware.UseAuth())

	//add the extra routes



	//start the server at the port
	log.Fatal(router.Run(":"+Port))
}
