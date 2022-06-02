package controller

import (
	"fmt"
	"ginjwt/dto"
	"ginjwt/response"
	"ginjwt/service"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

type UserController interface {
	Info(ctx *gin.Context)
	Update(ctx *gin.Context)
}

type userController struct {
	userSerice service.UserService
	jwtService service.JWTService
}

func NewUserController(
	userSerice service.UserService,
	jwtService service.JWTService,
) UserController {
	return &userController{
		userSerice: userSerice,
		jwtService: jwtService,
	}
}

func (c *userController) Info(ctx *gin.Context) {
	header := ctx.GetHeader("Authorization")
	if header == "" {
		resp := response.BuildErrorResponse("no token")
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, resp)
		return
	}
	token, err := c.jwtService.ValidateToken(header)
	if err != nil {
		resp := response.BuildErrorResponse("validate token failed: " + err.Error())
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, resp)
		return
	}
	if token == nil {
		resp := response.BuildErrorResponse("token invalid")
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, resp)
		return
	}

	clams := token.Claims.(jwt.MapClaims)
	userID := fmt.Sprintf("%v", clams["user_id"])
	user, err := c.userSerice.FindUserByID(userID)
	if err != nil {
		resp := response.BuildErrorResponse(err.Error())
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, resp)
		return
	}

	resp := response.BuildResponse(true, "", user)
	ctx.JSON(http.StatusOK, resp)
}

func (c *userController) Update(ctx *gin.Context) {
	var req dto.UpdateUserRequest
	if err := ctx.ShouldBind(&req); err != nil {
		resp := response.BuildErrorResponse(err.Error())
		ctx.AbortWithStatusJSON(http.StatusBadRequest, resp)
		return
	}

	user, err := c.userSerice.UpdateUser(&req)
	if err != nil {
		resp := response.BuildErrorResponse(err.Error())
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, resp)
		return
	}

	resp := response.BuildResponse(true, "", user)
	ctx.JSON(http.StatusOK, resp)
}

func userNil(user *dto.UserResponse) bool {
	return user == nil || user.ID == 0
}
