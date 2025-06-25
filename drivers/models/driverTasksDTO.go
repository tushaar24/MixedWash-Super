package models

import ()

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

type CreateDriverTaskDTO struct {
	OrderId        *string `json:"order_id"`
	TempOrderId    *string `json:"temp_order_id"`
	CustomerId     *string `json:"customer_id"`
	TempCustomerId *string `json:"temp_customer_id"`
	DriverId       *string `json:"driver_id"`
	Status         string `json:"status"`
	TypeTask       string `json:"task_type"`
}
