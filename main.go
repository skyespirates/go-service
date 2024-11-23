package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"rest-api/routes"
	"rest-api/utils"
)

func main() {
	utils.InitLogger()
	router := gin.Default()

	// Add logger middleware
	router.Use(utils.LoggerMiddleware())

	// Register task routes
	routes.RegisterTaskRoutes(router)

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok", "message": "homepage"})
	})

	// Listen and Server in 0.0.0.0:8080
	router.Run(":8080")
}
