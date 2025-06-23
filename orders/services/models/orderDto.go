package models

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/tushaar24/mixedWash-backend/utils"
	"strings"
	"time"
)

type TempOrderDTO struct {
	ID                  uuid.UUID      `json:"id"`
	UserID              uuid.UUID      `json:"user_id"`
	AddressID           uuid.UUID      `json:"address_id"`
	ServiceID           uuid.UUID      `json:"service_id"`
	PickupDate          utils.DateOnly `json:"pickup_date"`
	PickupSlotID        uuid.UUID      `json:"pickup_slot_id"`
	DeliveryDate        utils.DateOnly `json:"delivery_date"`
	DeliverySlotID      uuid.UUID      `json:"delivery_slot_id"`
	SpecialInstructions *string        `json:"special_instructions,omitempty"`
	EstimatedWeight     *float64       `json:"estimated_weight,omitempty"`
	Status              string         `json:"status"`
	TotalAmount         *float64       `json:"total_amount,omitempty"`
	CreatedAt           time.Time      `json:"created_at"`
	UpdatedAt           time.Time      `json:"updated_at"`
	Address             *Address       `json:"addresses_temp"`
	Service             *Service       `json:"services"`
	Customer            *TempCustomer  `json:"temp_customers"`
	PickupTime          *DateAndTime   `json:"pickup_time"`
	DeliveryTime        *DateAndTime   `json:"delivery_time"`
}

func (t TempOrderDTO) ToOrderDTO() OrderDTO {
	var profile *Profile
	if t.Customer != nil {
		profile = &Profile{
			Username:     t.Customer.CustomerName,
			MobileNumber: t.Customer.CustomerPhoneNumber,
			EmailAddress: t.Customer.CustomerEmailAddress,
		}
	}

	return OrderDTO{
		ID:                  t.ID,
		UserID:              t.UserID,
		AddressID:           t.AddressID,
		ServiceID:           t.ServiceID,
		PickupDate:          t.PickupDate,
		PickupSlotID:        t.PickupSlotID,
		DeliveryDate:        t.DeliveryDate,
		DeliverySlotID:      t.DeliverySlotID,
		SpecialInstructions: t.SpecialInstructions,
		EstimatedWeight:     t.EstimatedWeight,
		Status:              t.Status,
		TotalAmount:         t.TotalAmount,
		CreatedAt:           t.CreatedAt,
		UpdatedAt:           t.UpdatedAt,
		Address:             t.Address,
		Service:             t.Service,
		Profile:             profile,
		PickupTime:          t.PickupTime,
		DeliveryTime:        t.DeliveryTime,
	}
}

type TempCustomer struct {
	CustomerName         string `json:"customer_name"`
	CustomerPhoneNumber  string `json:"customer_phone_number"`
	CustomerEmailAddress string `json:"customer_email_address"`
}

type OrderDTO struct {
	ID                  uuid.UUID      `json:"id"`
	UserID              uuid.UUID      `json:"user_id"`
	AddressID           uuid.UUID      `json:"address_id"`
	ServiceID           uuid.UUID      `json:"service_id"`
	PickupDate          utils.DateOnly `json:"pickup_date"`
	PickupSlotID        uuid.UUID      `json:"pickup_slot_id"`
	DeliveryDate        utils.DateOnly `json:"delivery_date"`
	DeliverySlotID      uuid.UUID      `json:"delivery_slot_id"`
	SpecialInstructions *string        `json:"special_instructions,omitempty"`
	EstimatedWeight     *float64       `json:"estimated_weight,omitempty"`
	Status              string         `json:"status"`
	TotalAmount         *float64       `json:"total_amount,omitempty"`
	CreatedAt           time.Time      `json:"created_at"`
	UpdatedAt           time.Time      `json:"updated_at"`
	Address             *Address       `json:"addresses"`
	Service             *Service       `json:"services"`
	Profile             *Profile       `json:"profiles"`
	PickupTime          *DateAndTime   `json:"pickup_time"`
	DeliveryTime        *DateAndTime   `json:"delivery_time"`
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
	Name string `json:"name"`
}

type Profile struct {
	Username     string `json:"username"`
	MobileNumber string `json:"mobile_number"`
	EmailAddress string `json:"email"`
}

type DateAndTime struct {
	Label string `json:"label"`
}

func (o OrderDTO) ConvertToOrderDashboardModel() OrderDashboardModel {
	// --- User name ----------------------------------------------------------
	var userName string
	if o.Profile != nil {
		userName = o.Profile.Username
	}

	var phoneNumber string
	if o.Profile != nil {
		phoneNumber = o.Profile.MobileNumber
	}

	var emailAddress string

	if o.Profile != nil {
		emailAddress = o.Profile.EmailAddress
	}

	// --- Service name -------------------------------------------------------
	var serviceName string
	if o.Service != nil {
		serviceName = o.Service.Name
	}

	// --- Address & coordinates ---------------------------------------------
	var (
		addressStr  string
		coordinates string
	)
	if o.Address != nil {
		a := o.Address

		// Build a human-readable address (skip empty parts)
		parts := []string{}
		if a.HouseBuilding != "" {
			parts = append(parts, a.HouseBuilding)
		}
		if a.AddressLine1 != "" {
			parts = append(parts, a.AddressLine1)
		}
		if a.AddressLine2 != "" {
			parts = append(parts, a.AddressLine2)
		}
		if a.Area != "" {
			parts = append(parts, a.Area)
		}
		if a.City != "" {
			parts = append(parts, a.City)
		}
		if a.State != "" {
			parts = append(parts, a.State)
		}
		if a.PostalCode != "" {
			parts = append(parts, a.PostalCode)
		}
		addressStr = strings.Join(parts, ", ")
		coordinates = fmt.Sprintf("%f,%f", a.Latitude, a.Longitude)
	}

	// --- Pick-up & delivery date-time ---------------------------------------
	var pickupDT *PickupDateTime
	pickupDate, _ := time.Parse("2006-01-02 15:04:05 -0700 MST", o.PickupDate.String())
	if o.PickupTime != nil {
		pickupDT = &PickupDateTime{
			PickUpTime: o.PickupTime.Label,
			PickUpDate: pickupDate.Format("02/01/2006"),
		}
	}

	var deliveryDT *DeliveryDateTime
	deliveryDate, _ := time.Parse("2006-01-02 15:04:05 -0700 MST", o.DeliveryDate.String())
	if o.DeliveryTime != nil {
		deliveryDT = &DeliveryDateTime{
			DeliveryTime: o.DeliveryTime.Label,
			DeliveryDate: deliveryDate.Format("02/01/2006"),
		}
	}

	// --- Assemble the dashboard model --------------------------------------
	return OrderDashboardModel{
		OrderId:          o.ID.String(),
		UserName:         userName,
		Service:          serviceName,
		MobileNumber:     phoneNumber,
		Address:          addressStr,
		Coordinates:      coordinates,
		PickupDateTime:   pickupDT,
		DeliveryDateTime: deliveryDT,
		EmailAddress:     emailAddress,
	}
}


