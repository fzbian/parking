package main

import (
	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/layout"
	"github.com/fzbian/parking/views"
)

func main() {

	app := app.New()

	window := app.NewWindow("Parqueadero Colegio Jose Max Leon")

	window.Resize(fyne.Size{
		Width:  800,
		Height: 600,
	})

	VehicleTable := views.GetVehiclesTable()

	AddVehicleButton, ExitVehicleButton := views.LeftButtons(window)
	ExportRecordsButton, ExitWindowButton := views.RightButtons(window)

	LeftContainer := fyne.NewContainerWithLayout(layout.NewVBoxLayout(), AddVehicleButton, ExitVehicleButton)
	MidContainer := fyne.NewContainerWithLayout(layout.NewGridWrapLayout(fyne.NewSize(900, 650)), VehicleTable)
	RightContainer := fyne.NewContainerWithLayout(layout.NewVBoxLayout(), ExportRecordsButton, ExitWindowButton)

	HorizontalContainer := fyne.NewContainerWithLayout(layout.NewHBoxLayout(),
		LeftContainer,
		MidContainer,
		RightContainer,
	)

	window.SetContent(HorizontalContainer)
	window.ShowAndRun()
}
