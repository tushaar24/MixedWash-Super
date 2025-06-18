package models

import(
)

type OrderDashboardModel struct{
	OrderId string `json:"order_id"`
	UserName string `json:"user_name"`
	Service string `json:"service"`
	MobileNumber string `json:"phone_number"`
	Address string `json:"address"`
	Coordinates string `json:"coordinates"`
	PickupDateTime *PickupDateTime `json:"pickup_date_time"`
	DeliveryDateTime *DeliveryDateTime `json:"delivery_date_time"`
}

type PickupDateTime struct{
	PickUpTime string `json:"pickup_time"`
	PickUpDate string `json:"pickup_date"`
}

type DeliveryDateTime struct{
	DeliveryTime string `json:"pickup_time"`
	DeliveryDate string `json:"pickup_date"`
}
