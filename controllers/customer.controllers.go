// PATH: go-auth/controllers/auth.go

package controllers

import (
	"context"
	"database/sql"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	db "github.com/rovesaas/salesservice/db/sqlc"
	"github.com/rovesaas/salesservice/models"
)

func CheckOrgAdmin(role string) bool {
	return role == "orgadmin"
}

type CustomerController struct {
	db  *db.Queries
	ctx context.Context
}

func NewCustomerController(db *db.Queries, ctx context.Context) *CustomerController {
	return &CustomerController{db, ctx}
}

// CreateCustomer godoc
// @Summary Create a Customer
// @Description Create a Customer for an org.
// @Tags Customer
// @Accept json
// @Produce json
// @Param customer body models.CreateCustomerParams true "Customer Data"
// @Success 200 {object} db.Customer
// @Router /api/customer/ [post]
// @Security Bearer
func (c *CustomerController) CreateCustomer(ctx *gin.Context) {
	var payload *models.CreateCustomerParams
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	role, err := ctx.Get("role")
	if !err {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err})
		ctx.Abort()
		return
	}

	if !CheckOrgAdmin(role.(string)) {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "No Permissions"})
		ctx.Abort()
		return
	}
	tenant_id, err := ctx.Get("tenant_id")
	if !err {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err})
		ctx.Abort()
		return
	}
	args := &db.CreateCustomerParams{
		TenantID:               tenant_id.(string),
		CustomerID:             payload.CustomerID,
		CustomerName:           payload.CustomerName,
		CustomerAddress:        payload.CustomerAddress,
		CustomerCity:           payload.CustomerCity,
		CustomerState:          payload.CustomerState,
		CustomerCountry:        payload.CustomerCountry,
		CustomerTotalValue:     payload.CustomerTotalValue,
		CustomerStatus:         payload.CustomerStatus,
		CustomerAppType:        payload.CustomerAppType,
		CustomerReference:      payload.CustomerReference,
		CustomerAppSize:        payload.CustomerAppSize,
		CustomerPrimaryEmail:   payload.CustomerPrimaryEmail,
		CustomerPrimaryPhone:   payload.CustomerPrimaryPhone,
		CustomerSecondaryEmail: payload.CustomerSecondaryEmail,
		CustomerSecondaryPhone: payload.CustomerSecondaryPhone,
	}

	customer, error := c.db.CreateCustomer(ctx, *args)
	if error != nil {
		if strings.Contains(error.Error(), "duplicate key value violates unique constraint") {
			ctx.JSON(http.StatusConflict, gin.H{"error": "Customer Already exists with the Customer ID"})
			return
		}
		ctx.JSON(http.StatusBadGateway, gin.H{"error": error.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"customer": customer})

}

// GetCustomers godoc
// @Summary Get All Customers for Tenant
// @Description Get All Customers for Tenant.
// @Tags Customer
// @Accept json
// @Produce json
// @Param page query int false "Page" default(1)
// @Param limit query int false "Limit" default(10)
// @Success 200 {object} db.Customer
// @Router /api/customer/ [get]
// @Security Bearer
func (c *CustomerController) GetCustomers(ctx *gin.Context) {
	var page = ctx.DefaultQuery("page", "1")
	var limit = ctx.DefaultQuery("limit", "10")

	reqPageId, _ := strconv.Atoi(page)
	reqLimit, _ := strconv.Atoi(limit)
	offset := (reqPageId - 1) * reqLimit

	tenant_id, err := ctx.Get("tenant_id")
	if !err {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err})
		ctx.Abort()
		return
	}
	args := &db.ListCustomersParams{
		Limit:    int32(reqLimit),
		Offset:   int32(offset),
		TenantID: tenant_id.(string),
	}

	customers, error := c.db.ListCustomers(ctx, *args)

	if error != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"error": error.Error()})
		return
	}

	if customers == nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "No Customers found for the Tenant ID"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"customers": customers})
}

// GetCustomerById godoc
// @Summary Get a Customer by ID for Tenant
// @Description Get a Customer by ID for Tenant.
// @Tags Customer
// @Accept json
// @Produce json
// @Param customerid path string true "Customer ID"
// @Success 200 {object} db.Customer
// @Router /api/customer/{customerid} [get]
// @Security Bearer
func (c *CustomerController) GetCustomerById(ctx *gin.Context) {
	customerId := ctx.Param("customerid")

	tenant_id, err := ctx.Get("tenant_id")
	if !err {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err})
		ctx.Abort()
		return
	}
	args := &db.ListCustomerByIDParams{
		TenantID:     tenant_id.(string),
		CustomerUuid: uuid.MustParse(customerId),
	}

	customer, error := c.db.ListCustomerByID(ctx, *args)

	if error != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"error": error.Error()})
		return
	}

	if error == sql.ErrNoRows {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "No Customer found for the Tenant ID"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"customer": customer})
}

// DeleteCustomerById godoc
// @Summary Delete a Customer by ID for Tenant
// @Description Delete a Customer by ID for Tenant.
// @Tags Customer
// @Accept json
// @Produce json
// @Param customerid path string true "Customer ID"
// @Success 200
// @Router /api/customer/{customerid} [delete]
// @Security Bearer
func (c *CustomerController) DeleteCustomerById(ctx *gin.Context) {
	customerId := ctx.Param("customerid")

	tenant_id, err := ctx.Get("tenant_id")
	if !err {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err})
		ctx.Abort()
		return
	}
	role, err := ctx.Get("role")
	if !err {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err})
		ctx.Abort()
		return
	}
	if !CheckOrgAdmin(role.(string)) {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "No Permissions"})
		ctx.Abort()
		return
	}
	args := &db.DeleteCustomerByIDParams{
		TenantID:     tenant_id.(string),
		CustomerUuid: uuid.MustParse(customerId),
	}

	error := c.db.DeleteCustomerByID(ctx, *args)

	if error != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"error": error.Error()})
		return
	}

	if error == sql.ErrNoRows {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "No Customer found for the Tenant ID"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"Success": "Deleted the Customer"})
}
