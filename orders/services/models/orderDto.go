package models

import (
	"time"
	"github.com/google/uuid"
	"github.com/tushaar24/mixedWash-backend/utils"
)

type OrderDTO struct {
	ID                  uuid.UUID  `json:"id"`
	UserID              uuid.UUID  `json:"user_id"`
	AddressID           uuid.UUID  `json:"address_id"`
	ServiceID           uuid.UUID  `json:"service_id"`
	PickupDate          utils.DateOnly   `json:"pickup_date"`
	PickupSlotID        uuid.UUID  `json:"pickup_slot_id"`
	DeliveryDate        utils.DateOnly  `json:"delivery_date"`
	DeliverySlotID      uuid.UUID  `json:"delivery_slot_id"`
	SpecialInstructions *string    `json:"special_instructions,omitempty"`
	EstimatedWeight     *float64   `json:"estimated_weight,omitempty"`
	Status              string     `json:"status"`
	TotalAmount         *float64   `json:"total_amount,omitempty"`
	CreatedAt           time.Time  `json:"created_at"`
	UpdatedAt           time.Time  `json:"updated_at"`
	Address      *Address   `json:"addresses"`
	Service      *Service   `json:"services"`
	Profile      *Profile   `json:"profiles"`
}

type Address struct {
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

type Service struct {
	Name        string `json:"name"`
}

type Profile struct {
	Username     string `json:"username"`
	MobileNumber string `json:"mobile_number"`
}



