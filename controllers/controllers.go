// PATH: go-auth/controllers/auth.go

package controllers

import (
	"net/http"
	"salesservice/models"

	"github.com/gin-gonic/gin"
)

func CreateCustomer(c *gin.Context) {
	var customer models.Customers

	if err := c.ShouldBindJSON(&customer); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	var existingcustomer models.Customers

	models.DB.Where("customer_name = ?", customer.CustomerName).Where("tenant_id = ?", customer.TenantID).First(&existingcustomer)
	if existingcustomer.ID != 0 {
		c.JSON(400, gin.H{"error": "Customer Already exists for this Tenant"})
		return
	}

	models.DB.Create(&customer)
	c.JSON(200, gin.H{"success": "Customer Added for the Tenant Successfully"})
}

func GetAllCustomers(c *gin.Context) {
	var customers []models.Customers
	models.DB.Find(&customers)
	c.JSON(http.StatusOK, gin.H{"customers": customers})
}

func GetCustomerByID(c *gin.Context) {
	var customer models.Customers
	if err := models.DB.Where("id = ?", c.Param("id")).First(&customer).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record Not Found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": customer})
}
