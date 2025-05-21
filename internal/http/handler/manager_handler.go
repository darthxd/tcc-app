package handler

import (
	"fmt"
	"net/http"

	"github.com/darthxd/tcc-app/internal/config"
	"github.com/darthxd/tcc-app/internal/http/auth"
	"github.com/darthxd/tcc-app/internal/schemas"
	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
)

func AuthenticateManager(c echo.Context) schemas.Manager {
	db = config.GetDB()
	cookie, err := c.Cookie("session")
	if err != nil {
		log.Error(err)
	}
	sessions := auth.GetSessions()
	manager := schemas.Manager{}
	fmt.Print(sessions) // debug
	for _, s := range sessions {
		if s.SessionId == cookie.Value {
			if err := db.Where("user = ? AND password = ?", s.User, s.Password).First(&manager).Error; err != nil {
				log.Error(err)
			}
		}
	}
	return manager
}

func ManagerHome(c echo.Context) error {
	manager := AuthenticateManager(c)
	return c.Render(http.StatusOK, "manager_home", echo.Map{
		"manager": manager,
	})
}
