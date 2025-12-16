package todo

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Handler handles HTTP requests for todo items.
type Handler struct {
	svc Service
}

// NewHandler creates a new Handler with the given Service.
func NewHandler(svc Service) *Handler {
	return &Handler{svc: svc}
}

// RegisterRoutes registers /api/todos routes on a RouterGroup.
func (h *Handler) RegisterRoutes(rg *gin.RouterGroup) {
	rg.GET("/", h.ListTodos)
	rg.GET("/:id", h.GetTodo)
	rg.POST("/", h.CreateTodo)
	rg.PUT("/:id", h.UpdateTodo)
	rg.DELETE("/:id", h.DeleteTodo)
}

// ListTodos handles GET /api/todos
func (h *Handler) ListTodos(c *gin.Context) {
	todos, err := h.svc.ListTodos(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch todos"})
		return
	}
	c.JSON(http.StatusOK, todos)
}

// GetTodo handles GET /api/todos/:id
func (h *Handler) GetTodo(c *gin.Context) {
	id, err := parseIDParam(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	todo, err := h.svc.GetTodo(c.Request.Context(), id)
	if err != nil {
		if errors.Is(err, ErrTodoNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "todo not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch todo"})
		return
	}

	c.JSON(http.StatusOK, todo)
}

// CreateTodo handles POST /api/todos
type createTodoRequest struct {
	Title string `json:"title" binding:"required"`
}

// CreateTodo handles POST /api/todos
func (h *Handler) CreateTodo(c *gin.Context) {
	var req createTodoRequest
	if err := c.ShouldBindJSON(&req); err != nil || req.Title == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "title is required"})
		return
	}

	todo, err := h.svc.CreateTodo(c.Request.Context(), CreateTodoInput{
		Title: req.Title,
	})
	if err != nil {
		if errors.Is(err, ErrInvalidInput) {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create todo"})
		return
	}

	c.JSON(http.StatusCreated, todo)
}

// UpdateTodo handles PUT /api/todos/:id
type updateTodoRequest struct {
	Title     string `json:"title" binding:"required"`
	Completed bool   `json:"completed"`
}

// UpdateTodo handles PUT /api/todos/:id
func (h *Handler) UpdateTodo(c *gin.Context) {
	id, err := parseIDParam(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	var req updateTodoRequest
	if err := c.ShouldBindJSON(&req); err != nil || req.Title == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "title is required"})
		return
	}

	todo, err := h.svc.UpdateTodo(c.Request.Context(), id, UpdateTodoInput{
		Title:     req.Title,
		Completed: req.Completed,
	})
	if err != nil {
		if errors.Is(err, ErrTodoNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "todo not found"})
			return
		}
		if errors.Is(err, ErrInvalidInput) {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update todo"})
		return
	}

	c.JSON(http.StatusOK, todo)
}

// DeleteTodo handles DELETE /api/todos/:id
func (h *Handler) DeleteTodo(c *gin.Context) {
	id, err := parseIDParam(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	err = h.svc.DeleteTodo(c.Request.Context(), id)
	if err != nil {
		if errors.Is(err, ErrTodoNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "todo not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to delete todo"})
		return
	}

	c.Status(http.StatusNoContent)
}

// parseIDParam parses a string ID parameter to uint.
func parseIDParam(s string) (uint, error) {
	n, err := strconv.ParseUint(s, 10, 64)
	return uint(n), err
}
