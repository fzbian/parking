package controller

import (
	"errors"
	"github.com/fzbian/parking/models"
	"github.com/fzbian/parking/utils"
	"gorm.io/gorm"
)

func ParkingVehicle(request models.Vehicles) (error) {

	vehicle, err := GetOrCreateVehicle(request)
	if err != nil {
		panic(err)
	}

	IsIn := IsInParking(request.PlateNumber)
	if IsIn {
		return errors.New("vehicle is already in the parking lot")
	}

	spot := GetAvailableSpot(vehicle.VehicleType)

	if spot.ID == 0 {
		return errors.New("parking lot is full")
	}

	utils.Db.Table("vehicles_spots").Create(&models.VehiclesSpots{
		VehicleId: vehicle.Id,
		Spot:      spot.ID,
	})

	utils.Db.Table("spots").Where("id = ?", spot.ID).Update("in_use", true)
	return nil
}

func GetOrCreateVehicle(req models.Vehicles) (models.Vehicles, error) {
	validVehicleTypes := []string{"NORMAL", "VIP", "DISCAPACITADO", "EMERGENCIA", "PROVEEDOR"}
	isValidVehicleType := false
	for _, vt := range validVehicleTypes {
		if req.VehicleType == vt {
			isValidVehicleType = true
			break
		}
	}
	if !isValidVehicleType {
		return req, errors.New("invalid vehicle type")
	}

	vehicles := models.Vehicles{}

	err := utils.Db.Table("vehicles").Where("plate_number = ?", req.PlateNumber).FirstOrCreate(&req).Scan(&vehicles).Error
	if err != nil {
		return vehicles, err
	}
	return vehicles, nil
}

func GetAvailableSpot(spotType string) models.Spot {
	spot := models.Spot{}
	utils.Db.Table("spots").Where("type = ? AND in_use IS false", spotType).Find(&spot).Order("id")
	if spot.ID == 0 {
		utils.Db.Table("spots").Where("type = 'NORMAL' AND in_use IS false").Find(&spot).Order("id")
	}
	return spot
}

func IsInParking(plate string) bool {
	var vehicles models.Vehicles
	resultVehicles := utils.Db.Table("vehicles").Where("plate_number = ?", plate).First(&vehicles)
	if errors.Is(resultVehicles.Error, gorm.ErrRecordNotFound) {
		return false
	}

	var vehicleSpot models.VehiclesSpots
	resultVehiclesSpots := utils.Db.Table("vehicles_spots").Where("vehicle_id = ?", vehicles.Id).First(&vehicleSpot)
	if errors.Is(resultVehiclesSpots.Error, gorm.ErrRecordNotFound) {
		return false
	}
	return true
}

func ExitVehicle(plateNumber string) error {
	if !IsInParking(plateNumber) {
		return errors.New("vehicle is not in the parking lot")
	}

	var vehicle models.Vehicles
	utils.Db.Table("vehicles").Where("plate_number = ?", plateNumber).First(&vehicle)

	var vehicleSpot models.VehiclesSpots
	utils.Db.Table("vehicles_spots").Where("vehicle_id = ?", vehicle.Id).First(&vehicleSpot)
	utils.Db.Table("vehicles_spots").Where("vehicle_id = ?", vehicle.Id).Delete(&vehicleSpot)
	utils.Db.Table("spots").Where("id = ?", vehicleSpot.Spot).Update("in_use", false)

	var nextVehicleSpot models.VehiclesSpots
	if err := utils.Db.Table("vehicles_spots").Order("spot ASC").Where("spot > ?", vehicleSpot.Spot).First(&nextVehicleSpot).Error; err == nil {
		utils.Db.Table("vehicles_spots").Where("id = ?", nextVehicleSpot.Id).Update("spot", vehicleSpot.Spot)
	}

	return nil
}
