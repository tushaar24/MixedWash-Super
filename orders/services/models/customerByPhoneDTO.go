package models

type CustomerByPhoneDTO struct {
	CustomerId   string `json:"id"`
	CustomerName string `json:"username"`
}

type TempCustomerByPhoneDTO struct {
	CustomerID   string `json:"id"`
	CustomerName string `json:"customer_name"`
}
