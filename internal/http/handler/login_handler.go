package handler

import (
	"net/http"

	"github.com/darthxd/tcc-app/internal/http/auth"
	"github.com/labstack/echo"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

func LoginPageRender(c echo.Context) error {
	return c.Render(http.StatusOK, "login_page", echo.Map{
		"title": "Entre na sua conta",
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

func LoginStudent(c echo.Context) error {
	user := c.FormValue("user")
	password := c.FormValue("password")
	session, err := auth.AuthenticateStudent(user, password)
	if err != nil {
		c.JSON(http.StatusBadRequest, echo.Map{
			"error": err,
		})
	}
	auth.SetCookie(c, "session", session.SessionId)

	c.Response().Header().Set("HX-Redirect", "/aluno")
	return nil
}

func LogOut(c echo.Context) error {
	auth.DeleteCookie(c, "session")

	c.Response().Header().Set("HX-Redirect", "/login")
	return nil
}
