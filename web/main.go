package main

import (
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.File("/", "home.html")
	e.File("/home.css", "home.css")

	e.POST("/subscribe", subscribe)
	e.POST("/unsubscribe", unsubscribe)

	e.Logger.Fatal(e.Start(":"+os.Getenv("PORT")))
}

func subscribe(c echo.Context) error {
	return c.Redirect(http.StatusFound, "/")
}

func unsubscribe(c echo.Context) error {
	return c.Redirect(http.StatusFound, "/")
}
