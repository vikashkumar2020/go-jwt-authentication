package routes


import (
	"github.com/gin-gonic/gin"
	controller "go-jwt-authentication/controllers"
)

func AuthRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("users/signup", controller.Signup())
	incomingRoutes.POST("users/login", controller.Login())
}