// PATH: go-auth/controllers/auth.go

package controllers

import (
	"net/http"
	"salesservice/models"

	"github.com/gin-gonic/gin"
)

// CreateCustomer godoc
// @Summary Create a Customer
// @Description Create a Customer for an org.
// @Tags Customer
// @Accept json
// @Produce json
// @Param customer body models.Customers true "Customer Data"
// @Success 200 {object} models.Customers
// @Router /customer [post]
// @Security Bearer
func CreateCustomer(c *gin.Context) {
	var customer models.Customers

	if err := c.ShouldBindJSON(&customer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var existingcustomer models.Customers

	models.DB.Where("customer_id = ?", customer.CustomerId).Where("tenant_id = ?", customer.TenantID).First(&existingcustomer)
	if existingcustomer.ID != 0 {
		c.JSON(http.StatusConflict, gin.H{"error": "Customer Already exists for this Tenant"})
		return
	}

	models.DB.Create(&customer)
	c.JSON(http.StatusOK, gin.H{"success": "Customer Added for the Tenant Successfully"})
}

// GetCustomers godoc
// @Summary Get Customers
// @Description Retreive all the customers for an org.
// @Tags Customer
// @Accept json
// @Produce json
// @Success 200 {object} models.Customers
// @Router /customers [get]
// @Security Bearer
func GetAllCustomers(c *gin.Context) {
	var customers []models.Customers
	models.DB.Find(&customers)
	c.JSON(http.StatusOK, gin.H{"customers": customers})
}

// GetCustomer godoc
// @Summary Get Customer
// @Description Retreive a customer by ID for an org.
// @Tags Customer
// @Accept json
// @Produce json
// @Param customer_id path int true "CustomerID"
// @Success 200 {object} models.Customers
// @Router /customer/{customer_id} [get]
// @Security Bearer
func GetCustomerByID(c *gin.Context) {
	var customer models.Customers
	if err := models.DB.Where("id = ?", c.Param("id")).First(&customer).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record Not Found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": customer})
}
