package controller

import (
	"errors"
	"fmt"
	"time"

	"github.com/fzbian/parking/enums"
	"github.com/fzbian/parking/models"
	"github.com/fzbian/parking/utils"
	"gorm.io/gorm"
)

func ParkingVehicle(request models.Vehicles) (string, error) {

	if err := utils.ValidatePlateNumber(request.PlateNumber); err != nil {
		return "", err
	}

	VehiclesTypes := enums.GetVehicleTypes()
	ValidVehicleType := utils.ValidateVehicleType(request.VehicleType, VehiclesTypes)
	if !ValidVehicleType {
		return "", errors.New("El tipo de vehiculo no es valido")
	}

	vehicle, err := GetOrCreateVehicle(request)
	if err != nil {
		panic(err)
	}

	vehicleSpot := GetVehicleInSpot(vehicle.Id)
	if vehicleSpot.Id != 0 {
		return "", errors.New("El vehiculo ya se encuentra estacionado")
	}

	vehicles := GetVehicles(vehicle.Id)
	if request.VehicleType != vehicles.VehicleType {
		return "", errors.New("Este vehiculo ya se encuentra en la base de datos y el tipo de vehiculo no coinciden.")
	}

	spot := GetAvailableSpot(vehicle.VehicleType)

	if spot.ID == 0 {
		return "", errors.New("El parqueadero esta lleno")
	}

	utils.Db.Table("vehicles_spots").
		Create(&models.VehiclesSpots{
			VehicleId: vehicle.Id,
			Spot:      spot.ID,
			EntryTime: time.Now(),
		})

	utils.Db.Table("spots").
		Where("id = ?", spot.ID).
		Update("in_use", true)

	SuccessfullyMessage := fmt.Sprintf("El vehiculo %s ha sido estacionado en la bahia %d ubicada en la zona %s.",
		vehicle.PlateNumber, spot.ID, spot.Zone)

	CreateEntryTicket(vehicle.PlateNumber, vehicle.VehicleType, spot.Zone, spot.ID)

	return SuccessfullyMessage, nil
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

// Esta funcion hace algo
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

func ExitVehicle(plateNumber string) (string, error) {
	if err := utils.ValidatePlateNumber(plateNumber); err != nil {
		return "", err
	}

	vehicle := models.Vehicles{}
	resultVehicles := utils.Db.Table("vehicles").
		Where("plate_number = ?", plateNumber).
		First(&vehicle)
	if errors.Is(resultVehicles.Error, gorm.ErrRecordNotFound) {
		return "", errors.New("El vehiculo no existe")
	}

	vehicleSpot := GetVehicleInSpot(vehicle.Id)
	if vehicleSpot.Id == 0 {
		return "", errors.New("El vehiculo no se encuentra estacionado")
	}

	Zone := GetZoneFromVehicleSpot(vehicle.Id)

	utils.Db.Table("vehicles_spots").
		Where("vehicle_id = ? AND exit_time IS NULL", vehicleSpot.VehicleId).
		Update("exit_time", time.Now())

	utils.Db.Table("spots").
		Where("id = ?", vehicleSpot.Spot).
		Update("in_use", false)

	EntryTime := vehicleSpot.EntryTime.Format("15:04:05")
	CreateExitTicket(vehicle.PlateNumber, vehicle.VehicleType, Zone, EntryTime, vehicleSpot.Spot)

	return "El vehiculo ha salido correctamente", nil
}

func ExitAllVehicles() (string, error) {
	vehiclesSpots := GetVehiclesSpots()

	for _, vehicleSpot := range vehiclesSpots {
		vehicle := GetVehicles(vehicleSpot.VehicleId)
		Zone := GetZoneFromVehicleSpot(vehicle.Id)

		utils.Db.Table("vehicles_spots").
			Where("vehicle_id = ? AND exit_time IS NULL", vehicleSpot.VehicleId).
			Update("exit_time", time.Now())

		utils.Db.Table("spots").
			Where("id = ?", vehicleSpot.Spot).
			Update("in_use", false)

		EntryTime := vehicleSpot.EntryTime.Format("15:04:05")
		CreateExitTicket(vehicle.PlateNumber, vehicle.VehicleType, Zone, EntryTime, vehicleSpot.Spot)
	}

	return "Todos los vehiculos han salido correctamente", nil
}

// Generar reporte tiempo total de uso por zona, para establecer cuál fue la zona más utilizada
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

	return mostUsedZone, nil
}
