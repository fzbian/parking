package controller

import (
	"github.com/fzbian/parking/models"
	"github.com/fzbian/parking/utils"
	"time"
)

func GetVehicleInSpot(idVehicle int) models.VehiclesSpots {
	var vehicleSpot models.VehiclesSpots
	utils.Db.Table("vehicles_spots").
		Where("vehicle_id = ? AND exit_time IS NULL", idVehicle).
		First(&vehicleSpot)

	return vehicleSpot
}

func GetVehicles(idVehicle int) models.Vehicles {
	var vehicles models.Vehicles
	utils.Db.Table("vehicles").
		Where("id = ?", idVehicle).
		First(&vehicles)

	return vehicles
}

func GetVehiclesSpots() []models.VehiclesSpots {
	var vehiclesSpots []models.VehiclesSpots
	utils.Db.Table("vehicles_spots").
		Find(&vehiclesSpots)

	return vehiclesSpots
}

func GetAvailableSpot(spotType string) models.Spot {
	spot := models.Spot{}
	utils.Db.Table("spots").
		Where("type = ? AND in_use IS false", spotType).
		Find(&spot).
		Order("id")
	if spot.ID == 0 {
		utils.Db.Table("spots").
			Where("type = 'NORMAL' AND in_use IS false").
			Find(&spot).
			Order("id")
	}
	return spot
}

func GetZoneFromVehicleSpot(idVehicle int) string {
	var vehicleSpot models.VehiclesSpots
	utils.Db.Table("vehicles_spots").
		Where("vehicle_id = ? AND exit_time IS NULL", idVehicle).
		First(&vehicleSpot)

	var spot models.Spot
	utils.Db.Table("spots").
		Where("id = ?", vehicleSpot.Spot).
		First(&spot)

	return spot.Zone
}

func GetVehiclePlateNumberBySpotId(spotId int) string {
	var vehicleSpot models.VehiclesSpots
	utils.Db.Table("vehicles_spots").
		Where("spot = ? AND exit_time IS NULL", spotId).
		Find(&vehicleSpot)

	var vehicle models.Vehicles
	utils.Db.Table("vehicles").
		Where("id = ?", vehicleSpot.VehicleId).
		Find(&vehicle)

	return vehicle.PlateNumber
}

func GetVehiclesBySpotType(vehicleType string) ([]models.VehiclesSpots, error) {
	var vehiclesSpots []models.VehiclesSpots
	subquery := utils.Db.Table("vehicles").
		Where("vehicle_type = ?", vehicleType).
		Select("id")

	result := utils.Db.Table("vehicles_spots").
		Where("vehicle_id IN (?) AND entry_time > ?", subquery, time.Now().Add(-24*time.Hour)).
		Find(&vehiclesSpots)

	if result.Error != nil {
		return nil, result.Error
	}

	return vehiclesSpots, nil
}

func GetOrCreateVehicle(req models.Vehicles) (models.Vehicles, error) {

	vehicles := models.Vehicles{}

	err := utils.Db.Table("vehicles").
		Where("plate_number = ?", req.PlateNumber).
		FirstOrCreate(&req).
		Scan(&vehicles).Error
	if err != nil {
		return vehicles, err
	}
	return vehicles, nil
}
