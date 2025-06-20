package models

type AddAddressAdminDTO struct {
	UserID        string   `json:"user_id"`
	AddressLine1  string   `json:"address_line1"`
	AddressLine2  *string  `json:"address_line2"`
	City          string   `json:"city"`
	State         string   `json:"state"`
	PostalCode    string   `json:"postal_code"`
	IsDefault     *bool    `json:"is_default"`
	HouseBuilding *string  `json:"house_building"`
	Area          *string  `json:"area"`
	Latitude      *float64 `json:"latitude"`
	Longitude     *float64 `json:"longitude"`
}
