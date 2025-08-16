package main

import (
	"github.com/<you>/api-week2/internal/handlers"
	"github.com/<you>/api-week2/internal/store"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Health check
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	// Init store + handler
	mem := store.NewMemStore()
	h := &handlers.TaskHandler{Store: mem}

	// Task routes
	r.POST("/tasks", h.Create)
	r.GET("/tasks", h.List)
	r.GET("/tasks/:id", h.Get)
	r.PUT("/tasks/:id", h.Update)
	r.DELETE("/tasks/:id", h.Delete)

	r.Run(":8080")
}
