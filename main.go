package main

import (
	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/dialog"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
	"github.com/fzbian/parking/controller"
	"github.com/fzbian/parking/models"
	"github.com/fzbian/parking/views"
)

func main() {

	app := app.New()

	window := app.NewWindow("Parqueadero Colegio Jose Max Leon")

	window.Resize(fyne.Size{
		Width:  800,
		Height: 600,
	})

	data := views.GetData()
	table := views.GetTable(data)

	AddVehicleForm, AddVehiclePlateEntry, VehicleTypeEntry := views.AddVehicleDialog()
	ExitVehicleForm, ExitVehiclePlateEntry := views.ExitVehicleDialog()

	AddVehicleButton := widget.NewButton("Entrada de vehiculo", func() {

		AddVehiclePlateEntry.SetText("")
		VehicleTypeEntry.SetSelected("")

		dialog.ShowCustomConfirm("Entrada de vehiculo", "Agregar", "Cancelar", AddVehicleForm, func(b bool) {
			if b {
				message, err := controller.ParkingVehicle(models.Vehicles{
					PlateNumber: AddVehiclePlateEntry.Text,
					VehicleType: VehicleTypeEntry.Selected,
				})
				views.NewPopUp(message, window)
				if err != nil {
					views.NewPopUp(err.Error(), window)
				} else {
				}
			}
			AddVehicleForm.Hide()
		}, window)

		window.CenterOnScreen()
		AddVehicleForm.Show()
	})

	ExitVehicleButton := widget.NewButton("Salida de vehiculo", func() {

		ExitVehiclePlateEntry.SetText("")

		dialog.ShowCustomConfirm("Salida de vehiculo", "Agregar", "Cancelar", ExitVehicleForm, func(b bool) {
			if b {
				message, err := controller.ExitVehicle(ExitVehiclePlateEntry.Text)
				if err != nil {
					views.NewPopUp(err.Error(), window)
				}
				views.NewPopUp(message, window)
			}
			ExitVehicleForm.Hide()
		}, window)

		window.CenterOnScreen()
		ExitVehicleForm.Show()
	})

	ButtonContainer := fyne.NewContainerWithLayout(layout.NewVBoxLayout(), AddVehicleButton, ExitVehicleButton)
	RightContainer := fyne.NewContainerWithLayout(layout.NewGridWrapLayout(fyne.NewSize(900, 650)), table)
	MidContainer := fyne.NewContainerWithLayout(layout.NewHBoxLayout(), RightContainer)
	HorizontalContainer := fyne.NewContainerWithLayout(layout.NewHBoxLayout(), ButtonContainer, MidContainer)

	window.SetContent(HorizontalContainer)
	window.ShowAndRun()
}
