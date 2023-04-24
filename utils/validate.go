package utils

import (
	"errors"
	"regexp"
	"strings"
)

// ValidatePlateNumber validates a plate number.
// The function returns an error if the plate number is not valid.
func ValidatePlateNumber(PlateNumber string) error {

	// Convert the plate number to uppercase.
	PlateNumber = strings.ToUpper(PlateNumber)

	// Check if the plate number is 6 characters long.
	if len(PlateNumber) != 6 {
		return errors.New("Plate number must be 6 characters long")
	}

	// Check if the plate number matches the format ABC123.
	if regexp.MustCompile(`^[A-Z]{3}[0-9]{3}$`).MatchString(PlateNumber) == false {
		return errors.New("Plate number must be in the format ABC123")
	}

	// The plate number is valid.
	return nil

}

// ValidateVehicleType validates a vehicle type.
// The function returns true if the vehicle type is valid, and false otherwise.
func ValidateVehicleType(s string, list []string) bool {

	// Iterate over the list of valid vehicle types.
	for _, v := range list {
		// If the vehicle type is found in the list, return true.
		if s == v {
			return true
		}
	}

	// The vehicle type is not valid.
	return false

}
