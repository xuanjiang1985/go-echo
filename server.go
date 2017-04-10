package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
	pongo2 "github.com/mayowa/echo-pongo2"
	"net/http"
)

func main() {

	e := echo.New()
	r, _ := pongo2.NewRenderer("./template")
	e.SetRenderer(r)
	e.Static("/", "assets")

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.POST("/save", save)
	e.GET("/hello", hello)

	e.Run(standard.New(":8080"))
}

func save(c echo.Context) error {
	name := c.FormValue("name")
	email := c.FormValue("email")
	return c.String(http.StatusOK, name+email)
}

func hello(c echo.Context) error {
	data := map[string]interface{}{"World": "zhou gang 123"}
	return c.Render(http.StatusOK, "hello.html", data)
}
