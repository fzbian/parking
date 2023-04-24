package controller

import (
	"github.com/fzbian/parking/models"
	"github.com/fzbian/parking/utils"
	"time"
)

// GetVehicleInSpot gets the vehicle spot for the given vehicle ID.
func GetVehicleInSpot(idVehicle int) models.VehiclesSpots {

	// Create a new variable to store the vehicle spot.
	var vehicleSpot models.VehiclesSpots

	// Use the Db object to query the vehicles_spots table.
	// The Where method filters the results to only include rows where
	// the vehicle_id column matches the given ID and the exit_time column is NULL.
	// The First method returns the first row in the results.
	utils.Db.Table("vehicles_spots").
		Where("vehicle_id = ? AND exit_time IS NULL", idVehicle).
		First(&vehicleSpot)

	// Return the vehicle spot.
	return vehicleSpot

}

// GetVehicles gets the vehicle for the given vehicle ID.
func GetVehicles(idVehicle int) models.Vehicles {

	// Create a new variable to store the vehicle.
	var vehicles models.Vehicles

	// Use the Db object to query the vehicles table.
	// The Where method filters the results to only include rows where the id column matches the given ID.
	// The First method returns the first row in the results.
	utils.Db.Table("vehicles").
		Where("id = ?", idVehicle).
		First(&vehicles)

	// Return the vehicle.
	return vehicles

}

// GetVehiclesSpots gets all the vehicle spots.
func GetVehiclesSpots() []models.VehiclesSpots {

	// Create a new slice to store the vehicle spots.
	var vehiclesSpots []models.VehiclesSpots

	// Use the Db object to query the `vehicles_spots` table.
	// The `Find` method returns all the rows in the table.
	utils.Db.Table("vehicles_spots").
		Find(&vehiclesSpots)

	// Return the vehicle spots.
	return vehiclesSpots

}

// GetAvailableSpot gets the first available spot of the given type.
func GetAvailableSpot(spotType string) models.Spot {

	// Create a new spot variable.
	spot := models.Spot{}

	// Use the Db object to query the spots table.
	// The Where method filters the results to only include rows where the
	// type column matches the given type and the in_use column is false.
	// The Find method returns the first row in the results.
	// The Order method sorts the results by the id column.
	utils.Db.Table("spots").
		Where("type = ? AND in_use IS false", spotType).
		Find(&spot).
		Order("id")

	// If the spot ID is 0, then there are no available spots of the given type.
	if spot.ID == 0 {

		// Try to find an available normal spot.
		utils.Db.Table("spots").
			Where("type = 'NORMAL' AND in_use IS false").
			Find(&spot).
			Order("id")
	}

	// Return the spot.
	return spot

}

// GetZoneFromVehicleSpot gets the zone of the vehicle spot for the given vehicle ID.
func GetZoneFromVehicleSpot(idVehicle int) string {

	// Create a new variable to store the vehicle spot.
	var vehicleSpot models.VehiclesSpots

	// Use the Db object to query the vehicles_spots table.
	// The Where method filters the results to only include rows where the
	//vehicle_id column matches the given ID and the exit_time column is NULL.
	// The First method returns the first row in the results.
	utils.Db.Table("vehicles_spots").
		Where("vehicle_id = ? AND exit_time IS NULL", idVehicle).
		First(&vehicleSpot)

	// Create a new variable to store the spot.
	var spot models.Spot

	// Use the Db object to query the spots table.
	// The Where method filters the results to only include rows where the id
	// column matches the spot column of the vehicles_spots table.
	// The First method returns the first row in the results.
	utils.Db.Table("spots").
		Where("id = ?", vehicleSpot.Spot).
		First(&spot)

	// Return the zone of the spot.
	return spot.Zone

}

// GetVehiclePlateNumberBySpotId returns the plate number of the vehicle that
// is currently parked in the given spot.
func GetVehiclePlateNumberBySpotId(spotId int) string {

	// Create a variable to store the vehicle spot information.
	var vehicleSpot models.VehiclesSpots
	// Query the database for the vehicle spot with the given ID.
	// The WHERE clause ensures that we only return vehicle spots that are still in
	// the parking lot (i.e. exit_time is NULL).
	utils.Db.Table("vehicles_spots").
		Where("spot = ? AND exit_time IS NULL", spotId).
		Find(&vehicleSpot)

	// Create a variable to store the vehicle information.
	var vehicle models.Vehicles

	// Query the database for the vehicle with the ID from the vehicle spot information.
	utils.Db.Table("vehicles").
		Where("id = ?", vehicleSpot.VehicleId).
		Find(&vehicle)

	// Return the plate number of the vehicle.
	return vehicle.PlateNumber

}

// GetVehiclesBySpotType returns all the vehicles that have been parked in the past 24 hours of the given type.
func GetVehiclesBySpotType(vehicleType string) ([]models.VehiclesSpots, error) {
	// Create a slice to store the vehicle spots.
	var vehiclesSpots []models.VehiclesSpots

	// Create a subquery to get the IDs of all vehicles of the given type.
	subquery := utils.Db.Table("vehicles").
		Where("vehicle_type = ?", vehicleType).
		Select("id")

	// Query the database for all vehicle spots that are of the given
	// type and that have been parked in the past 24 hours.
	result := utils.Db.Table("vehicles_spots").
		Where("vehicle_id IN (?) AND entry_time > ?", subquery, time.Now().Add(-24*time.Hour)).
		Find(&vehiclesSpots)

	// If there was an error, return an error.
	if result.Error != nil {
		return nil, result.Error
	}

	// Return the list of vehicle spots.
	return vehiclesSpots, nil

}

// GetOrCreateVehicle gets or creates a vehicle in the database.
// The req parameter is a models.Vehicles object that contains the information about the vehicle to get or create.
// The function returns a models.Vehicles object and an error.
func GetOrCreateVehicle(req models.Vehicles) (models.Vehicles, error) {

	// Create a new `models.Vehicles` object.
	vehicles := models.Vehicles{}

	// Query the database for a vehicle with the given plate number.
	// If a vehicle is found, the `FirstOrCreate()` method will update the existing
	// vehicle with the information from the `req` object.
	// If a vehicle is not found, the `FirstOrCreate()` method will create a new vehicle
	// with the information from the `req` object.
	err := utils.Db.Table("vehicles").
		Where("plate_number = ?", req.PlateNumber).
		FirstOrCreate(&req).
		Scan(&vehicles).Error
	if err != nil {
		// If there was an error, return the error.
		return vehicles, err
	}

	// Return the vehicle.
	return vehicles, nil

}
