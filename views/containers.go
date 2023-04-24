package views

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
)

var MainContainer *fyne.Container

// GetMainContainer returns the main container of the application.
func GetMainContainer(window fyne.Window, logo fyne.Resource) *fyne.Container {

	// Obtains the buttons, tables and images to be placed in each of the containers.
	VehicleListTable := GetTable()
	LogoObject := canvas.NewImageFromResource(logo)
	AddVehicleButton, ExitVehicleButton, ExitAllVehicles := LeftButtons(window, VehicleListTable)
	ReportsByVehicleType, ReportsByZoneType, ReportsProvidersCollection, ZoneMostUsed, ExitButton := RightButtons(window)

	// Create a container for the logo of the school.
	LogoContainer := container.New(layout.NewGridWrapLayout(fyne.Size{
		Width:  175,
		Height: 175,
	}), LogoObject)

	// Create the containers for the buttons and logo of the school.
	LeftContainer := container.New(layout.NewVBoxLayout(),
		AddVehicleButton,
		ExitVehicleButton,
		LogoContainer,
	)

	// Create the container for the table.
	MidContainer := container.New(layout.NewGridWrapLayout(
		fyne.Size{
			Width:  900,
			Height: 650,
		}), VehicleListTable)

	// Create the container for the buttons.
	RightContainer := container.New(layout.NewVBoxLayout(),
		ReportsByVehicleType,
		ReportsByZoneType,
		ReportsProvidersCollection,
		ZoneMostUsed,
		ExitAllVehicles,
		ExitButton,
	)

	// Create the main container.
	MainContainer = container.New(layout.NewHBoxLayout(),
		LeftContainer,
		MidContainer,
		RightContainer,
	)

	// Return the main container.
	return MainContainer

}
