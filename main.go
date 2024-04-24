package main

import (
	"net/http"

	"github.com/4anyanat/assessment-tax/routes"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func healthHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, "OK")
}

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/health", healthHandler)

	routes.Router(e)

	e.Logger.Fatal(e.Start(":8080"))
}
