package auth

import (
	"net/http"

	"github.com/darthxd/tcc-app/config"
	"github.com/labstack/echo"
)

func LoginMiddleware() echo.HandlerFunc {
	return func(c echo.Context) error {
		var auth_type string
		is_authenticated := false
		db = config.GetDB()
		cookie, err := c.Cookie("session")
		if err != nil {
			return c.Render(http.StatusOK, "login_page", nil)
		}
		for _, s := range sessions {
			if s.SessionId == cookie.Value && s.Type == "student" {
				is_authenticated = true
				auth_type = "student"
			} else if s.SessionId == cookie.Value && s.Type == "teacher" {
				is_authenticated = true
				auth_type = "teacher"
			} else if s.SessionId == cookie.Value && s.Type == "manager" {
				is_authenticated = true
				auth_type = "manager"
			}
		}
		if is_authenticated {
			switch auth_type {
			case "student":
				return c.Redirect(http.StatusFound, "/aluno/")
			case "teacher":
				return c.Redirect(http.StatusFound, "/professor/")
			case "manager":
				return c.Redirect(http.StatusFound, "/gerenciamento/")
			default:
				return c.Render(http.StatusOK, "login_page", nil)
			}
		} else {
			return c.Render(http.StatusOK, "login_page", nil)
		}
	}
}

func DefaultMiddleware(s_type string, next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var auth_type string
		is_authenticated := false
		db = config.GetDB()
		cookie, err := c.Cookie("session")
		if err != nil { // Checks if the session cookie exists
			return c.Redirect(http.StatusFound, "/login/")
		}
		for _, s := range sessions { // Check the authentication type
			if s.SessionId == cookie.Value && s.Type == "student" {
				is_authenticated = true
				auth_type = "student"
			} else if s.SessionId == cookie.Value && s.Type == "teacher" {
				is_authenticated = true
				auth_type = "teacher"
			} else if s.SessionId == cookie.Value && s.Type == "manager" {
				is_authenticated = true
				auth_type = "manager"
			}
		}
		if is_authenticated {
			switch s_type { // Switch between the authentication types
	 		case "student":
				switch auth_type {
				case "student":
					return next(c)
				case "teacher":
					return c.Redirect(http.StatusFound, "/professor/")
				case "manager":
					return c.Redirect(http.StatusFound, "/gerenciamento/")
				default:
					return c.Redirect(http.StatusFound, "/login/")
			}
			case "teacher":
				switch auth_type {
				case "student":
					return c.Redirect(http.StatusFound, "/aluno/")
				case "teacher":
					return next(c)
				case "manager":
					return c.Redirect(http.StatusFound, "/gerenciamento/")
				default:
					return c.Redirect(http.StatusFound, "/login/")
				}
			case "manager":
				switch auth_type {
				case "student":
					return c.Redirect(http.StatusFound, "/aluno/")
				case "teacher":
					return c.Redirect(http.StatusFound, "/professor/")
				case "manager":
					return next(c)
				default:
					return c.Redirect(http.StatusFound, "/login/")
				}
			default:
				return c.Redirect(http.StatusFound, "/login/")
			}
		} else { // If no authentication, return to the login page
			return c.Redirect(http.StatusFound, "/login/")
		}
	}
}
