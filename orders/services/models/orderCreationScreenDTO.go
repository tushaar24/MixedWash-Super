package models

import ()

type OrderCreationScreenDTO struct {
	Services     []OrderCreationScreenServicesDTO             `json:"services"`
	DeliverySlot []OrderCreationScreenDeliverySlotDTO `json:"delivery_slot"`
	PickupSlot   []OrderCreationScreenPickupSlotDTO   `json:"pickup_slot"`
}

type OrderCreationScreenServicesDTO struct {
	ServiceID   string `json:"id"`
	ServiceName string `json:"name"`
}

type OrderCreationScreenDeliverySlotDTO struct {
	SlotId    string `json:"id"`
	SlotLabel string `json:"label"`
}

type OrderCreationScreenPickupSlotDTO struct {
	SlotId    string `json:"id"`
	SlotLabel string `json:"label"`
}
