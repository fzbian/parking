package utils

import (
	"github.com/fzbian/parking/config"
	"github.com/fzbian/parking/models"
)

var (
	// Main DB connection
	Db = config.Connect()
)

func CreateSpots() {
	Db.Where("in_use = 0").Delete(models.Spot{})

	for i := 0; i < 30; i++ {
		a := &models.Spot{Zone: "A"}
		if i >= 0 && i <= 4 {
			a.Type = "VIP"
		} else {
			a.Type = "NORMAL"
		}
		Db.Create(a)
	}
	for i := 0; i < 30; i++ {
		a := &models.Spot{Zone: "B", Type: "NORMAL"}
		Db.Create(a)
	}
	for i := 0; i < 40; i++ {
		a := &models.Spot{Zone: "C"}
		if i >= 0 && i <= 4 {
			a.Type = "DISCAPACITADO"
		} else if i >= 19 && i <= 21 {
			a.Type = "EMERGENCIA"
		} else if i >= 37 && i <= 40 {
			a.Type = "PROVEEDOR"
		} else {
			a.Type = "NORMAL"
		}
		Db.Create(a)
	}
}
