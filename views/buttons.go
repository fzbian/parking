package views

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/fzbian/parking/controller"
	"github.com/fzbian/parking/models"
)

// LeftButtons creates the buttons that will be on the left side and returns the pointers
// to widget.Button to be used in the containers.
func LeftButtons(window fyne.Window, table *widget.Table) (*widget.Button, *widget.Button, *widget.Button) {

	// Gets the dialogs of the forms with their inputs
	AddVehicleForm, AddVehiclePlateEntry, VehicleTypeEntry := AddVehicleDialog()
	ExitVehicleForm, ExitVehiclePlateEntry := ExitVehicleDialog()

	// Create AddVehicleButton and set the function to be executed when the button is clicked.
	AddVehicleButton := widget.NewButton("Vehicle Entry", func() {

		// Set the text of the inputs to empty.
		AddVehiclePlateEntry.SetText("")
		VehicleTypeEntry.SetSelected("")

		// Show the dialog with the form and the buttons.
		dialog.ShowCustomConfirm("Vehicle Entry", "Add", "Cancel", AddVehicleForm, func(b bool) {
			// If the user clicks on the Add button, the function will be executed.
			if b {
				// Call the controller function to add the vehicle to the database.
				message, err := controller.ParkingVehicle(models.Vehicles{
					PlateNumber: AddVehiclePlateEntry.Text,
					VehicleType: VehicleTypeEntry.Selected,
				})
				// If there is an error, show the error message in a dialog.
				if err != nil {
					dialog.ShowInformation("Error", err.Error(), window)
					return
				}
				// If there is no error, show the message in a dialog.
				dialog.ShowInformation("Information", message, window)
				// Update the table with the new data.
				table = GetTable()
				MidContainer := container.New(layout.NewGridWrapLayout(
					fyne.Size{
						Width:  900,
						Height: 650,
					}), table)
				MainContainer.Objects[1] = MidContainer
			}
			// Hide the dialog.
			AddVehicleForm.Hide()
		}, window)

		// Center the dialog on the screen and show it.
		window.CenterOnScreen()
		AddVehicleForm.Show()
	})

	// Create ExitVehicleButton and set the function to be executed when the button is clicked.
	ExitVehicleButton := widget.NewButton("Vehicle Exit", func() {

		// Set the text of the input to empty.
		ExitVehiclePlateEntry.SetText("")

		// Show the dialog with the form and the buttons.
		dialog.ShowCustomConfirm("Vehicle Exit", "Exit", "Cancel", ExitVehicleForm, func(b bool) {
			// If the user clicks on the Exit button, the function will be executed.
			if b {
				// Call the controller function to exit the vehicle from the database.
				message, err := controller.ExitVehicle(ExitVehiclePlateEntry.Text)
				// If there is an error, show the error message in a dialog.
				if err != nil {
					dialog.ShowInformation("Error", err.Error(), window)
					return
				}
				// Update the table with the new data.
				table = GetTable()
				MidContainer := container.New(layout.NewGridWrapLayout(
					fyne.Size{
						Width:  900,
						Height: 650,
					}), table)
				MainContainer.Objects[1] = MidContainer
				// If there is no error, show the message in a dialog.
				dialog.ShowInformation("Information", message, window)
			}
			// Hide the dialog.
			ExitVehicleForm.Hide()
		}, window)

		// Center the dialog on the screen and show it.
		window.CenterOnScreen()
		ExitVehicleForm.Show()
	})

	// Create ExitAllVehicles and set the function to be executed when the button is clicked.
	ExitAllVehicles := widget.NewButton("Exit all vehicles", func() {

		// Show the dialog with the form and the buttons.
		dialog.ShowCustomConfirm("Exit all vehicles", "Sure", "Cancel", widget.NewLabel("Are you sure you want to remove all vehicles from the parking lot?"), func(b bool) {
			// If the user clicks on the Sure button, the function will be executed.
			if b {
				// Call the controller function to exit all vehicles from the database.
				message, err := controller.ExitAllVehicles()
				// If there is an error, show the error message in a dialog.
				if err != nil {
					dialog.ShowInformation("Error", err.Error(), window)
					return
				}
				// Update the table with the new data.
				table = GetTable()
				MidContainer := container.New(layout.NewGridWrapLayout(
					fyne.Size{
						Width:  900,
						Height: 650,
					}), table)
				MainContainer.Objects[1] = MidContainer
				// If there is no error, show the message in a dialog.
				dialog.ShowInformation("Information", message, window)
			}
		}, window)
	})

	// Return the pointers to the buttons.
	return AddVehicleButton, ExitVehicleButton, ExitAllVehicles
}

