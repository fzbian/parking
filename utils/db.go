package utils

import (
	"github.com/fzbian/parking/config"
	"github.com/fzbian/parking/models"
)

var (
	// Db main connection
	Db = config.Connect()
)

func CreateSpots() {
	Db.Where("in_use = 0").Delete(models.Spot{})

	for i := 0; i < 30; i++ {
		a := &models.Spot{Zone: "A"}
		if i >= 0 && i <= 4 {
			a.Type = "Vip"
		} else {
			a.Type = "Normal"
		}
		Db.Create(a)
	}
	for i := 0; i < 30; i++ {
		a := &models.Spot{Zone: "B", Type: "Normal"}
		Db.Create(a)
	}
	for i := 0; i < 40; i++ {
		a := &models.Spot{Zone: "C"}
		if i >= 0 && i <= 4 {
			a.Type = "Handicapped"
		} else if i >= 19 && i <= 21 {
			a.Type = "Emergency"
		} else if i >= 37 && i <= 40 {
			a.Type = "Provider"
		} else {
			a.Type = "Normal"
		}
		Db.Create(a)
	}
}
