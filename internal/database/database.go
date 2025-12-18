package database

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/VahidR/todo-backend/internal/todo"
)

// NewMySQL initializes a new GORM DB connection to a MySQL database.
// It returns the *gorm DB instance.
// HINT: The 'constructor' function for DB connection.
func NewMySQL(dbConfig string) *gorm.DB {
	// HINT: It opens the gorm database connection with MySQL driver and gorm config
	db, err := gorm.Open(mysql.Open(dbConfig), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to DB: %v", err)
	}

	// Auto-migrate schema
	// HINT: This will migrate (create/update) the "todos" table based on the Todo model.
	if err := db.AutoMigrate(&todo.Todo{}); err != nil {
		log.Fatalf("failed to migrate DB: %v", err)
	}

	return db
}
