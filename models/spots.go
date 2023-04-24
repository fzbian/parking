package models

import "fmt"

// Spot represents a parking spot.
type Spot struct {
	ID    int    `json:"id" db:"id"`
	Zone  string `json:"name" db:"zone"`
	Type  string `json:"type" db:"type"`
	InUse bool   `json:"isOccupied" db:"in_use"`
}

// GetIDString returns the ID as a string.
func (s *Spot) GetIDString() string {
	return fmt.Sprint(s.ID)
}
