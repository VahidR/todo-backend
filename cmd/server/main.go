package main

import (
	"log"

	"github.com/VahidR/todo-backend/internal/config"
	"github.com/VahidR/todo-backend/internal/database"
	"github.com/VahidR/todo-backend/internal/router"
	"github.com/VahidR/todo-backend/internal/todo"
)

func main() {
	cfg := config.Load()
	db := database.NewMySQL(cfg.DBDSN)

	// Wiring: repository -> service -> handler -> router
	todoRepo := todo.NewRepository(db)
	todoService := todo.NewService(todoRepo)
	todoHandler := todo.NewHandler(todoService)

	r := router.New(todoHandler)

	log.Printf("Starting server on :%s (%s)", cfg.Port, cfg.Env)
	if err := r.Run(":" + cfg.Port); err != nil {
		log.Fatalf("server stopped: %v", err)
	}
}
