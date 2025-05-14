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

func StudentInfo(c echo.Context) error {
	db = config.GetDB()
	cookie, err := c.Cookie("session")
	if err != nil {
		log.Error(err)
	}
	sessions := auth.GetSessions()
	student := schemas.Student{}
	fmt.Print(sessions) // debug
	for _, s := range sessions {
		if s.SessionId == cookie.Value {
			if err := db.Where("rm = ? AND password = ?", s.User, s.Password).First(&student).Error; err != nil {
				log.Error(err)
			}
		}
	}

	return c.Render(http.StatusOK, "student_home", echo.Map{
		"title":   fmt.Sprintf("%s - Informações", student.Name),
		"active":  "info",
		"student": student,
	})
}

func StudentMail(c echo.Context) error {
	student := schemas.Student{}
	sessions := auth.GetSessions()
	cookie, _ := c.Cookie("session")
	for _, s := range sessions {
		if s.SessionId == cookie.Value {
			if s.Type == "student" {
				if err := db.Where("rm = ? AND password = ?", s.User, s.Password).First(&student).Error; err != nil {
					log.Error(err)
				}
			}
		}
	}

	return c.Render(http.StatusOK, "student_mail", echo.Map{
		"title":   fmt.Sprintf("%s - E-mail", student.Name),
		"active":  "email",
		"student": student,
	})
}

func StudentAccount(c echo.Context) error {
	student := schemas.Student{}
	sessions := auth.GetSessions()
	cookie, _ := c.Cookie("session")
	for _, s := range sessions {
		if s.SessionId == cookie.Value {
			if s.Type == "student" {
				if err := db.Where("rm = ? AND password = ?", s.User, s.Password).First(&student).Error; err != nil {
					log.Error(err)
				}
			}
		}
	}

	return c.Render(http.StatusOK, "student_account", echo.Map{
		"title":   fmt.Sprintf("%s - Conta", student.Name),
		"active":  "conta",
		"student": student,
	})
}
