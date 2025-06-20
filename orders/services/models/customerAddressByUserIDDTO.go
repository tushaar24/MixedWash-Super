package models

import (
	"strings"
)

type CustomerAddressByUserIdDTO struct {
	AddressId     string  `json:"id"`
	AddressLine1  string  `json:"address_line1"`
	AddressLine2  string  `json:"address_line2"`
	City          string  `json:"city"`
	State         string  `json:"state"`
	HouseBuilding string  `json:"house_building"`
	Area          string  `json:"area"`
	PostalCode    string  `json:"postal_code"`
}

func (dto *CustomerAddressByUserIdDTO) ToModel() CustomerAddressesByUserIdModel {
	var parts []string

	if dto.HouseBuilding != "" {
		parts = append(parts, dto.HouseBuilding)
	}
	if dto.AddressLine1 != "" {
		parts = append(parts, dto.AddressLine1)
	}
	if dto.AddressLine2 != "" {
		parts = append(parts, dto.AddressLine2)
	}
	if dto.Area != "" {
		parts = append(parts, dto.Area)
	}
	if dto.City != "" {
		parts = append(parts, dto.City)
	}
	if dto.State != "" {
		parts = append(parts, dto.State)
	}
	if dto.PostalCode != "" {
		parts = append(parts, dto.PostalCode)
	}

	fullAddress := strings.Join(parts, ", ")

	return CustomerAddressesByUserIdModel{
		AddressId: dto.AddressId,
		Address:   fullAddress,
	}
}

