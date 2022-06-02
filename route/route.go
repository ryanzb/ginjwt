package route

import (
	"ginjwt/controller"

	"github.com/gin-gonic/gin"
)

type Route struct {
	authController controller.AuthController
	userController controller.UserController
}

func New(
	authController controller.AuthController, 
	userController controller.UserController,
) *Route {
	return &Route{
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

	{
		app.GET("/api/user", r.userController.Info)
		app.PUT("/api/user", r.userController.Update)
	}
}