package routes

import (
	"github.com/4anyanat/assessment-tax/handlers"
	"github.com/labstack/echo/v4"
)

func Router(e *echo.Echo) {
	e.POST("/tax/calculations", handlers.Tax_Cal_Handler)
}
