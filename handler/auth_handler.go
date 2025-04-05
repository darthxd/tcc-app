package handler

import (
	"net/http"

	"github.com/labstack/echo"
)

func AuthRender(c echo.Context) error {
	return c.Render(http.StatusOK,"auth", echo.Map{})
}
