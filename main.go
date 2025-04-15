package main

import (
	"log"
	"net/http"
	"text/template"

	"github.com/labstack/echo"

	"github.com/darthxd/tcc-app/auth"
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
	port := config.GetPort()

	// Set up the main routes

	e.GET("/", func(c echo.Context) error { return c.Redirect(http.StatusFound, "/login/") })

	// Student routes
	student := e.Group("/aluno")
	{
		student.GET("", func(c echo.Context) error { return c.Redirect(http.StatusFound, "/aluno/") })
		student.GET("/", auth.DefaultMiddleware("student", func(c echo.Context) error {
			return c.Redirect(http.StatusFound, "/aluno/info")
		}))
		student.GET("/info", auth.DefaultMiddleware("student", handler.StudentInfo))
		student.GET("/email", auth.DefaultMiddleware("student", handler.StudentMail))
		student.GET("/sair", auth.DefaultMiddleware("student", handler.LogOut))
	}

	// Teacher routes
	teacher := e.Group("/professor")
	{
		teacher.GET("/", auth.DefaultMiddleware("teacher", handler.TeacherRender))
	}

	// Management routes
	manager := e.Group("/gerenciamento")
	{
		manager.GET("/", auth.DefaultMiddleware("manager", handler.ManagerRender))
	}

	// Login page routes
	login := e.Group("/login")
	{
		login.GET("", func(c echo.Context) error { return c.Redirect(http.StatusFound, "/login/") })
		login.GET("/", auth.LoginMiddleware())
		login.GET("/aluno", handler.LoginStudentRender)
		login.GET("/professor", handler.LoginTeacherRender)
		login.GET("/supervisao", handler.LoginManagerRender)
		login.POST("/aluno", handler.LoginStudent)
	}

	// Run the server
	if err := e.Start(port); err != nil {
		log.Fatal(err)
	}
}
