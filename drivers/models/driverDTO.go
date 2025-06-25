package models

import ()

type DriverDTO struct {
	ID          string `json:"id"           db:"id"`
	Name        string `json:"name"         db:"name"`
	PhoneNumber string `json:"phone_number" db:"phone_number"` // stored as NUMERIC
	Salary      string `json:"salary"       db:"salary"`       // NUMERIC/REAL
}
