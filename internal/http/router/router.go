package router

import (
	"log"
	"text/template"

	"github.com/darthxd/tcc-app/internal/http/handler"
	"github.com/labstack/echo"
)

func RouterInit(port string) {
	// Set up basic server settings
	e := echo.New()
	t := &handler.Template{Templates: template.Must(template.New("").Funcs(template.FuncMap{
		"formatGrade": FormatGrade,
		"formatDate":  FormatDate,
	}).ParseGlob("internal/http/client/views/*.html"))}
	e.Renderer = t
	e.Static("/assets", "internal/http/client/static")

	routesInit(e)

	if err := e.Start(port); err != nil {
		log.Fatal(err)
	}

}
