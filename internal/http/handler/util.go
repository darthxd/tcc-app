package handler

import (
	"io"
	"text/template"
	"time"

	"github.com/labstack/echo"
)

type Template struct {
	Templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, e echo.Context) error {
	return t.Templates.ExecuteTemplate(w, name, data)
}

type studentRequest struct {
	Name      string    `form:"name"`
	RM        string    `form:"rm"`
	Phone     string    `form:"phone"`
	Birthdate time.Time `form:"birthdate"`
	Course    string    `form:"course"`
	Grade     string    `form:"grade"`
	Email     string    `form:"email"`
	Password  string    `form:"password"`
	CPF       string    `form:"cpf"`
	RA        string    `form:"ra"`
}
