package controller

import (
	"errors"
	"github.com/fzbian/parking/models"
	"github.com/fzbian/parking/utils"
	"time"
)

func GetTotalTimeByZone(request models.Spot) (string, error) {
	var vehiclesSpots []models.VehiclesSpots
	subquery := utils.Db.Table("spots").
		Where("zone = ?", request.Zone).
		Select("id")

	result := utils.Db.Table("vehicles_spots").
		Where("spot IN (?)", subquery).
		Find(&vehiclesSpots)

	if result.Error != nil {
		return "", result.Error
	}

	var totalTime time.Duration
	for _, vehicleSpot := range vehiclesSpots {
		totalTime += vehicleSpot.ExitTime.Sub(vehicleSpot.EntryTime)
	}

	return totalTime.String(), nil
}

func GetMostUsedZone() (string, error) {
	var mostUsedZone string
	var mostUsedZoneTime time.Duration

	zone := "A"
	for i := 0; i < 3; i++ {
		zoneTime, err := GetTotalTimeByZone(models.Spot{Zone: zone})
		if err != nil {
			return "", err
		}

		zoneTimeDuration, err := time.ParseDuration(zoneTime)
		if err != nil {
			return "", err
		}

		if zoneTimeDuration > mostUsedZoneTime {
			mostUsedZone = zone
			mostUsedZoneTime = zoneTimeDuration
		}

		zone = string(int(zone[0]) + 1)
	}

	if mostUsedZone == "" {
		return "", errors.New("No hay registros aun para realizar este reporte.")
	}

	return mostUsedZone, nil
}

func GetVehiclesProvidersToPay() ([]models.Vehicles, error) {
	var vehicles []models.Vehicles
	subquery := utils.Db.Table("vehicles").
		Where("vehicle_type = 'PROVEEDOR'").
		Select("id")

	subquery2 := utils.Db.Table("vehicles_spots").
		Where("vehicle_id IN (?) AND entry_time < ?", subquery, time.Now().Add(-30*time.Minute)).
		Select("vehicle_id")

	result := utils.Db.Table("vehicles").
		Where("id IN (?)", subquery2).
		Find(&vehicles)

	if result.Error != nil {
		return nil, result.Error
	}

	return vehicles, nil
}

func GetVehiclesByZoneType(zoneType string) ([]models.Vehicles, error) {
	var vehicles []models.Vehicles
	subquery := utils.Db.Table("spots").
		Where("zone = ?", zoneType).
		Select("id")

	subquery2 := utils.Db.Table("vehicles_spots").
		Where("spot IN (?) AND entry_time > ?", subquery, time.Now().Add(-24*time.Hour)).
		Select("vehicle_id")

	result := utils.Db.Table("vehicles").
		Where("id IN (?)", subquery2).
		Find(&vehicles)

	if result.Error != nil {
		return nil, result.Error
	}

	return vehicles, nil
}

func GetUsedSpotsByNormalVehicles() (int, error) {
	var vehiclesSpots []models.VehiclesSpots
	subquery := utils.Db.Table("vehicles").
		Where("vehicle_type = 'NORMAL'").
		Select("id")

	result := utils.Db.Table("vehicles_spots").
		Where("vehicle_id IN (?)", subquery).
		Find(&vehiclesSpots)

	if result.Error != nil {
		return 0, result.Error
	}

	return len(vehiclesSpots), nil
}

func GetUsedSpotsByVIPVehicles() (int, error) {
	var vehiclesSpots []models.VehiclesSpots
	subquery := utils.Db.Table("vehicles").
		Where("vehicle_type = 'VIP'").
		Select("id")

	result := utils.Db.Table("vehicles_spots").
		Where("vehicle_id IN (?)", subquery).
		Find(&vehiclesSpots)

	if result.Error != nil {
		return 0, result.Error
	}

	return len(vehiclesSpots), nil
}

func GetUsedSpotsByHandicappedVehicles() (int, error) {
	var vehiclesSpots []models.VehiclesSpots
	subquery := utils.Db.Table("vehicles").
		Where("vehicle_type = 'DISCAPACITADO'").
		Select("id")

	result := utils.Db.Table("vehicles_spots").
		Where("vehicle_id IN (?)", subquery).
		Find(&vehiclesSpots)

	if result.Error != nil {
		return 0, result.Error
	}

	return len(vehiclesSpots), nil
}

func GetUsedSpotsByEmergencyVehicles() (int, error) {
	var vehiclesSpots []models.VehiclesSpots
	subquery := utils.Db.Table("vehicles").
		Where("vehicle_type = 'EMERGENCIA'").
		Select("id")

	result := utils.Db.Table("vehicles_spots").
		Where("vehicle_id IN (?)", subquery).
		Find(&vehiclesSpots)

	if result.Error != nil {
		return 0, result.Error
	}

	return len(vehiclesSpots), nil
}

func GetUsedSpotsByProviderVehicles() (int, error) {
	var vehiclesSpots []models.VehiclesSpots
	subquery := utils.Db.Table("vehicles").
		Where("vehicle_type = 'PROVEEDOR'").
		Select("id")

	result := utils.Db.Table("vehicles_spots").
		Where("vehicle_id IN (?)", subquery).
		Find(&vehiclesSpots)

	if result.Error != nil {
		return 0, result.Error
	}

	return len(vehiclesSpots), nil
}

func GetVehiclesInParking() (int, error) {
	var vehiclesSpots []models.VehiclesSpots
	result := utils.Db.Table("vehicles_spots").
		Where("exit_time IS NULL").
		Find(&vehiclesSpots)

	if result.Error != nil {
		return 0, result.Error
	}

	return len(vehiclesSpots), nil
}
