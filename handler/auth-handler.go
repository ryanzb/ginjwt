package handler

import (
	"errors"
	"ginjwt/dto"
	"ginjwt/response"
	"ginjwt/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type AuthHandler interface {
	Login(ctx *gin.Context)
	Register(ctx *gin.Context)
}

type authHandler struct {
	authService service.AuthService
	jwtService  service.JWTService
	userService service.UserService
}

func NewAuthHandler(
	authService service.AuthService,
	jwtService service.JWTService,
	userService service.UserService,
) AuthHandler {
	return &authHandler{
		authService: authService,
		jwtService:  jwtService,
		userService: userService,
	}
}

func (c *authHandler) Login(ctx *gin.Context) {
	var req dto.LoginRequest
	if err := ctx.ShouldBind(&req); err != nil {
		resp := response.BuildErrorResponse("failed to login: " + err.Error())
		ctx.AbortWithStatusJSON(http.StatusBadRequest, resp)
		return
	}

	// 验证用户名和密码
	valid, err := c.authService.Verify(req.Email, req.Password)
	if err != nil {
		resp := response.BuildErrorResponse("failed to login: " + err.Error())
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, resp)
		return
	}

	if !valid {
		resp := response.BuildErrorResponse("failed to login: email or password wrong")
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, resp)
		return
	}

	user, _ := c.userService.FindUserByEmail(req.Email)

	// 生成token
	token, err := c.jwtService.GenerateToken(strconv.Itoa(int(user.ID)))
	if err != nil {
		resp := response.BuildErrorResponse("failed to login: " + err.Error())
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, resp)
		return
	}

	user.Token = token
	resp := response.BuildResponse(true, "", user)
	ctx.JSON(http.StatusOK, resp)
}

func (c *authHandler) Register(ctx *gin.Context) {
	var req dto.RegisterRequest
	if err := ctx.ShouldBind(&req); err != nil {
		resp := response.BuildErrorResponse("failed to register: " + err.Error())
		ctx.AbortWithStatusJSON(http.StatusBadRequest, resp)
		return
	}

	// 生成user
	user, err := c.userService.CreateUser(&req)
	if err != nil {
		resp := response.BuildErrorResponse("failed to register: " + err.Error())
		if errors.Is(err, service.ErrUserExists) {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, resp)
		} else {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, resp)
		}
		return
	}

	// 生成token
	token, err := c.jwtService.GenerateToken(strconv.Itoa(int(user.ID)))
	if err != nil {
		resp := response.BuildErrorResponse("failed to login: " + err.Error())
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, resp)
		return
	}

	user.Token = token
	resp := response.BuildResponse(true, "", user)
	ctx.JSON(http.StatusOK, resp)
}
