package routes
import(
	"github.com/gin-gonic/gin"
	// "Mod/controllers"
)
//passed down a router from the main function 
//will implement it

func UserRouter(router *gin.Engine){
	//using the gin router create different routes to be handled

	router.POST("user/signUp",)
	router.POST("/user/login")
	router.POST("/user/add-products")
	router.GET("/searchProduct")
	router.GET("/product-view")
}