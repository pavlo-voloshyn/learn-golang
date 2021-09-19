package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pavlo-voloshyn/init/entities"
	"github.com/pavlo-voloshyn/init/services"
)

type AuthController interface {
	Login(ctx *gin.Context)
}

type authController struct {
	loginService services.AuthService
	jWtService   services.JWTService
}

func NewAuthController(authService services.AuthService,
	jWtService services.JWTService) AuthController {
	return &authController{
		loginService: authService,
		jWtService:   jWtService,
	}
}

func (controller *authController) Login(c *gin.Context) {
	var credential entities.LoginCredentials
	err := c.BindJSON(&credential)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err)
	}
	isUserAuthenticated := controller.loginService.Login(credential.Email, credential.Password)
	if isUserAuthenticated {
		token := controller.jWtService.Generate(credential.Email, true)
		c.JSON(200, gin.H{
			"token": token,
		})
		return
	}

	c.AbortWithStatusJSON(http.StatusBadRequest, "not found")
}
