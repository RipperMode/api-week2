package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// Create router
	r := gin.Default() // includes Logger and Recovery middleware by default

	// Health check route
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})
	})

	// Start server on port 8080
	r.Run(":8080")
}
