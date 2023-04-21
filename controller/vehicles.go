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

	if err := request.ValidatePlateNumber(); err != nil {
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

	utils.Db.Table("vehicles_spots").Create(&models.VehiclesSpots{
		VehicleId: vehicle.Id,
		Spot:      spot.ID,
		EntryTime: time.Now(),
	})

	utils.Db.Table("spots").Where("id = ?", spot.ID).Update("in_use", true)

	SuccessfullyMessage := fmt.Sprintf("El vehiculo %s ha sido estacionado en la bahia %d ubicada en la zona %s.",
		vehicle.PlateNumber, spot.ID, spot.Zone)

	return SuccessfullyMessage, nil
}

func GetOrCreateVehicle(req models.Vehicles) (models.Vehicles, error) {

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

func GetVehicleInSpot(idVehicle int) models.VehiclesSpots {
	var vehicleSpot models.VehiclesSpots
	utils.Db.Table("vehicles_spots").Where("vehicle_id = ? AND exit_time IS NULL", idVehicle).First(&vehicleSpot)

	return vehicleSpot
}

func GetVehicles(idVehicle int) models.Vehicles {
	var vehicles models.Vehicles
	utils.Db.Table("vehicles").Where("id = ?", idVehicle).First(&vehicles)

	return vehicles
}

func GetVehiclePlateNumberBySpotId(spotId int) string {
	var vehicleSpot models.VehiclesSpots
	utils.Db.Table("vehicles_spots").Where("spot = ? AND exit_time IS NULL", spotId).Find(&vehicleSpot)

	var vehicle models.Vehicles
	utils.Db.Table("vehicles").Where("id = ?", vehicleSpot.VehicleId).Find(&vehicle)

	return vehicle.PlateNumber
}

func ExitVehicle(plateNumber string) (string, error) {
	if err := utils.ValidatePlateNumber(plateNumber); err != nil {
		return "", err
	}

	vehicle := models.Vehicles{}
	resultVehicles := utils.Db.Table("vehicles").Where("plate_number = ?", plateNumber).First(&vehicle)
	if errors.Is(resultVehicles.Error, gorm.ErrRecordNotFound) {
		return "", errors.New("El vehiculo no existe")
	}

	vehicleSpot := GetVehicleInSpot(vehicle.Id)
	if vehicleSpot.Id == 0 {
		return "", errors.New("El vehiculo no se encuentra estacionado")
	}

	utils.Db.Table("vehicles_spots").
		Where("vehicle_id = ? AND exit_time IS NULL", vehicleSpot.VehicleId).
		Update("exit_time", time.Now())

	utils.Db.Table("spots").
		Where("id = ?", vehicleSpot.Spot).
		Update("in_use", false)

	return "El vehiculo ha salido correctamente", nil
}
