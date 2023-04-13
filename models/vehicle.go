package models

import "time"

type Color string

const (
	Blanco Color = "BLANCO"
	Azul   Color = "AZUL"
	Rojo   Color = "ROJO"
	Verde  Color = "VERDE"
	Dorado Color = "DORADO"
)

type Zone string

const (
	A Zone = "A"
	B Zone = "B"
	C Zone = "C"
)

type Vehicle struct {
	ID          int       `json:"id" db:"id"`
	PlateNumber string    `json:"plate_number" db:"plate_number"`
	VehicleType string    `json:"vehicle_type" db:"vehicle_type"`
	Color       Color     `json:"color" db:"color"`
	Zone        Zone      `json:"zone" db:"zone"`
	ParkingSpot int       `json:"parking_spot" db:"parking_spot"`
	EntryTime   time.Time `json:"entry_time" db:"entry_time"`
	ExitTime    time.Time `json:"exit_time" db:"exit_time"`
}
