package controller

import (
	"errors"
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

func ExitVehicle(plateNumber string) error {

	var vehicle *models.Vehicle

	if plateNumber == "" {
		return errors.New("plate number is required")
	}

	db, err := config.Connect()
	if err != nil {
		panic(err)
	}

	// Find the vehicle by its plate number
	result := db.Where("plate_number = ?", plateNumber).First(&vehicle)
	if result.Error != nil {
		return errors.New("vehicle not found")
	}

	// Delete the vehicle
	result = db.Delete(&vehicle)
	if result.Error != nil {
		return errors.New("vehicle can be deleted")
	}

	return nil
}
