package main

import (
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// main function starts the server
func main() {
	// Setup the application and run it
	e := setupServer()
	httpPort := getPort()
	e.Logger.Fatal(e.Start(":" + httpPort))
}

// setupServer initializes the Echo server with routes
func setupServer() *echo.Echo {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", func(c echo.Context) error {
		return c.HTML(http.StatusOK, "Hello, Docker! <3")
	})

	e.GET("/health", func(c echo.Context) error {
		return c.JSON(http.StatusOK, struct{ Status string }{Status: "OK"})
	})

	return e
}

// getPort retrieves the port from environment variables or defaults to 8080
func getPort() string {
	httpPort := os.Getenv("PORT")
	if httpPort == "" {
		httpPort = "8080"
	}
	return httpPort
}

// IntMin returns the minimum of two integers
func IntMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}
