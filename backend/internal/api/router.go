package api

import (
	"backend/internal/api/routes"
	"os"

	"github.com/gin-gonic/gin"
)

// SetupRouter initializes the router and registers all routes
func SetupRouter() *gin.Engine {
	// Set Gin mode
	gin.SetMode(os.Getenv("GIN_MODE"))

	// Create router
	router := gin.Default()

	// Development routes
	if os.Getenv("GIN_MODE") != "release" {
		devGroup := router.Group("/dev")
		{
			devGroup.POST("/seed", routes.SeedHandler)
		}
	}

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	router.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"message": "Page not found"})
	})

	router.GET("/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Hello, World!"})
	})

	return router
}
