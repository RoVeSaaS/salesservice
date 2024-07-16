// PATH: go-auth/controllers/auth.go

package controllers

import (
	"fmt"
	"net/http"
	"salesservice/models"

	"github.com/gin-gonic/gin"
)

type TenantCheck struct {
	TenantID string `json:"tenant_id" binding:"required"`
}

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
	ClaimTenantId, _ := c.Get("tenant_id")
	fmt.Println(ClaimTenantId)
	if ClaimTenantId != customer.TenantID {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Passed Org ID is not matching with Token org ID"})
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
// @Param TenantCheck body TenantCheck true "TenantCheck"
// @Success 200 {object} models.Customers
// @Router /customers [get]
// @Security Bearer
func GetAllCustomers(c *gin.Context) {
	var customers []models.Customers
	var TenantCheck TenantCheck
	if err := c.ShouldBindJSON(&TenantCheck); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ClaimTenantId, _ := c.Get("tenant_id")
	fmt.Println(ClaimTenantId)
	if ClaimTenantId != TenantCheck.TenantID {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Passed Org ID is not matching with Token org ID"})
		return
	}
	if err := models.DB.Where("tenant_id = ?", ClaimTenantId).Find(&customers).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Not a valid org"})
	}
	c.JSON(http.StatusOK, gin.H{"customers": customers})
}

// GetCustomer godoc
// @Summary Get Customer
// @Description Retreive a customer by ID for an org.
// @Tags Customer
// @Accept json
// @Produce json
// @Param customer_id path int true "CustomerID"
// @Param TenantCheck body TenantCheck true "TenantCheck"
// @Success 200 {object} models.Customers
// @Router /customer/{customer_id} [get]
// @Security Bearer
func GetCustomerByID(c *gin.Context) {
	var customer models.Customers
	var TenantCheck TenantCheck
	if err := c.ShouldBindJSON(&TenantCheck); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ClaimTenantId, _ := c.Get("tenant_id")
	fmt.Println(ClaimTenantId)
	if ClaimTenantId != TenantCheck.TenantID {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Passed Org ID is not matching with Token org ID"})
		return
	}
	if err := models.DB.Where("id = ?", c.Param("id")).Where("tenant_id = ?", ClaimTenantId).First(&customer).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record Not Found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": customer})
}
