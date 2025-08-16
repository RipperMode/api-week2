package handlers

import (
	"net/http"
	"strconv"

	"github.com/<you>/api-week2/internal/store"
	"github.com/gin-gonic/gin"
)

type TaskHandler struct {
	Store *store.MemStore
}

// POST /tasks
func (h *TaskHandler) Create(c *gin.Context) {
	var input struct {
		Title string `json:"title" binding:"required"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
		return
	}
	task := h.Store.Create(input.Title)
	c.JSON(http.StatusCreated, task)
}

// GET /tasks
func (h *TaskHandler) List(c *gin.Context) {
	c.JSON(http.StatusOK, h.Store.List())
}

// GET /tasks/:id
func (h *TaskHandler) Get(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	task, ok := h.Store.Get(id)
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
		return
	}
	c.JSON(http.StatusOK, task)
}

// PUT /tasks/:id
func (h *TaskHandler) Update(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	var input struct {
		Title string `json:"title"`
		Done  bool   `json:"done"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
		return
	}
	task, ok := h.Store.Update(id, input.Title, input.Done)
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
		return
	}
	c.JSON(http.StatusOK, task)
}

// DELETE /tasks/:id
func (h *TaskHandler) Delete(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	if !h.Store.Delete(id) {
		c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
		return
	}
	c.Status(http.StatusNoContent)
}
