package main

import (
	"log"

	"api-test/internal/config"
	"api-test/internal/handlers"
	"api-test/internal/store"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.LoadConfig()

	mem := store.NewMemStore()
	h := &handlers.TaskHandler{Store: mem}

	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	r.POST("/tasks", h.Create)
	r.GET("/tasks", h.List)
	r.GET("/tasks/:id", h.Get)
	r.PUT("/tasks/:id", h.Update)
	r.DELETE("/tasks/:id", h.Delete)

	if err := r.Run(":" + cfg.Port); err != nil {
		log.Fatal(err)
	}
}
