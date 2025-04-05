package handler

import (
	"net/http"

	"github.com/labstack/echo"
)

func TeacherRender(c echo.Context) error {
	return c.Render(http.StatusOK,"teacher_home", echo.Map{})
}
