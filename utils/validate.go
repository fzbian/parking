package utils

import (
	"errors"
	"regexp"
	"strings"
)

func ValidatePlateNumber(cadena string) error {
	cadena = strings.ToUpper(cadena)

	if len(cadena) != 6 {
		return errors.New("plate number must be 6 characters long")
	}

	if regexp.MustCompile(`^[A-Z]{3}[0-9]{3}$`).MatchString(cadena) == false {
		return errors.New("plate number must be in the format ABC123")
	}
	return nil
}

func ValidateVehicleType(s string, list []string) bool {
	for _, v := range list {
		if s == v {
			return true
		}
	}
	return false
}
