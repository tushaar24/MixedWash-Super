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
	TaskPriority   int8   `json:"task_prioritY"`
}

type DriverTaskResponseDTO struct {
	Id           string          `json:"id"`
	Customer     *models.Profile `json:"customer"`
	Address      *models.Address `json:"address"`
	Status       string          `json:"status"`
	TaskType     string          `json:"task_type"`
	DriverId     string          `json:"driver"`
	TaskPriority int8            `json:"task_priority"`
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
