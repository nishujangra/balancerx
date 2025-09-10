package main

import (
	"fmt"
	"os"

	"github.com/labstack/echo/v4"
)

func main() {
	// Default port if none provided
	port := "9000"
	if len(os.Args) > 1 {
		port = os.Args[1]
	}

	e := echo.New()

	// Health check endpoint
	e.GET("/health", func(c echo.Context) error {
		return c.String(200, fmt.Sprintf("Ok at %s!", port))
	})

	// Root endpoint (dummy response)
	e.GET("/", func(c echo.Context) error {
		return c.String(200, fmt.Sprintf("Hello from Echo server on port %s!", port))
	})

	// Start server
	e.Logger.Fatal(e.Start(":" + port))
}
