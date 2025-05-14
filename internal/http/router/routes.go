package router

import (
	"net/http"

	"github.com/darthxd/tcc-app/internal/http/auth"
	"github.com/darthxd/tcc-app/internal/http/handler"
	"github.com/labstack/echo"
)

func routesInit(e *echo.Echo) {
	// Default page redirection
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
		student.GET("/conta", auth.DefaultMiddleware("student", handler.StudentAccount))
		student.GET("/sair", auth.DefaultMiddleware("student", handler.LogOut))
	}

	// Teacher routes
	teacher := e.Group("/professor")
	{
		teacher.GET("", func(c echo.Context) error { return c.Redirect(http.StatusFound, "/professor/") })
		teacher.GET("/", auth.DefaultMiddleware("teacher", handler.TeacherRender))
	}

	// Management routes
	manager := e.Group("/gerenciamento")
	{
		manager.GET("", func(c echo.Context) error { return c.Redirect(http.StatusFound, "/gerenciamento/") })
		manager.GET("/", auth.DefaultMiddleware("manager", handler.ManagerRender))
	}

	// Login page routes
	login := e.Group("/login")
	{
		login.GET("", func(c echo.Context) error { return c.Redirect(http.StatusFound, "/login/") })
		login.GET("/", auth.LoginMiddleware())
		login.GET("/aluno", handler.LoginStudentRender)
		login.GET("/professor", handler.LoginTeacherRender)
		login.GET("/gerenciamento", handler.LoginManagerRender)
		login.POST("/aluno", handler.LoginStudent)
	}

}
