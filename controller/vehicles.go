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

// ParkingVehicle function receives a request of type "models.Vehicles" and returns a success message and an error if one occurs.
func ParkingVehicle(request models.Vehicles) (string, error) {

	// Validates that the vehicle's plate number is valid.
	if err := utils.ValidatePlateNumber(request.PlateNumber); err != nil {
		return "", err
	}

	// Gets the allowed vehicle types and validates that the requested vehicle type is within the allowed types.
	VehiclesTypes := enums.GetVehicleTypes()
	ValidVehicleType := utils.ValidateVehicleType(request.VehicleType, VehiclesTypes)
	if !ValidVehicleType {
		return "", errors.New("The vehicle type is not valid")
	}

	// Gets the existing vehicle or creates a new one if it does not exist.
	vehicle, err := GetOrCreateVehicle(request)
	if err != nil {
		panic(err)
	}

	// Verifies if the vehicle is already parked in a spot.
	vehicleSpot := GetVehicleInSpot(vehicle.Id)
	if vehicleSpot.Id != 0 {
		return "", errors.New("The vehicle is already parked")
	}

	// Gets the vehicle from the database.
	vehicles := GetVehicles(vehicle.Id)

	// Verifies that the vehicle type of the vehicle obtained from the database is the same as the vehicle type sent in the request.
	if request.VehicleType != vehicles.VehicleType {
		return "", errors.New("This vehicle already exists in the database and the vehicle type does not match.")
	}

	// Searches for an available spot for the requested vehicle type.
	spot := GetAvailableSpot(vehicle.VehicleType)

	// If there are no available spots, it returns an error.
	if spot.ID == 0 {
		return "", errors.New("The parking lot is full")
	}

	// Gets the current time
	now := time.Now()
	timePtr := &now

	// Creates a new row in the "vehicles_spots" table with the vehicle ID, the spot ID, and the entry time.
	utils.Db.Table("vehicles_spots").
		Create(&models.VehiclesSpots{
			VehicleId: vehicle.Id,
			Spot:      spot.ID,
			EntryTime: timePtr,
		})

	// Updates the spot's status to "in use".
	utils.Db.Table("spots").
		Where("id = ?", spot.ID).
		Update("in_use", true)

	// Creates a success message and returns it with a nil error.
	SuccessfullyMessage := fmt.Sprintf("The vehicle %s has been parked in spot %d located in zone %s.",
		vehicle.PlateNumber, spot.ID, spot.Zone)

	// Creates an entry ticket for the vehicle.
	CreateEntryTicket(vehicle.PlateNumber, vehicle.VehicleType, spot.Zone, spot.ID)

	return SuccessfullyMessage, nil

}

// The ExitVehicle function receives a plateNumber as input and returns a string and an error.
func ExitVehicle(plateNumber string) (string, error) {

	// Validates that the plate number is valid using the ValidatePlateNumber function from the utils package.
	if err := utils.ValidatePlateNumber(plateNumber); err != nil {
		return "", err
	}

	// Queries the database for a vehicle with the provided plate number.
	vehicle := models.Vehicles{}
	resultVehicles := utils.Db.Table("vehicles").
		Where("plate_number = ?", plateNumber).
		First(&vehicle)

	// If the result is a "record not found" error, returns an error indicating that the vehicle does not exist.
	if errors.Is(resultVehicles.Error, gorm.ErrRecordNotFound) {
		return "", errors.New("The vehicle does not exist")
	}

	// Gets the spot where the vehicle is parked.
	vehicleSpot := GetVehicleInSpot(vehicle.Id)

	// If the vehicle is not parked in any spot, returns an error.
	if vehicleSpot.Id == 0 {
		return "", errors.New("The vehicle is not parked")
	}

	// Gets the zone where the vehicle is parked.
	Zone := GetZoneFromVehicleSpot(vehicle.Id)

	// Updates the exit_time field of the vehicles_spots table with the current time for the vehicle that is leaving.
	utils.Db.Table("vehicles_spots").
		Where("vehicle_id = ? AND exit_time IS NULL", vehicleSpot.VehicleId).
		Update("exit_time", time.Now())

	// Updates the in_use field of the spots table to false for the spot where the vehicle was parked.
	utils.Db.Table("spots").
		Where("id = ?", vehicleSpot.Spot).
		Update("in_use", false)

	// Gets the entry time of the vehicle in the format "HH:MM:SS".
	EntryTime := vehicleSpot.EntryTime.Format("15:04:05")

	// Creates an exit ticket for the vehicle.
	CreateExitTicket(vehicle.PlateNumber, vehicle.VehicleType, Zone, EntryTime, vehicleSpot.Spot)

	// Returns a success message.
	return "The vehicle has been successfully delivered", nil

}

// ExitAllVehicles exits all parked vehicles in the parking lot
func ExitAllVehicles() (string, error) {

	// Get all parked vehicles
	vehiclesSpots := GetVehiclesSpots()

	// Loop through all parked vehicles
	for _, vehicleSpot := range vehiclesSpots {

		// Get the corresponding vehicle
		vehicle := GetVehicles(vehicleSpot.VehicleId)

		// Get the zone of the parked vehicle
		Zone := GetZoneFromVehicleSpot(vehicle.Id)

		// Update the exit time of the parked vehicle
		utils.Db.Table("vehicles_spots").
			Where("vehicle_id = ? AND exit_time IS NULL", vehicleSpot.VehicleId).
			Update("exit_time", time.Now())

		// Mark the spot as available
		utils.Db.Table("spots").
			Where("id = ?", vehicleSpot.Spot).
			Update("in_use", false)

		// Create an exit ticket for the parked vehicle
		EntryTime := vehicleSpot.EntryTime.Format("15:04:05")
		CreateExitTicket(vehicle.PlateNumber, vehicle.VehicleType, Zone, EntryTime, vehicleSpot.Spot)
	}

	// Return a success message
	return "All vehicles have been successfully delivered", nil

}
