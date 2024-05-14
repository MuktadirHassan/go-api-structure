package main

import (
	"fitness-api/config"

	"github.com/labstack/echo/v4"
)

func main() {
	// Initialize logger
	config.InitLogger()
	// Load environment variables
	config.LoadEnv()
	// Initialize database
	config.InitializeDB()

	e := echo.New()

	// Start the server
	e.Logger.Fatal(e.Start(":8080"))
}
