package models

import (
	"errors"
	"regexp"
	"strings"
	"time"
)

type Vehicles struct {
	Id          int    `json:"id" db:"id"`
	PlateNumber string `json:"plate_number" db:"plate_number"`
	VehicleType string `json:"vehicle_type" db:"vehicle_type"`
}

type VehiclesSpots struct {
	Id        int        `json:"id" db:"id"`
	VehicleId int        `json:"vehicle_id" db:"vehicle_id"`
	Spot      int        `json:"spot" db:"spot"`
	EntryTime time.Time  `json:"entry_time" db:"entry_time"`
	ExitTime  *time.Time `json:"exit_time" db:"exit_time"`
}

func (v *Vehicles) ValidatePlateNumber() error {
	v.PlateNumber = strings.ToUpper(v.PlateNumber)

	if len(v.PlateNumber) != 6 {
		return errors.New("plate number must be 6 characters long")
	}

	if regexp.MustCompile(`^[A-Z]{3}[0-9]{3}$`).MatchString(v.PlateNumber) == false {
		return errors.New("plate number must be in the format ABC123")
	}

	return nil
}
