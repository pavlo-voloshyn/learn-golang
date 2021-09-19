package main

import (
	"github.com/gin-gonic/gin"
	"github.com/pavlo-voloshyn/init/routers"
)

func main() {
	router := gin.Default()
	routers.InitRoutes(router)
	router.Run(":8080")
}
