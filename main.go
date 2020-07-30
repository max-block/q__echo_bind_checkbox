package main

import (
	"github.com/labstack/echo/v4"
	"html/template"
	"io"
	"log"
	"net/http"
)

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, _ echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

type PostForm struct {
	Name  string
	Agree bool
}

func main() {
	e := echo.New()
	e.Debug = true
	e.Renderer = &Template{
		templates: template.Must(template.ParseGlob("./templates/*.gohtml")),
	}
	e.GET("/", func(c echo.Context) error {
		return c.Render(http.StatusOK, "index.gohtml", nil)
	})
	e.POST("/post", func(c echo.Context) error {
		var form PostForm
		err := c.Bind(&form)
		if err != nil {
			return c.String(http.StatusInternalServerError, err.Error())
		}
		return c.JSON(http.StatusOK, form)

	})

	log.Fatalln(e.Start(":3000"))
}
