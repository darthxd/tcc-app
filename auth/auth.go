package auth

import (
	"math/rand"
	"net/http"
	"time"

	"github.com/darthxd/tcc-app/config"
	"github.com/darthxd/tcc-app/models"
	"github.com/labstack/echo"
	"gorm.io/gorm"
)

var (
	sessions []Session
	db       *gorm.DB
)

const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

type Session struct {
	Name      string
	User      string
	Password  string
	Type      string
	ID        uint
	SessionId string
}

func GetSessions() []Session {
	return sessions
}

func GenerateSessionId() string {
	b := make([]byte, 8)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func SetCookie(c echo.Context, cookie_name, cookie_value string) error {
	cookie := new(http.Cookie)
	cookie.Name = cookie_name
	cookie.Value = cookie_value
	cookie.Expires = time.Now().AddDate(999, 999, 999)
	cookie.Path = "/"
	c.SetCookie(cookie)

	return nil
}

func DeleteCookie(c echo.Context, cookie_name string) error {
	_, err := c.Cookie(cookie_name)
	if err != nil {
		return err
	}

	cookie := new(http.Cookie)
	cookie.Name = cookie_name
	cookie.Value = ""
	cookie.MaxAge = -1
	cookie.Expires = time.Now().Add(-7*24*time.Hour)
	cookie.Path = "/"
	c.SetCookie(cookie)

	return nil
}

func AuthenticateStudent(user, password string) (Session, error) {
	db = config.GetDB()

	session_id := GenerateSessionId()
	session_exists := false

	// Check if the user and password matches an database entry
	student := models.Student{}
	if err := db.Where("rm = ? AND password = ?", user, password).First(&student).Error; err != nil {
		return Session{}, err
	}

	// Create a new session struct
	session := Session{
		Name:      student.Name,
		User:      user,
		Password:  password,
		Type:      "student",
		ID:        student.ID,
		SessionId: session_id,
	}

	// Checks if the session already exists in the sessions array
	for _, s := range sessions {
		if s.Type == "student" && s.ID == student.ID {
			session.SessionId = s.SessionId
			session_exists = true
		}
	}

	// If the session does not exist, append the new session to the sessions array
	if !session_exists {
		sessions = append(sessions, session)
	}

	return session, nil
}

func AuthenticateTeacher(user, password string) (Session, error) {
	db = config.GetDB()

	session_id := GenerateSessionId()
	session_exists := false

	// Check if the user and password matches an database entry
	teacher := models.Teacher{}
	if err := db.Where("user = ? AND password = ?", user, password).First(&teacher).Error; err != nil {
		return Session{}, err
	}

	// Create a new session struct
	session := Session{
		Name:      teacher.Name,
		User:      user,
		Password:  password,
		Type:      "teacher",
		ID:        teacher.ID,
		SessionId: session_id,
	}

	// Checks if the session already exists in the sessions array
	for _, s := range sessions {
		if s.Type == "teacher" && s.ID == teacher.ID {
			session.SessionId = s.SessionId
			session_exists = true
		}
	}

	// If the session does not exist, append the new session to the sessions array
	if !session_exists {
		sessions = append(sessions, session)
	}

	return session, nil
}

func AuthenticateManager(user, password string) (Session, error) {
	db = config.GetDB()

	session_id := GenerateSessionId()
	session_exists := false

	// Check if the user and password matches an database entry
	manager := models.Manager{}
	if err := db.Where("user = ? AND password = ?", user, password).First(&manager).Error; err != nil {
		return Session{}, err
	}

	// Create a new session struct
	session := Session{
		Name:      manager.Name,
		User:      user,
		Password:  password,
		Type:      "manager",
		ID:        manager.ID,
		SessionId: session_id,
	}

	// Checks if the session already exists in the sessions array
	for _, s := range sessions {
		if s.Type == "manager" && s.ID == manager.ID {
			session.SessionId = s.SessionId
			session_exists = true
		}
	}

	// If the session does not exist, append the new session to the sessions array
	if !session_exists {
		sessions = append(sessions, session)
	}

	return session, nil
}
