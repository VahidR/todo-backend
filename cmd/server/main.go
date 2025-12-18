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
	// 1. Load configuration and initialize database connection
	cfg := config.Load()
	// 2. Initialize database connection
	db := database.NewMySQL(cfg.DBDSN)

	// 3. Wiring: repository -> service -> handler -> router
	todoRepo := todo.NewRepository(db)
	todoService := todo.NewService(todoRepo)
	todoHandler := todo.NewHandler(todoService)

	// 4. Initialize and start the router
	r := router.New(todoHandler)

	log.Printf("Starting server on :%s (%s)", cfg.Port, cfg.Env)
	// 5. Finally, start the server
	if err := r.Run(":" + cfg.Port); err != nil {
		log.Fatalf("server stopped: %v", err)
	}
}
