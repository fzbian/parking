package models

import (
	"time"
)

type Payment struct {
	ID          int       `json:"id" db:"id"`
	VehicleID   int       `json:"vehicleId" db:"vehicleId"`
	PaymentType string    `json:"paymentType" db:"paymentType"`
	Amount      float64   `json:"amount" db:"amount"`
	CreatedAt   time.Time `json:"createdAt" db:"createdAt"`
}
