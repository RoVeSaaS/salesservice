package models

import "gorm.io/gorm"

type Customers struct {
	gorm.Model
	CustomerName string `json:"customer_name" binding:"required"`
	CustomerId   string `json:"customer_id" binding:"required" gorm:"unique"`
	Address      string `json:"address"`
	City         string `json:"city"`
	State        string `json:"state"`
	Country      string `json:"country"`
	PinCode      int    `json:"pincode"`
	IsActive     *bool  `json:"is_active" binding:"required"`
	TenantID     string `json:"tenant_id" binding:"required"`
}
