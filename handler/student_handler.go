package handler

import (
	"net/http"

	"github.com/labstack/echo"
)

func StudentInfo(c echo.Context) error {
	return c.Render(http.StatusOK,"student_home", echo.Map{
		"title":"Informações",
		"active":"info",
	})
}

func StudentMail(c echo.Context) error {
	return c.Render(http.StatusOK, "student_mail", echo.Map{
		"title":"E-mail",
		"active":"email",
	})
}
