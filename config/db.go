package config

import (
	"log"

	"github.com/darthxd/tcc-app/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func DatabaseInit(connStr string) *gorm.DB {
	db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{})
	if err != nil {
		log.Fatal("Error connecting to the database")
	}
	db.AutoMigrate(&models.Student{}, &models.Teacher{}, &models.Manager{})
	return db
}
