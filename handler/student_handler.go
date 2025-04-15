package handler

import (
	"fmt"
	"net/http"

	"github.com/darthxd/tcc-app/config"
	"github.com/darthxd/tcc-app/models"
	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
)

func StudentInfo(c echo.Context) error {
	// Get the database
	db = config.GetDB()

	// Get the session cookie and checks if it exists
	cookie, err := c.Cookie("session")
	if err != nil {
		log.Error(err)
	}

	// Start a new student struct
	student := models.Student{}
	fmt.Print(sessions) // debug

	// Checks if the session student matches an database entry
	for _,s := range(sessions){
		if s.SessionId == cookie.Value {
			if err := db.Where("rm = ? AND password = ?", s.User, s.Password).First(&student).Error; err != nil {
				log.Error(err)
			}
		}
	}	

	// Render the page passing the student object
	return c.Render(http.StatusOK,"student_home", echo.Map{
		"title":"Informações",
		"active":"info",
		"student":student,
	})
}

func StudentMail(c echo.Context) error {
	return c.Render(http.StatusOK, "student_mail", echo.Map{
		"title":"E-mail",
		"active":"email",
	})
}
