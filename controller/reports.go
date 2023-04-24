package controller

import (
	"errors"
	"github.com/fzbian/parking/models"
	"github.com/fzbian/parking/utils"
	"time"
)

// GetTotalTimeByZone function receives a request of type "models.Spot" and returns the total time
// that the vehicles have been parked in the zone and an error if one occurs.
// TODO: bug when reusing the function updating the data
func GetTotalTimeByZone(request models.Spot) (string, error) {

	// Declare an empty slice of models.VehiclesSpots structs
	var vehiclesSpots []models.VehiclesSpots

	// Build a subquery to select all spot ids that match the given zone and store the query in subquery variable
	subquery := utils.Db.Table("spots").
		Where("zone = ?", request.Zone).
		Select("id")

	// Build a query to select all vehicles spots where the spot is in the subquery's
	// list of spot ids and store the results in vehiclesSpots slice
	result := utils.Db.Table("vehicles_spots").
		Where("spot IN (?)", subquery).
		Find(&vehiclesSpots)

	// If there was an error with the query, return an empty string and the error
	if result.Error != nil {
		return "", result.Error
	}

	// Declare a variable to store the total time and initialize it to 0
	var totalTime time.Duration
	// Loop through each vehicleSpot in the vehiclesSpots slice and add the difference
	// between its exit time and entry time to totalTime
	for _, vehicleSpot := range vehiclesSpots {
		totalTime += vehicleSpot.ExitTime.Sub(vehicleSpot.EntryTime)
	}

	// Convert totalTime to a string and return it along with nil error
	return totalTime.String(), nil

}

// GetMostUsedZone function returns the most used zone and an error if one occurs.
func GetMostUsedZone() (string, error) {

	// Declare two variables: mostUsedZone as string and mostUsedZoneTime as time.Duration,
	// and initialize them to empty string and 0 respectively
	var mostUsedZone string
	var mostUsedZoneTime time.Duration

	// Start with zone A and loop through zones A, B, and C
	zone := "A"
	for i := 0; i < 3; i++ {
		// Call the GetTotalTimeByZone function with a Spot struct that has the current zone,
		// and store the result in zoneTime variable
		zoneTime, err := GetTotalTimeByZone(models.Spot{
			Zone: zone})
		if err != nil {
			return "", err
		}

		// Parse the zoneTime into a time.Duration variable called zoneTimeDuration
		zoneTimeDuration, err := time.ParseDuration(zoneTime)
		if err != nil {
			return "", err
		}

		// If the zoneTimeDuration is greater than the current mostUsedZoneTime,
		// update the mostUsedZone and mostUsedZoneTime variables
		if zoneTimeDuration > mostUsedZoneTime {
			mostUsedZone = zone
			mostUsedZoneTime = zoneTimeDuration
		}

		// Move to the next zone
		zone = string(int(zone[0]) + 1)
	}

	// If there was no zone with any records, return an error
	if mostUsedZone == "" {
		return "", errors.New("There are no records yet to make this report.")
	}

	// Return the most used zone and nil error
	return mostUsedZone, nil

}

// GetVehiclesProvidersToPay returns the supplier vehicles in a filter of those that took more
// than 30 minutes to be returned
func GetVehiclesProvidersToPay() ([]models.Vehicles, error) {

	// Declare a variable called vehicles as a slice of models.Vehicles
	var vehicles []models.Vehicles

	// Define a subquery that selects the IDs of all vehicles with a vehicle_type of 'Provider'
	subquery := utils.Db.Table("vehicles").
		Where("vehicle_type = 'Provider'").
		Select("id")

	// Define a second subquery that selects the vehicle IDs of all vehicles that have parked for
	// more than 30 minutes the subquery selects the vehicle IDs from the vehicles_spots table where
	// the vehicle ID is in the IDs from the first subquery, and the entry time is more than 30 minutes ago
	subquery2 := utils.Db.Table("vehicles_spots").
		Where("vehicle_id IN (?) AND entry_time < ?", subquery, time.Now().Add(-30*time.Minute)).
		Select("vehicle_id")

	// Run a query that selects all vehicles from the vehicles table where the ID is in the IDs
	// from the second subquery
	result := utils.Db.Table("vehicles").
		Where("id IN (?)", subquery2).
		Find(&vehicles)

	// If there was an error during the query, return nil slice and the error
	if result.Error != nil {
		return nil, result.Error
	}

	// Return the vehicles slice and nil error
	return vehicles, nil

}

// GetVehiclesByZoneType returns the vehicles that have been parked in a zone for more than 24 hours
func GetVehiclesByZoneType(zoneType string) ([]models.Vehicles, error) {

	// Define an empty slice of models.Vehicles struct
	var vehicles []models.Vehicles

	// Create a subquery to retrieve all the spots that match the given zone type
	subquery := utils.Db.Table("spots").
		Where("zone = ?", zoneType).
		Select("id")

	// Create another subquery to retrieve all the vehicles that parked in the spots matching the subquery above
	// within the last 24 hours
	subquery2 := utils.Db.Table("vehicles_spots").
		Where("spot IN (?) AND entry_time > ?", subquery, time.Now().Add(-24*time.Hour)).
		Select("vehicle_id")

	// Retrieve all the vehicles whose id matches the subquery2
	result := utils.Db.Table("vehicles").
		Where("id IN (?)", subquery2).
		Find(&vehicles)

	// Check for any error while querying the database and return it along with the retrieved vehicles
	if result.Error != nil {
		return nil, result.Error
	}

	// Return the vehicles slice and nil error
	return vehicles, nil

}
