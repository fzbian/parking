package views

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"github.com/fzbian/parking/controller"
	"github.com/fzbian/parking/models"
)

func LeftButtons(window fyne.Window) (*widget.Button, *widget.Button) {

	AddVehicleForm, AddVehiclePlateEntry, VehicleTypeEntry := AddVehicleDialog()
	ExitVehicleForm, ExitVehiclePlateEntry := ExitVehicleDialog()

	AddVehicleButton := widget.NewButton("Entrada de vehiculo", func() {

		AddVehiclePlateEntry.SetText("")
		VehicleTypeEntry.SetSelected("")

		dialog.ShowCustomConfirm("Entrada de vehiculo", "Agregar", "Cancelar", AddVehicleForm, func(b bool) {
			if b {
				message, err := controller.ParkingVehicle(models.Vehicles{
					PlateNumber: AddVehiclePlateEntry.Text,
					VehicleType: VehicleTypeEntry.Selected,
				})
				NewPopUp(message, window)
				// Refresh table
				if err != nil {
					NewPopUp(err.Error(), window)
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
					NewPopUp(err.Error(), window)
				}
				NewPopUp(message, window)
				// Refresh table
			}
			ExitVehicleForm.Hide()
		}, window)

		window.CenterOnScreen()
		ExitVehicleForm.Show()
	})

	return AddVehicleButton, ExitVehicleButton
}

func RightButtons(window fyne.Window) (*widget.Button, *widget.Button) {
	ExportRecordsButton := widget.NewButton("Exportar registros", func() {
		NewPopUp("Funcion pendiente", window)
	})

	ExitButton := widget.NewButton("Salir", func() {
		window.Close()
	})

	return ExportRecordsButton, ExitButton
}