// RightButtons creates the buttons that will be on the right side and returns the pointers
// to widget.Button to be used in the containers.
func RightButtons(window fyne.Window) (*widget.Button, *widget.Button, *widget.Button, *widget.Button, *widget.Button) {

	// Gets the dialogs of the forms with their inputs
	RecordsVehiclesTypeForm, VehicleTypeEntry := RecordsVehiclesTypeDialog()
	RecordsZoneForm, ZoneEntry := RecordsZoneDialog()

	// Create ReportsByVehicleType and set the function to be executed when the button is clicked.
	ReportsByVehicleType := widget.NewButton("Report by vehicle type", func() {

		// Set the text of the input to empty.
		VehicleTypeEntry.SetSelected("")

		// Show the dialog with the form and the buttons.
		dialog.ShowCustomConfirm("Report by vehicle type", "Generate", "Cancel", RecordsVehiclesTypeForm, func(b bool) {
			// If the user clicks on the Generate button, the function will be executed.
			if b {
				// Call the controller function to get the vehicles by type.
				result, err := controller.GetVehiclesBySpotType(VehicleTypeEntry.Selected)
				// If there is an error, show the error message in a dialog.
				if err != nil {
					dialog.ShowInformation("Error", err.Error(), window)
					return
				}
				// If there is no error, show the message in a dialog.
				dialog.ShowInformation("Information", fmt.Sprintf("In the last 24 hours, %d vehicles of type %s, entered the parking lot.", len(result), VehicleTypeEntry.Selected), window)
			}
			// Hide the dialog.
			RecordsVehiclesTypeForm.Hide()
		}, window)

		// Center the dialog on the screen and show it.
		window.CenterOnScreen()
		RecordsVehiclesTypeForm.Show()
	})

	// Create ReportsByZone and set the function to be executed when the button is clicked.
	ReportsByZone := widget.NewButton("Report by zone", func() {

		// Set the text of the input to empty.
		ZoneEntry.SetSelected("")

		// Show the dialog with the form and the buttons.
		dialog.ShowCustomConfirm("Report by zone", "Generate", "Cancel", RecordsZoneForm, func(b bool) {
			// If the user clicks on the Generate button, the function will be executed.
			if b {
				// Call the controller function to get the vehicles by zone.
				result, err := controller.GetVehiclesByZoneType(ZoneEntry.Selected)
				// If there is an error, show the error message in a dialog.
				if err != nil {
					dialog.ShowInformation("Error", err.Error(), window)
					return
				}
				// If there is no error, show the message in a dialog.
				dialog.ShowInformation("Information", fmt.Sprintf("In the last 24 hours, %d vehicles in zone %s entered the parking lot.", len(result), ZoneEntry.Selected), window)
			}
			// Hide the dialog.
			RecordsZoneForm.Hide()
		}, window)

		// Center the dialog on the screen and show it.
		window.CenterOnScreen()
		RecordsZoneForm.Show()
	})

	// Create ReportsProvidersCollection and set the function to be executed when the button is clicked.
	ReportsProvidersCollection := widget.NewButton("Money collection", func() {

		// Call the controller function to get the vehicles by type.
		toPay, err := controller.GetVehiclesProvidersToPay()
		// If there is an error, show the error message in a dialog.
		if err != nil {
			dialog.ShowInformation("Error", err.Error(), window)
			return
		}
		// If there is no to pay providers, show the message in a dialog.
		if len(toPay) == 0 {
			dialog.ShowInformation("Error", "No providers have entered in the last 24 hours.", window)
			return
		}
		dialog.ShowInformation("Information", fmt.Sprintf("In total, $%d pesos have been collected.", len(toPay)*5000), window)
	})

	// ZoneMostUsed creates the button that will be on the right side and returns the pointer
	ZoneMostUsed := widget.NewButton("Zone most used", func() {

		// Call the controller function to get the most used zone.
		zoneMostUsed, err := controller.GetMostUsedZone()
		// If there is an error, show the error message in a dialog.
		if err != nil {
			dialog.ShowInformation("Error", err.Error(), window)
			return
		}
		// If there is no error, show the message in a dialog.
		dialog.ShowInformation("Information", fmt.Sprintf("The most used zone has been the %s zone.", zoneMostUsed), window)
	})

	// Create ExitButton and set the function to be executed when the button is clicked.
	ExitButton := widget.NewButton("Exit", func() {
		// Close the window.
		window.Close()
	})

	// Return the pointers to the buttons.
	return ReportsByVehicleType, ReportsByZone, ReportsProvidersCollection, ZoneMostUsed, ExitButton
}
