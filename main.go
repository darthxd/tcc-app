package main

import (
	"log"
	"text/template"

	"github.com/labstack/echo"

	"github.com/darthxd/tcc-app/handler"
)

func main() {
	e := echo.New()
	t := &handler.Template{Templates: template.Must(template.ParseGlob("public/views/*.html"))}
	e.Renderer = t
	e.Static("/assets", "public/static")

	student := e.Group("/aluno")
	{
		student.GET("/", handler.StudentRender)
	}

	teacher := e.Group("/professor")
	{
		teacher.GET("/", handler.TeacherRender)
	}

	manager := e.Group("/gerenciamento")
	{
		manager.GET("/", handler.ManagerRender)
	}

	auth := e.Group("/conta")
	{
		auth.GET("", handler.AuthRender)
	}

	if err := e.Start(":8080"); err != nil {
		log.Fatal(err)
	}
}
