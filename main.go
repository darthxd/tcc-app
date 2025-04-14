package main

import (
	"log"
	"os"
	"text/template"

	"github.com/joho/godotenv"
	"github.com/labstack/echo"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/darthxd/tcc-app/handler"
	"github.com/darthxd/tcc-app/models"
)

func main() {
	// Set up basic server settings 
	e := echo.New()
	t := &handler.Template{Templates: template.Must(template.ParseGlob("public/views/*.html"))}
	e.Renderer = t
	e.Static("/assets", "public/static")

	// Load .env and start the database
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
	connStr := os.Getenv("DATABASE_URL")
	port := os.Getenv("PORT")
	db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{})
	if err != nil {
		log.Fatal("Error connecting to the database")
	}
	db.AutoMigrate(&models.Student{})

	// Set up the main routes
	student := e.Group("/aluno")
	{
		student.GET("", handler.StudentInfo)
		student.GET("/info", handler.StudentInfo)
		student.GET("/email", handler.StudentMail)
	}

	teacher := e.Group("/professor")
	{
		teacher.GET("/", handler.TeacherRender)
	}

	manager := e.Group("/gerenciamento")
	{
		manager.GET("/", handler.ManagerRender)
	}

	login := e.Group("/login")
	{
		login.GET("/", handler.LoginRender)
		login.GET("/aluno", handler.LoginStudentRender)
		login.GET("/professor", handler.LoginTeacherRender)
		login.GET("/supervisao", handler.LoginManagerRender)
	}

	// Run the server
	if err := e.Start(port); err != nil {
		log.Fatal(err)
	}
}
