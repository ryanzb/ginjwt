package middleware

import (
	"fmt"
	"ginjwt/response"
	"ginjwt/service"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func AuthorizeJWT(jwtService service.JWTService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		header := ctx.GetHeader("Authorization")
		if header == "" {
			resp := response.BuildErrorResponse("no token provided")
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, resp)
			return
		}

		token, err := jwtService.ValidateToken(header)
		if err != nil {
			resp := response.BuildErrorResponse(err.Error())
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, resp)
			return
		}
		if token == nil || !token.Valid {
			resp := response.BuildErrorResponse("token invalid")
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, resp)
			return
		}

		clams := token.Claims.(jwt.MapClaims)
		userID := fmt.Sprintf("%v", clams["user_id"])
		ctx.Set("user_id", userID)
	}
}