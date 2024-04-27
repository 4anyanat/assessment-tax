package routes

import (
	"github.com/4anyanat/assessment-tax/handlers"
	"github.com/labstack/echo/v4"
)

func Router(e *echo.Echo) {
	e.Static("/", "static")
	e.GET("/health", handlers.HealthHandler)
	e.POST("/tax/calculations", handlers.Tax_Cal_Handler)
	e.POST("/tax/calculations/upload-csv", handlers.Tax_Csv_Handler)
}
