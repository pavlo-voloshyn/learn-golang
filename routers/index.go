package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/pavlo-voloshyn/init/midlewares"
)

func InitRoutes(router *gin.Engine) {
	api := router.Group("/api", midlewares.AuthorizeJWT())
	SetStudentRouts(api)
	auth := router.Group("/auth")
	SetAuthRoutes(auth)
}
