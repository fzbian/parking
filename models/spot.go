package models

type Spot struct {
	ID         int    `json:"id" db:"id"`
	Name       string `json:"name" db:"name"`
	IsOccupied bool   `json:"isOccupied" db:"isOccupied"`
}
