package handler

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/darthxd/tcc-app/config"
	"github.com/darthxd/tcc-app/models"
	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
	"gorm.io/gorm"
)

var (
	sessions []Session
	db       *gorm.DB
)

type Session struct {
	Name      string
	User      string
	Password  string
	Type      string
	ID        uint
	SessionId string
}

const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func GenerateSessionId() string {
	b := make([]byte, 8)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func LoginRender(c echo.Context) error {
	return c.Render(http.StatusOK, "login_page", echo.Map{
		"title": "Entre na sua conta",
	})
}

func LoginStudentRender(c echo.Context) error {
	return c.Render(http.StatusOK, "login_student", nil)
}

func LoginTeacherRender(c echo.Context) error {
	return c.Render(http.StatusOK, "login_teacher", nil)
}

func LoginManagerRender(c echo.Context) error {
	return c.Render(http.StatusOK, "login_manager", nil)
}

func LoginStudent(c echo.Context) error {
	// Get the DB, username and password
	db = config.GetDB()
	user := c.FormValue("user")
	password := c.FormValue("password")

	// Set the login type
	login_type := "student"

	// Generate a new session ID
	session_id := GenerateSessionId()

	// Set the session_exists variable to false
	session_exists := false

	// Create a student struct
	student := models.Student{}

	// Check if the user and password matches an database entry
	if err := db.Where("rm = ? AND password = ?", user, password).First(&student).Error; err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error": err,
		})
	}

	// Create a new session struct
	student_login := Session{
		Name:      student.Name,
		User:      user,
		Password:  password,
		Type:      login_type,
		ID:        student.ID,
		SessionId: session_id,
	}

	// Checks if the session already exists in the sessions array
	for _, s := range sessions {
		if s.Type == login_type && s.ID == student.ID {
			session_id = s.SessionId
			session_exists = true
		}
	}

	// If the session does not exist, append the new session to the sessions array
	if !session_exists {
		sessions = append(sessions, student_login)
	}

	// Create the session cookie
	cookie := new(http.Cookie)
	cookie.Name = "session"
	cookie.Value = session_id
	cookie.Expires = time.Now().AddDate(999, 999, 999)
	cookie.Path = "/"
	c.SetCookie(cookie)

	// Redirect to the student page
	c.Response().Header().Set("HX-Redirect", "/aluno")
	return nil
}

func LogOut(c echo.Context) error {
	// Check if the user is authenticated (in progress)
	_, err := c.Cookie("session")
	if err != nil {
		log.Error(err)
	}

	// Delete the session cookie
	cookie := new(http.Cookie)
	cookie.Name = "session"
	cookie.Value = ""
	cookie.MaxAge = -1
	cookie.Expires = time.Now().Add(-7*24*time.Hour)
	cookie.Path = "/"
	c.SetCookie(cookie)

	// Redirect to the login page
	c.Response().Header().Set("HX-Redirect", "/login/")
	return nil 
}
