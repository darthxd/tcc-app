package auth

import (
	"github.com/darthxd/tcc-app/config"
	"github.com/labstack/echo"
)

func StudentMiddleware(next, stop echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		is_authenticated := false
		db = config.GetDB()
		cookie, err := c.Cookie("session")
		if err != nil {
			return stop(c)
		}
		for _, s := range sessions {
			if s.SessionId == cookie.Value && s.Type == "student" {
				is_authenticated = true
			}
		}
		if is_authenticated {
			return next(c)
		} else {
			return stop(c)
		}
	}
}
