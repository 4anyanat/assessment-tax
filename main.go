package main

import (
	"github.com/4anyanat/assessment-tax/routes"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)


func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	

	routes.Router(e)

	e.Logger.Fatal(e.Start(":8080"))
}
