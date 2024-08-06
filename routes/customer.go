package routes

import (
	"os"

	"github.com/rovesaas/salesservice/controllers"
	middleware "github.com/rovesaas/salesservice/middlewares"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"

	"github.com/gin-gonic/gin"
)

type CustomerRoutes struct {
	customerController controllers.CustomerController
}

func NewRouteCustomer(customerController controllers.CustomerController) CustomerRoutes {
	return CustomerRoutes{customerController}
}

func (cr *CustomerRoutes) CustomerRoute(rg *gin.RouterGroup) {
	rg.Use(otelgin.Middleware(os.Getenv("OTEL_SERVICE_NAME")))
	rg.Use(middleware.AuthenticationMiddleware())
	router := rg.Group("customer")
	router.POST("/", cr.customerController.CreateCustomer)
	router.GET("/", cr.customerController.GetCustomers)
	router.GET("/:customerid", cr.customerController.GetCustomerById)
	router.DELETE("/:customerid", cr.customerController.DeleteCustomerById)
}
