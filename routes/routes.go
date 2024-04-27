package routes

import (
	"github.com/4anyanat/assessment-tax/handlers"
	"github.com/4anyanat/assessment-tax/database"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Router(e *echo.Echo) {
	e.Static("/", "static")
	e.GET("/health", handlers.HealthHandler)
	e.POST("/tax/calculations", handlers.Tax_Cal_Handler)
	e.POST("/tax/calculations/upload-csv", handlers.Tax_Csv_Handler)

	g := e.Group("/admin")

	g.Use(middleware.BasicAuth(func(username, password string, c echo.Context) (bool, error) {
		if username == "adminTax" && password == "admin!" {
			return true, nil
		}
		return false, nil
	}))

	g.POST("/deductions/personal", database.TaxesUpdate)
}
