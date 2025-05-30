package handler

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/darthxd/tcc-app/internal/config"
	"github.com/darthxd/tcc-app/internal/http/auth"
	"github.com/darthxd/tcc-app/internal/schemas"
	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
)

func (student studentRequest) genPassword() {
	password := fmt.Sprintf("%s%s", strings.Split(student.Name, " ")[0], student.RM)
	student.Password = password
}

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
	var students []schemas.Student
	manager := AuthenticateManager(c)
	result := db.Find(&students)
	if result.Error != nil {
		log.Print(result.Error)
	}
	return c.Render(http.StatusOK, "manager_home", echo.Map{
		"title":    "Alunos cadastrados",
		"active":   "alunos-cadastrados",
		"manager":  manager,
		"students": students,
	})
}

func ManagerNewStudent(c echo.Context) error {
	manager := AuthenticateManager(c)
	if c.Request().Method == http.MethodPost {

		request := studentRequest{}

		c.Bind(&request)

		request.genPassword()

		log.Print(request)

		// student := schemas.Student{
		// 	Name:      request.Name,
		// 	RM:        request.RM,
		// 	Phone:     request.Phone,
		// 	Birthdate: request.Birthdate,
		// 	Course:    request.Course,
		// 	Grade:     request.Grade,
		// 	Email:     request.Email,
		// 	Password:  request.Password,
		// 	CPF:       request.CPF,
		// 	RA:        request.RA,
		// }

		// if err := db.Create(&student).Error; err != nil {
		// 	log.Fatal(err)
		// }

		return c.Redirect(200, "/")

	} else {
		return c.Render(http.StatusOK, "manager_newstudent", echo.Map{
			"title":   "Cadastrar aluno",
			"active":  "cadastrar-aluno",
			"manager": manager,
		})
	}
}
