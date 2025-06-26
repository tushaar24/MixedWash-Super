package models

import "github.com/tushaar24/mixedWash-backend/orders/services/models"

type DriverTasksDTO struct {
	Id             string `json:"id"`
	OrderId        string `json:"order_id"`
	TempOrderId    string `json:"temp_order_id"`
	CustomerId     string `json:"customer_id"`
	TempCustomerId string `json:"temp_customer_id"`
	DriverId       string `json:"driver_id"`
	Status         string `json:"status"`
	TypeTask       string `json:"task_type"`
}

type DriverTaskResponseDTO struct {
	Id       string          `json:"id"`
	Customer *models.Profile `json:"customer"`
	Address  *models.Address `json:"address"`
	Status   string          `json:"status"`
	TaskType string          `json:"task_type"`
	DriverId string          `json:"driver"`
}

type DriverTaskCustomerResponseDTO struct {
	Username     string `json:"username"`
	MobileNumber string `json:"mobile_number"`
	EmailAddress string `json:"email"`
}

type DriverTaskAddressResponseDTO struct {
	AddressLine1  string  `json:"address_line1"`
	AddressLine2  string  `json:"address_line2"`
	City          string  `json:"city"`
	State         string  `json:"state"`
	HouseBuilding string  `json:"house_building"`
	Area          string  `json:"area"`
	PostalCode    string  `json:"postal_code"`
	Latitude      float64 `json:"latitude"`
	Longitude     float64 `json:"longitude"`
}

type CreateDriverTaskDTO struct {
	OrderId        *string `json:"order_id"`
	TempOrderId    *string `json:"temp_order_id"`
	CustomerId     *string `json:"customer_id"`
	TempCustomerId *string `json:"temp_customer_id"`
	DriverId       *string `json:"driver_id"`
	Status         string  `json:"status"`
	TypeTask       string  `json:"task_type"`
}
