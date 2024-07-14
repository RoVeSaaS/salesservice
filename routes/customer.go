package routes

import (
	"salesservice/controllers"
	middleware "salesservice/middlewares"

	"github.com/gin-gonic/gin"
)

func CustomerRoutes(r *gin.Engine) {
	r.Use(middleware.AuthenticationMiddleware())
	r.POST("/customer", controllers.CreateCustomer)
	r.GET("/customer/:id", controllers.GetCustomerByID)
	r.GET("/customers", controllers.GetAllCustomers)

}
