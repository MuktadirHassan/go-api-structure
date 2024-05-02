package main

import (
	"fitness-api/cmd/handlers"
	"fitness-api/cmd/storage"

	"github.com/labstack/echo/v4"
)

func main() {

	e := echo.New()
	// Define the routes
	e.GET("/", handlers.Home)
	// Initialize the database
	storage.InitDB()
	// Start the server
	e.Logger.Fatal(e.Start(":8080"))
}
