package database

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/VahidR/todo-backend/internal/todo"
)

// NewMySQL initializes a new GORM DB connection to a MySQL database.
func NewMySQL(dsn string) *gorm.DB {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to DB: %v", err)
	}

	// Auto-migrate schema
	// This will create/update the "todos" table based on the Todo model.
	if err := db.AutoMigrate(&todo.Todo{}); err != nil {
		log.Fatalf("failed to migrate DB: %v", err)
	}

	return db
}
