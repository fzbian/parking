package models

import (
	"time"
)

// Vehicles represents a vehicle.
type Vehicles struct {
	Id          int    `json:"id" db:"id"`
	PlateNumber string `json:"plate_number" db:"plate_number"`
	VehicleType string `json:"vehicle_type" db:"vehicle_type"`
}

// VehiclesSpots represents a vehicle spot.
type VehiclesSpots struct {
	Id        int        `json:"id" db:"id"`
	VehicleId int        `json:"vehicle_id" db:"vehicle_id"`
	Spot      int        `json:"spot" db:"spot"`
	EntryTime time.Time  `json:"entry_time" db:"entry_time"`
	ExitTime  *time.Time `json:"exit_time" db:"exit_time"`
}
