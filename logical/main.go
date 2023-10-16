package main

import (
	"github.com/labstack/echo/v4"
	"html/template"
	"io"
	"net/http"
)

type Renderer struct {
	Template *template.Template
	Debug    bool
	Location string
}

func main() {
	e := echo.New()

	e.Renderer = NewRenderer("template/pages/index.html", true)

	e.GET("/index", func(c echo.Context) error {
		return c.Render(http.StatusOK, "index.html", nil)
	})

	//e.Static("/", "")
	e.Static("/", "template")

	e.Logger.Fatal(e.Start(":9000"))
}

func NewRenderer(location string, debug bool) *Renderer {
	tmplt := new(Renderer)
	tmplt.Location = location
	tmplt.Debug = debug

	tmplt.ReloadTemplates()

	return tmplt
}

func (r *Renderer) ReloadTemplates() {
	r.Template = template.Must(template.ParseGlob(r.Location))
}

func (r *Renderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	if r.Debug {
		r.ReloadTemplates()
	}
	return r.Template.ExecuteTemplate(w, name, data)
}
