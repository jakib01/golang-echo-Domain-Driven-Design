package main

import (
	"github.com/joho/godotenv"
	"os"
	"simple-rest-go-echo/Config"
	"simple-rest-go-echo/Models"
	"simple-rest-go-echo/Routes"

	"github.com/labstack/echo/v4"
)

func main() {
	// Load environment variables from .env file
	//if err := godotenv.Load(); err != nil {
	//	log.Fatalf("Error loading .env file: %v", err)
	//}

	err := godotenv.Load()
	if err != nil {
		return
	}

	// Initialize Echo instance
	e := echo.New()

	/// Connect to the database
	Config.DatabaseInit()
	defer Config.GetDB().DB()

	// Perform migrations using AutoMigrate
	db := Config.GetDB()
	err = db.AutoMigrate(&Models.Course{})
	if err != nil {
		panic(err)
	}

	// Set up Routes
	Routes.SetupRoutes(e)

	// Start the server
	serverPort := os.Getenv("SERVER_PORT")
	e.Logger.Fatal(e.Start(":" + serverPort))
}
