package controller

import (
	"ginjwt/dto"
	"ginjwt/response"
	"ginjwt/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
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
	userID := ctx.GetString("user_id")
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

	userID := ctx.GetString("user_id")
	if strconv.Itoa(int(req.ID)) != userID {
		resp := response.BuildErrorResponse("user id wrong")
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
