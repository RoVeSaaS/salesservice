package routes

import (
	"salesservice/controllers"

	"github.com/gin-gonic/gin"
)

func CustomerRoutes(r *gin.Engine) {
	r.POST("/customer", controllers.CreateCustomer)
	r.GET("/customer/:id", controllers.GetCustomerByID)
	r.GET("/customers", controllers.GetAllCustomers)

}
