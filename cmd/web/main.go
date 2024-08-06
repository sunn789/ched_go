package main

import (
	"errors"
	"html/template"
	"io"

	"github.com/labstack/echo/v4"
	"github.com/sunn789/ched.ir/handler"
)

const port = ":4000"

type TemplateRegistry struct {
	templates map[string]*template.Template
}
type Application struct {
}

func (t *TemplateRegistry) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	tmpl, ok := t.templates[name]
	if !ok {
		err := errors.New("Template not found -> " + name)
		return err
	}
	return tmpl.ExecuteTemplate(w, "base.html", data)
}

func main() {
	e := echo.New()

	// Instantiate a template registry with an array of template set
	// Ref: https://gist.github.com/rand99/808e6e9702c00ce64803d94abff65678
	templates := make(map[string]*template.Template)
	templates["home.html"] = template.Must(template.ParseFiles("views/home.html", "views/base.html"))
	templates["about.html"] = template.Must(template.ParseFiles("views/about.html", "views/base.html"))
	e.Renderer = &TemplateRegistry{
		templates: templates,
	}

	// Route => handler
	e.GET("/", handler.HomeHandler)
	e.GET("/about", handler.AboutHandler)

	// Start the Echo server
	e.Logger.Fatal(e.Start(port))
}
