package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Config holds the application configuration.
type Config struct {
	Port  string
	DBDSN string
	Env   string
}

// Load reads configuration from .env file (if present) and environment variables.
func Load() *Config {
	// Load .env file if it exists (won't override existing env vars)
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using environment variables")
	}

	cfg := &Config{
		Port:  getEnv("PORT", "8080"),
		DBDSN: getEnv("DB_DSN", ""),
		Env:   getEnv("ENV", "development"),
	}

	if cfg.DBDSN == "" {
		log.Fatal("DB_DSN is required, e.g. user:pass@tcp(localhost:3306)/todo_db?parseTime=true&charset=utf8mb4&loc=Local")
	}
	return cfg
}

// getEnv retrieves the value of the environment variable named by the key.
// If the variable is empty or not present, it returns the provided default value.
func getEnv(key, def string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return def
}
