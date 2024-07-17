package main

import (
	"fmt"
	"log"
	"os"
	_ "salesservice/docs"
	"salesservice/models"
	"salesservice/routes"

	"github.com/gin-gonic/gin"
	"github.com/honeycombio/otel-config-go/otelconfig"
	"github.com/joho/godotenv"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title SalesService APIs
// @version 1.0
// @description Testing Swagger APIs.
// @termsOfService http://swagger.io/terms/// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io
// @securityDefinitions.apikey Bearer
// @in header
// @name Authorization
// @Security JWT
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8080
// @BasePath /
// @schemes http https
func main() {
	// @schemes http https
	// Create a new gin instance

	r := gin.Default()

	// Load .env file and Create a new connection to the database
	err := godotenv.Load()
	// if err != nil {
	// 	log.Fatal("Error loading .env file")
	// }
	fmt.Println(err)
	config := models.Config{
		Host:           os.Getenv("DB_HOST"),
		Port:           os.Getenv("DB_PORT"),
		User:           os.Getenv("DB_USER"),
		Password:       os.Getenv("DB_PASSWORD"),
		DBName:         os.Getenv("DB_NAME"),
		SSLMode:        os.Getenv("DB_SSLMODE"),
		WorkOSClientId: os.Getenv("WORKOS_CLIENT_ID"),
	}
	otelShutdown, err := otelconfig.ConfigureOpenTelemetry()
	if err != nil {
		log.Fatalf("error setting up OTel SDK - %e", err)
	}

	defer otelShutdown()
	// Initialize DB
	models.InitDB(config)

	r.Use(otelgin.Middleware(os.Getenv("OTEL_SERVICE_NAME")))
	// Load the routes
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	routes.CustomerRoutes(r)
	// Run the server
	r.Run(":8080")
}
