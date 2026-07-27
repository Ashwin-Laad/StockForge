package middleware

import (
	"Mod/token"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Authorization() gin.HandlerFunc {

	//middleware to check the token
	//if token not valid then does not move forward

	return func(ctx *gin.Context) {

		//gets the token from the front end

		Ctoken := ctx.Request.Header.Get("token")
		if Ctoken == "" {
			ctx.JSON(http.StatusInternalServerError, "No Token provided")
			ctx.Abort()
			return

		}

		s, msg := token.ValidateToken(Ctoken)
		if msg != "" {
			ctx.JSON(http.StatusInternalServerError, "No Token provided")
			ctx.Abort()
			return
		}
		//here s stores all the data from the token
		//using get i can receive all the values

		ctx.Set("email", s.Email)
		ctx.Set("name", s.Name)

		//we need to use next for middleware for it to move forward
		ctx.Next()

	}
}
