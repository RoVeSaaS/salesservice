// PATH: go-auth/main.go

package main

import (
	"log"
	"os"
	middleware "salesservice/middlewares"
	"salesservice/models"
	"salesservice/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Create a new gin instance
	r := gin.Default()

	// Load .env file and Create a new connection to the database
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	config := models.Config{
		Host:           os.Getenv("DB_HOST"),
		Port:           os.Getenv("DB_PORT"),
		User:           os.Getenv("DB_USER"),
		Password:       os.Getenv("DB_PASSWORD"),
		DBName:         os.Getenv("DB_NAME"),
		SSLMode:        os.Getenv("DB_SSLMODE"),
		WorkOSClientId: os.Getenv("WORKOS_CLIENT_ID"),
	}

	// Initialize DB
	models.InitDB(config)
	r.Use(middleware.AuthenticationMiddleware())
	// Load the routes
	routes.CustomerRoutes(r)

	// Run the server
	r.Run(":8080")
}
