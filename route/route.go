package route

import (
	"ginjwt/handler"
	"ginjwt/middleware"
	"ginjwt/service"

	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(NewRoute)

type Route struct {
	App            *gin.Engine
	jwtService     service.JWTService
	authHandler handler.AuthHandler
	userHandler handler.UserHandler
}

func NewRoute(
	jwtService service.JWTService,
	authHandler handler.AuthHandler,
	userHandler handler.UserHandler,
) *Route {
	r := &Route{
		App:            gin.Default(),
		jwtService:     jwtService,
		authHandler: authHandler,
		userHandler: userHandler,
	}
	r.Register()
	return r
}

func (r *Route) Register() {
	authGroup := r.App.Group("/api/auth")
	{
		authGroup.POST("/login", r.authHandler.Login)
		authGroup.POST("/register", r.authHandler.Register)
	}

	userGroup := r.App.Group("/api/user", middleware.AuthorizeJWT(r.jwtService))
	{
		userGroup.GET("", r.userHandler.Info)
		userGroup.PUT("", r.userHandler.Update)
	}
}
