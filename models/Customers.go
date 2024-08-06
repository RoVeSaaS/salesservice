package models

type CreateCustomerParams struct {
	CustomerID             string `json:"customer_id" "binding":"required"`
	CustomerName           string `json:"customer_name" "binding":"required"`
	CustomerAddress        string `json:"customer_address"`
	CustomerCity           string `json:"customer_city"`
	CustomerState          string `json:"customer_state"`
	CustomerCountry        string `json:"customer_country"`
	CustomerTotalValue     int64  `json:"customer_total_value"`
	CustomerStatus         string `json:"customer_status"`
	CustomerAppType        string `json:"customer_app_type"`
	CustomerReference      string `json:"customer_reference"`
	CustomerAppSize        string `json:"customer_app_size"`
	CustomerPrimaryEmail   string `json:"customer_primary_email"`
	CustomerPrimaryPhone   string `json:"customer_primary_phone"`
	CustomerSecondaryEmail string `json:"customer_secondary_email"`
	CustomerSecondaryPhone string `json:"customer_secondary_phone"`
}
