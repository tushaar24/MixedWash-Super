package models

import "github.com/google/uuid"

type OrderCreationDTO struct {
	UserID          uuid.UUID `json:"user_id"`
	AddressID       uuid.UUID `json:"address_id"`
	ServiceID       uuid.UUID `json:"service_id"`
	PickupDate      string    `json:"pickup_date"`
	PickupSlotID    uuid.UUID `json:"pickup_slot_id"`
	DeliveryDate    string    `json:"delivery_date"`
	DeliverySlotID  uuid.UUID `json:"delivery_slot_id"`
	SpecialInstr    string    `json:"special_instructions"`
	EstimatedWeight float64   `json:"estimated_weight"`
	Status          string    `json:"status"`
	TotalAmount     float64   `json:"total_amount"`
}
