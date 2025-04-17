package router

import (
	"text/template"
	"log"

	"github.com/darthxd/tcc-app/handler"
	"github.com/labstack/echo"
)

func RouterInit(port string) {
	// Set up basic server settings
	e := echo.New()
	t := &handler.Template{Templates: template.Must(template.ParseGlob("public/views/*.html"))}
	e.Renderer = t
	e.Static("/assets", "public/static")

	routesInit(e)

	if err := e.Start(port); err != nil {
		log.Fatal(err)
	}

}
