package controller

import (
	"github.com/fzbian/parking/config"
	"github.com/fzbian/parking/models"
)

func InsertVehicle(vehicle *models.Vehicle) error {
	db, err := config.Connect()
	if err != nil {
		return err
	}

	result := db.Create(vehicle)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
