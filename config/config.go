package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

var (
	db   *gorm.DB
	port string
)

func Init() {
	// Load .env and start the database
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
	connStr := os.Getenv("DATABASE_URL")
	port = os.Getenv("PORT")
	db = DatabaseInit(connStr)
}

func GetDB() *gorm.DB {
	return db
}

func GetPort() string {
	return port
}
