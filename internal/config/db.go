package config

import (
	"log"

	"github.com/darthxd/tcc-app/internal/schemas"
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

func InitDatabase() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("./database/main.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Error connecting to the database")
	}
	db.AutoMigrate(&schemas.Student{}, &schemas.Teacher{}, &schemas.Manager{})
	return db
}
