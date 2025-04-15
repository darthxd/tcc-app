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
		student.GET("/", auth.StudentMiddleware(func(c echo.Context) error {
			return c.Redirect(http.StatusFound, "/aluno/info")
		}, func(c echo.Context) error { 
			return c.Redirect(http.StatusFound, "/login/") 
		}))
		student.GET("/info", auth.StudentMiddleware(handler.StudentInfo, func(c echo.Context) error {
			return c.Redirect(http.StatusFound, "/login/")
		}))
		student.GET("/email", auth.StudentMiddleware(handler.StudentMail, func(c echo.Context) error {
			return c.Redirect(http.StatusFound, "/login/")
		}))
		student.GET("/sair", auth.StudentMiddleware(handler.LogOut, func(c echo.Context) error {
			return c.Redirect(http.StatusFound, "/login/")
		}))
	}

	// Teacher routes
	teacher := e.Group("/professor")
	{
		teacher.GET("/", handler.TeacherRender)
	}

	// Management routes
	manager := e.Group("/gerenciamento")
	{
		manager.GET("/", handler.ManagerRender)
	}

	// Login page routes
	login := e.Group("/login")
	{
		login.GET("", func(c echo.Context) error { return c.Redirect(http.StatusFound, "/login/") })
		login.GET("/", handler.LoginPageRender)
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
