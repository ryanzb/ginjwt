package route

import (
	"ginjwt/controller"
	"ginjwt/middleware"
	"ginjwt/service"

	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(NewRoute)

type Route struct {
	App            *gin.Engine
	jwtService     service.JWTService
	authController controller.AuthController
	userController controller.UserController
}

func NewRoute(
	jwtService service.JWTService,
	authController controller.AuthController,
	userController controller.UserController,
) *Route {
	r := &Route{
		App:            gin.Default(),
		jwtService:     jwtService,
		authController: authController,
		userController: userController,
	}
	r.Register()
	return r
}

func (r *Route) Register() {
	authGroup := r.App.Group("/api/auth")
	{
		authGroup.POST("/login", r.authController.Login)
		authGroup.POST("/register", r.authController.Register)
	}

	userGroup := r.App.Group("/api/user", middleware.AuthorizeJWT(r.jwtService))
	{
		userGroup.GET("", r.userController.Info)
		userGroup.PUT("", r.userController.Update)
	}
}
