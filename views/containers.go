package views

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
)

var MainContainer *fyne.Container

func GetMainContainer(window fyne.Window, logo fyne.Resource) *fyne.Container {
	VehicleListTable := GetTable()
	LogoObject := canvas.NewImageFromResource(logo)
	AddVehicleButton, ExitVehicleButton := LeftButtons(window, VehicleListTable)
	ReportsByVehicleType, ReportsByZoneType, ExitButton := RightButtons(window)

	LogoContainer := container.New(layout.NewGridWrapLayout(fyne.Size{
		Width:  175,
		Height: 175,
	}), LogoObject)

	LeftContainer := container.New(layout.NewVBoxLayout(),
		AddVehicleButton,
		ExitVehicleButton,
		LogoContainer,
	)

	MidContainer := container.New(layout.NewGridWrapLayout(
		fyne.Size{
			Width:  900,
			Height: 650,
		}), VehicleListTable)

	RightContainer := container.New(layout.NewVBoxLayout(),
		ReportsByVehicleType,
		ReportsByZoneType,
		ExitButton,
	)

	MainContainer = container.New(layout.NewHBoxLayout(),
		LeftContainer,
		MidContainer,
		RightContainer,
	)

	return MainContainer
}
