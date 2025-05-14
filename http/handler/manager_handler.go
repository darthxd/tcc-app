package handler

import (
	"net/http"

	"github.com/labstack/echo"
)

func ManagerRender(c echo.Context) error {
	return c.Render(http.StatusOK,"manager_home", echo.Map{})
}
