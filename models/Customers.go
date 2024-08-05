package models

type CreateCustomerParams struct {
	CustomerName       string `json:"customer_name"`
	CustomerAddress    string `json:"customer_address"`
	CustomerCity       string `json:"customer_city"`
	CustomerState      string `json:"customer_state"`
	CustomerCountry    string `json:"customer_country"`
	CustomerTotalValue int64  `json:"customer_total_value"`
	CustomerStatus     string `json:"customer_status"`
}
