package config

import (
	"gorm.io/gorm"
)

var (
	db   *gorm.DB
	port string
)

func Init() {
	port = ":8080"
	// db = InitSQLite()
	db = InitMySQL()
}

func GetDB() *gorm.DB {
	return db
}

func GetPort() string {
	return port
}
