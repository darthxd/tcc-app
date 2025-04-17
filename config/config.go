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
		log.Print("Error loading .env file")
		port = ":8080"
	} else {
		// connStr := os.Getenv("DATABASE_URL")
		port = os.Getenv("PORT")
	}
	// db = DatabaseInit(connStr)
	db = SQLiteInit()
}

func GetDB() *gorm.DB {
	return db
}

func GetPort() string {
	return port
}
