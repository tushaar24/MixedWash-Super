package models

import ()

type CustomerCreationDTO struct {
	CustomerName        string `json:"customer_name"`
	CustomerPhoneNumber string `json:"customer_phone_number"`
	CustomerEmailId     string `json:"customer_email_address"`
}
