package route

import (
	"ginjwt/controller"
	"ginjwt/middleware"
	"ginjwt/service"

	"github.com/gin-gonic/gin"
)

type Route struct {
	jwtService service.JWTService
	authController controller.AuthController
	userController controller.UserController
}

func New(
	jwtService service.JWTService,
	authController controller.AuthController,
	userController controller.UserController,
) *Route {
	return &Route{
		jwtService: jwtService,
		authController: authController,
		userController: userController,
	}
}

func (r *Route) Register(app *gin.Engine) {
	authGroup := app.Group("/api/auth")
	{
		authGroup.POST("/login", r.authController.Login)
		authGroup.POST("/register", r.authController.Register)
	}

	userGroup := app.Group("/api/user", middleware.AuthorizeJWT(r.jwtService))
	{
		userGroup.GET("", r.userController.Info)
		userGroup.PUT("", r.userController.Update)
	}
}
