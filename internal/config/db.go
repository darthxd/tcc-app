package config

import (
	"log"

	"github.com/darthxd/tcc-app/internal/schemas"
	"github.com/glebarez/sqlite"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitSQLite() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("./database/main.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Error connecting to the database")
	}
	db.AutoMigrate(&schemas.Student{}, &schemas.Teacher{}, &schemas.Manager{})
	return db
}

func InitMySQL() *gorm.DB {
	dsn := "root:@tcp(127.0.0.1:3306)/main?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Error connecting to MySQL")
	}
	db.AutoMigrate(&schemas.Student{}, &schemas.Teacher{}, &schemas.Manager{})
	return db
}
