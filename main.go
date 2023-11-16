package main

import (
	"github.com/joho/godotenv"
	"github.com/labstack/gommon/log"
	"golang-echo-Domain-Driven-Design/Config"
	"golang-echo-Domain-Driven-Design/Routes"
	"os"

	"github.com/labstack/echo/v4"
)

func main() {
	// Load environment variables from .env file
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Initialize Echo instance
	e := echo.New()

	/// Connect to the database
	Config.DatabaseInit()
	defer Config.GetDB().DB()

	// Perform migrations using AutoMigrate
	db := Config.GetDB()

	//err = db.AutoMigrate(&Models.Course{})
	//if err != nil {
	//	panic(err)
	//}

	// Set up Routes
	Routes.SetupRoutes(e, db)

	// Start the server
	serverPort := os.Getenv("SERVER_PORT")
	e.Logger.Fatal(e.Start(":" + serverPort))
}
