package main

import (
	"context"
	"os"
	"os/signal"
	"net/http"
	"time"
	
	"github.com/4anyanat/assessment-tax/routes"
	"github.com/4anyanat/assessment-tax/database"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Signal interface {
	String() string
	Signal()
}

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	routes.Router(e)
	database.DatabaseInit()

	// Get port number from environment variable PORT
	portNo := os.Getenv("PORT")
	port := ":" + portNo

	go func() {
		e.Logger.Fatal(e.Start(port))
		if err := e.Start(port); err != nil && err != http.ErrServerClosed { 
			e.Logger.Fatal("shutting down the server")
		}
	}()

	shutdown := make(chan os.Signal, 1)

	signal.Notify(shutdown, os.Interrupt)

	<-shutdown
	e.Logger.Fatal("Shutting down . . .")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
}
