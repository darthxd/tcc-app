package handler

import (
	"io"
	"text/template"

	"github.com/labstack/echo"
)

type Template struct {
	Templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, e echo.Context) error {
	return t.Templates.ExecuteTemplate(w, name, data)
}
