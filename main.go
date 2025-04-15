package main

import (
	"log"
	"text/template"

	"github.com/labstack/echo"

	"github.com/darthxd/tcc-app/config"
	"github.com/darthxd/tcc-app/handler"
)

func main() {
	// Set up basic server settings 
	e := echo.New()
	t := &handler.Template{Templates: template.Must(template.ParseGlob("public/views/*.html"))}
	e.Renderer = t
	e.Static("/assets", "public/static")

	config.Init()

	// Set up the main routes
	student := e.Group("/aluno")
	{
		student.GET("", handler.StudentInfo)
		student.GET("/info", handler.StudentInfo)
		student.GET("/email", handler.StudentMail)
		student.GET("/sair", handler.LogOut)
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
		login.POST("/aluno", handler.LoginStudent)
	}

	// Run the server
	if err := e.Start(":5432"); err != nil {
		log.Fatal(err)
	}
}
