package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

//RetornaOlaMundo : retorna
func RetornaOlaMundo(c echo.Context) error {

	return c.String(http.StatusOK, "Olá, Mundo!")
}

func main() {
	e := echo.New()

	// Routes
	//e.GET("/", hello)
	e.Static("/css", "dist/css")
	e.Static("/img", "dist/img")
	e.Static("/js", "dist/js")

	e.File("/", "dist/index.html")
	e.File("/favicon.ico", "dist/favicon.ico")

	e.GET("/ola", RetornaOlaMundo)

	// Start server
	e.Logger.Fatal(e.Start(":8100"))
}
