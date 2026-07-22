package controllers


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
//binding this fucntion to the interfaec
func (app *Application) AddProductToCart(ctx *gin.Context)(gin.HandlerFunc){

	ProductID=ctx.query("id")
	if ProductID==""{
		//abort need the product id
		ctx.AbortWithError(505,errors.New("The product id is missing"))
		return 
	}

	ProductID=ctx.query("Userid")
	if ProductID==""{
		//abort need the product id
		ctx.AbortWithError(505,errors.New("The user id is missing"))
		return 
	}

 //database add product to cart


}