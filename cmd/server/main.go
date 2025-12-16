package main

import (
	"log"

	"github.com/VahidR/todo-backend/internal/config"
	"github.com/VahidR/todo-backend/internal/database"
	"github.com/VahidR/todo-backend/internal/router"
	"github.com/VahidR/todo-backend/internal/todo"
)

// The main function initializes the application components and starts the server.
func main() {
	// Load configuration and initialize database connection
	cfg := config.Load()
	db := database.NewMySQL(cfg.DBDSN)

	// Wiring: repository -> service -> handler -> router
	todoRepo := todo.NewRepository(db)
	todoService := todo.NewService(todoRepo)
	todoHandler := todo.NewHandler(todoService)

	// Initialize and start the router
	r := router.New(todoHandler)

	log.Printf("Starting server on :%s (%s)", cfg.Port, cfg.Env)
	if err := r.Run(":" + cfg.Port); err != nil {
		log.Fatalf("server stopped: %v", err)
	}
}
