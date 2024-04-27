package main

import (
	"os"
	"github.com/4anyanat/assessment-tax/routes"
	"github.com/4anyanat/assessment-tax/database"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)


func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	routes.Router(e)
	database.DatabaseInit()

	// Get port number from environment variable PORT
	port := os.Getenv("PORT")

	e.Logger.Fatal(e.Start(port))
}
