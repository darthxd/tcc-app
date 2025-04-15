package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

var (
	db   *gorm.DB
)

func Init() {
	// Load .env and start the database
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
	connStr := os.Getenv("DATABASE_URL")
	db = DatabaseInit(connStr)
}

func GetDB() *gorm.DB {
	return db
}
