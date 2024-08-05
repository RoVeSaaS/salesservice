// PATH: go-auth/controllers/auth.go

package controllers

import (
	"context"
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
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
		TenantID:           tenant_id.(string),
		CustomerName:       sql.NullString{String: payload.CustomerName, Valid: payload.CustomerName != ""},
		CustomerAddress:    sql.NullString{String: payload.CustomerAddress, Valid: payload.CustomerName != ""},
		CustomerCity:       sql.NullString{String: payload.CustomerCity, Valid: payload.CustomerName != ""},
		CustomerState:      sql.NullString{String: payload.CustomerState, Valid: payload.CustomerName != ""},
		CustomerCountry:    sql.NullString{String: payload.CustomerCountry, Valid: payload.CustomerName != ""},
		CustomerTotalValue: sql.NullInt64{Int64: payload.CustomerTotalValue, Valid: payload.CustomerName != ""},
		CustomerStatus:     sql.NullString{String: payload.CustomerStatus, Valid: payload.CustomerName != ""},
	}

	customer, error := c.db.CreateCustomer(ctx, *args)
	if error != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"error": error.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"customer": customer})

}
