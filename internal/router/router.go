package router

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"github.com/VahidR/todo-backend/internal/todo"
)

// New creates a new Gin engine with routes registered.
// Returns the *gin.Engine instance.
// HINT: It 'injects' the todo.Handler
func New(todoHandler *todo.Handler) *gin.Engine {
	// HINT: Create a new Gin router
	router := gin.New()
	// HINT: Use() middleware with Logger and Recovery (crash-free) passed inside
	// HINT: Logger() is used to log HTTP requests
	// HINT: Recovery() is used to recover from any panics and writes a 500 if there was one
	router.Use(gin.Logger(), gin.Recovery())

	// CORS configuration
	// HINT: Obviously, we inject CORS middleware in the router. Another use case of Use() method.
	router.Use(cors.New(cors.Config{
		// this should be read from the ENV in a real-world application
		AllowOrigins:     []string{"http://localhost:4321"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// Register API routes
	// HINT: This is how they group the apis under /api and /api/todos
	api := router.Group("/api")
	{
		// HINT: This is how  they embed the inner resources
		todos := api.Group("/todos")
		// HINT: Attachment of routes to the handler
		// HINT: Technically, it is the START of the request flows insided the application from a dev perspective.
		todoHandler.RegisterRoutes(todos)
	}

	return router
}
