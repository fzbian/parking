package models

import "time"

type Vehicle struct {
	ID          int       `json:"id" db:"id"`
	PlateNumber string    `json:"plateNumber" db:"plateNumber"`
	VehicleType string    `json:"vehicleType" db:"vehicleType"`
	Color       string    `json:"color" db:"color"`
	Zone        string    `json:"zone" db:"zone"`
	ParkingSpot int       `json:"parkingSpot" db:"parkingSpot"`
	EntryTime   time.Time `json:"entryTime" db:"entryTime"`
	ExitTime    time.Time `json:"exitTime" db:"exitTime"`
}
