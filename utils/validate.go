package utils

func ValidatePlateNumber(cadena string) bool {
	if len(cadena) != 6 {
		return false
	}

	for i := 0; i < 3; i++ {
		if cadena[i] < 'A' || cadena[i] > 'Z' {
			return false
		}
	}

	for i := 3; i < 6; i++ {
		if cadena[i] < '0' || cadena[i] > '9' {
			return false
		}
	}

	for i := 0; i < 3; i++ {
		if cadena[i] >= 'a' && cadena[i] <= 'z' {
			return false
		}
	}
	return true
}

func ValidateVehicleType(s string, list []string) bool {
	for _, v := range list {
		if s == v {
			return true
		}
	}
	return false
}
