package main

import (
	"github.com/gin-gonic/gin"
	routes "go-jwt-authentication/routes"
	"os"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "8000"
	}

	router := gin.New()
	router.Use(gin.Logger())

	routes.AuthRoutes(router)
	routes.UserRoutes(router)

	router.GET("/api-1", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"Success": "Access Granted for api 1",
		})
	})

	router.GET("/api-2", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"Success": "Access Granted for api 2",
		})
	})

	router.Run(":" + port)

}
