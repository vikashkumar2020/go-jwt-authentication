package routes


import (
	"github.com/gin-gonic/gin"
	controller "go-jwt-authentication/controllers"
	middleware "go-jwt-authentication/middleware"
)

func UserRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.Use(middleware.Authenticate())
	incomingRoutes.GET("/users", controller.GetUsers())
	incomingRoutes.GET("/users/:id", controller.GetUser())
}