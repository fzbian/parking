package controller

import (
	"errors"
	"fmt"
	"github.com/fzbian/parking/enums"
	"github.com/fzbian/parking/models"
	"github.com/fzbian/parking/utils"
	"gorm.io/gorm"
	"time"
)

func ParkingVehicle(request models.Vehicles) (string, error) {

	if !utils.ValidatePlateNumber(request.PlateNumber) {
		return "", errors.New("El formato de la placa no es valido")
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

	IsIn := IsInParking(request.PlateNumber)
	if IsIn {
		return "", errors.New("El vehiculo ya se encuentra estacionado")
	}

	spot := GetAvailableSpot(vehicle.VehicleType)

	if spot.ID == 0 {
		return "", errors.New("El parqueadero esta lleno")
	}

	fmt.Println("EntryTime: ", time.Now())
	fmt.Println("ExitTime: ", time.Time{})

	utils.Db.Table("vehicles_spots").Create(&models.VehiclesSpots{
		VehicleId: vehicle.Id,
		Spot:      spot.ID,
		EntryTime: time.Now(),
	})

	utils.Db.Table("spots").Where("id = ?", spot.ID).Update("in_use", true)

	SuccesfullMessage := fmt.Sprintf("El vehiculo %s ha sido estacionado en la bahia %d ubicada en la zona %s.", vehicle.PlateNumber, spot.ID, spot.Zone)
	return SuccesfullMessage, nil
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

func IsInParking(plate string) bool {
	var vehicles models.Vehicles
	resultVehicles := utils.Db.Table("vehicles").Where("plate_number = ?", plate).First(&vehicles)
	if errors.Is(resultVehicles.Error, gorm.ErrRecordNotFound) {
		return false
	}

	if vehicles.Id > 0 {
		var vehicleSpot models.VehiclesSpots
		resultVehiclesSpots := utils.Db.Table("vehicles_spots").Where("vehicle_id = ?", vehicles.Id).First(&vehicleSpot)
		if errors.Is(resultVehiclesSpots.Error, gorm.ErrRecordNotFound) {
			return false
		}
		return true
	}

	return false
}

func GetVehiclePlateNumberBySpotId(spotId int) (string, error) {
	var vehicleSpot models.VehiclesSpots
	resultVehiclesSpots := utils.Db.Table("vehicles_spots").Where("spot = ?", spotId).First(&vehicleSpot)
	if errors.Is(resultVehiclesSpots.Error, gorm.ErrRecordNotFound) {
		return "", errors.New("No se encontro el vehiculo")
	}

	var vehicle models.Vehicles
	resultVehicles := utils.Db.Table("vehicles").Where("id = ?", vehicleSpot.VehicleId).First(&vehicle)
	if errors.Is(resultVehicles.Error, gorm.ErrRecordNotFound) {
		return "", errors.New("No se encontro el vehiculo")
	}

	return vehicle.PlateNumber, nil
}

func ExitVehicle(plateNumber string) (string, error) {

	if !utils.ValidatePlateNumber(plateNumber) {
		return "", errors.New("El formato de la placa no es valido")
	}

	if !IsInParking(plateNumber) {
		return "", errors.New("El vehiculo no se encuentra estacionado")
	}

	var vehicle models.Vehicles
	utils.Db.Table("vehicles").Where("plate_number = ?", plateNumber).First(&vehicle)

	var vehicleSpot models.VehiclesSpots
	utils.Db.Table("vehicles_spots").Where("vehicle_id = ?", vehicle.Id).First(&vehicleSpot)
	utils.Db.Table("vehicles_spots").Where("vehicle_id = ?", vehicleSpot.VehicleId).Update("exit_time", time.Now())
	utils.Db.Table("spots").Where("id = ?", vehicleSpot.Spot).Update("in_use", false)

	var nextVehicleSpot models.VehiclesSpots
	if err := utils.Db.Table("vehicles_spots").Order("spot ASC").Where("spot > ?", vehicleSpot.Spot).First(&nextVehicleSpot).Error; err == nil {
		utils.Db.Table("vehicles_spots").Where("id = ?", nextVehicleSpot.Id).Update("spot", vehicleSpot.Spot)
	}

	return "El vehiculo ha salido correctamente", nil
}
