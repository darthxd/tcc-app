package handler

import (
	"net/http"

	"github.com/labstack/echo"
)

type Login struct {
	Name     string
	User     string
	Password string
	Type     string
	ID       int
}

func LoginRender(c echo.Context) error {
	return c.Render(http.StatusOK, "login_page", echo.Map{
		"title":"Login",
	})
}

func LoginStudentRender(c echo.Context) error {
	return c.Render(http.StatusOK, "login_student", nil)
}

func LoginTeacherRender(c echo.Context) error {
	return c.Render(http.StatusOK, "login_teacher", nil)
}

func LoginManagerRender(c echo.Context) error {
	return c.Render(http.StatusOK, "login_manager", nil)
}
