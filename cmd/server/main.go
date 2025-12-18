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
	// HINT: Dependency Injection (DI) happens in the package-level.
	// HINT: So you for having access to confing.Load() method, you need to import the config package here.
	cfg := config.Load()
	// 2. Initialize database connection
	// HINT: Similarly, for having access to database.NewMySQL() method, you need to import the database package here.
	db := database.NewMySQL(cfg.DBDSN)

	// 3. Wiring: repository -> service -> handler -> router
	// HINT: inject every layer in the main just like below.
	todoRepo := todo.NewRepository(db)
	todoService := todo.NewService(todoRepo)
	todoHandler := todo.NewHandler(todoService)

	// 4. Initialize and start the router
	// HINT: they pass "handler" to "router". Makes sense, as router routes to different handlers.
	r := router.New(todoHandler)

	log.Printf("Starting server on :%s (%s)", cfg.Port, cfg.Env)
	// 5. Finally, start the server
	if err := r.Run(":" + cfg.Port); err != nil {
		log.Fatalf("server stopped: %v", err)
	}
}
