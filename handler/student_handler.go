package handler

import (
	"net/http"

	"github.com/labstack/echo"
)

func StudentRender(c echo.Context) error {
	return c.Render(http.StatusOK,"student_home", echo.Map{})
}
