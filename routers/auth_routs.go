package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/pavlo-voloshyn/init/controllers"
	"github.com/pavlo-voloshyn/init/services"
)

var (
	jwt_service     = services.NewJWTService()
	auth_service    = services.NewAuthService()
	auth_controller = controllers.NewAuthController(auth_service, jwt_service)
)

func SetAuthRoutes(rg *gin.RouterGroup) {
	rg.GET("/login", auth_controller.Login)
}
