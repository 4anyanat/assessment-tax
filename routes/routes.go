package routes

import (
	"os"
	"fmt"

	"github.com/4anyanat/assessment-tax/handlers"
	"github.com/4anyanat/assessment-tax/database"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Router(e *echo.Echo) {
	e.Static("/", "static")

	//  Get health status
	e.GET("/health", handlers.HealthHandler)

	// Get calculated tax information by inputs
	e.POST("/tax/calculations", handlers.Tax_Cal_Handler)

	// Get calculated tax information by uploaded .csv file inputs
	e.POST("/tax/calculations/upload-csv", handlers.Tax_Csv_Handler)


	// Group the restricted routes that only authenticated admin with authorized credentials can access
	g := e.Group("/admin")

	// Get environment variables of ADMIN_USERNAME and ADMIN_PASSWORD
	adminUsername := os.Getenv("ADMIN_USERNAME")
	adminPassword := os.Getenv("ADMIN_PASSWORD")

	if adminUsername != "" && adminPassword != ""{
		fmt.Println("Credentials are available")
	} else {
		fmt.Println("Credentials are not available")
	}
	g.Use(middleware.BasicAuth(func(username, password string, c echo.Context) (bool, error) {
		if username == adminUsername && password == adminPassword {
			return true, nil
		}
		return false, nil
	}))

	// Update the personal deduction amount in the database by inputs
	g.POST("/deductions/personal", database.TaxesUpdate)

	// Update the k-receipt allowance amount in the database by inputs
	g.POST("/deductions/k-receipt", database.AllowanceUpdate)
}
