package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/rovesaas/salesservice/controllers"
	_ "github.com/rovesaas/salesservice/docs"

	dbCon "github.com/rovesaas/salesservice/db/sqlc"
	"github.com/rovesaas/salesservice/routes"

	"github.com/gin-gonic/gin"
	"github.com/honeycombio/otel-config-go/otelconfig"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "github.com/lib/pq"
)

// @title SalesService APIs
// @version 1.0
// @description Sales Service APIs.
// @contact.name RoVe SaaS Support
// @contact.email rovesaas@gmail.com
// @securityDefinitions.apikey Bearer
// @in header
// @name Authorization
// @Security JWT
// @license.name Mozilla Public License Version 2.0
// @license.url https://www.mozilla.org/en-US/MPL/2.0/
// @host localhost:8080
// @BasePath /
// @schemes http https

var (
	server *gin.Engine
	db     *dbCon.Queries
	ctx    context.Context

	CustomerController controllers.CustomerController
	CustomerRoutes     routes.CustomerRoutes
)

func init() {
	ctx = context.TODO()
	_ = godotenv.Load()
	connection, err := sql.Open(os.Getenv("DB_DRIVER"), os.Getenv("DB_SOURCE"))
	if err != nil {
		log.Fatal("Couldn't Connect to the DB", err)
	}

	db = dbCon.New(connection)
	fmt.Println("Connected Successfully to the Datastore")
	CustomerController = *controllers.NewCustomerController(db, ctx)
	CustomerRoutes = routes.NewRouteCustomer(CustomerController)

	server = gin.Default()
}
func main() {
	// @schemes http https
	otelShutdown, err := otelconfig.ConfigureOpenTelemetry(
		otelconfig.WithMetricsEnabled(false),
	)
	if err != nil {
		log.Fatalf("error setting up OTel SDK - %e", err)
	}

	defer otelShutdown()

	server.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	router := server.Group("/api")
	CustomerRoutes.CustomerRoute(router)

	// Run the server
	server.Run(":8080")
}
