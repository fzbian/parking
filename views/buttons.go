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

func LeftButtons(window fyne.Window, table *widget.Table) (*widget.Button, *widget.Button, *widget.Button) {

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
				if err != nil {
					dialog.ShowInformation("Error", err.Error(), window)
					return
				}
				dialog.ShowInformation("Informacion", message, window)
				table = GetTable()
				MidContainer := container.New(layout.NewGridWrapLayout(
					fyne.Size{
						Width:  900,
						Height: 650,
					}), table)
				MainContainer.Objects[1] = MidContainer
			}
			AddVehicleForm.Hide()
		}, window)

		window.CenterOnScreen()
		AddVehicleForm.Show()
	})

	ExitVehicleButton := widget.NewButton("Salida de vehiculo", func() {

		ExitVehiclePlateEntry.SetText("")

		dialog.ShowCustomConfirm("Salida de vehiculo", "Salida", "Cancelar", ExitVehicleForm, func(b bool) {
			if b {
				message, err := controller.ExitVehicle(ExitVehiclePlateEntry.Text)
				if err != nil {
					dialog.ShowInformation("Error", err.Error(), window)
					return
				}
				table = GetTable()
				MidContainer := container.New(layout.NewGridWrapLayout(
					fyne.Size{
						Width:  900,
						Height: 650,
					}), table)
				MainContainer.Objects[1] = MidContainer
				dialog.ShowInformation("Informacion", message, window)
			}
			ExitVehicleForm.Hide()
		}, window)

		window.CenterOnScreen()
		ExitVehicleForm.Show()
	})

	ExitAllVehicles := widget.NewButton("Salida a todos los vehiculos", func() {

		dialog.ShowCustomConfirm("Salida a todos los vehiculos", "Salida", "Cancelar", widget.NewLabel("¿Esta seguro que desea sacar a todos los vehiculos del parqueadero?"), func(b bool) {
			if b {
				message, err := controller.ExitAllVehicles()
				if err != nil {
					dialog.ShowInformation("Error", err.Error(), window)
					return
				}
				table = GetTable()
				MidContainer := container.New(layout.NewGridWrapLayout(
					fyne.Size{
						Width:  900,
						Height: 650,
					}), table)
				MainContainer.Objects[1] = MidContainer
				dialog.ShowInformation("Informacion", message, window)
			}
		}, window)
	})

	return AddVehicleButton, ExitVehicleButton, ExitAllVehicles
}

func RightButtons(window fyne.Window) (*widget.Button, *widget.Button, *widget.Button, *widget.Button) {
	RecordsVehiclesTypeForm, VehicleTypeEntry := RecordsVehiclesTypeDialog()
	RecordsZoneForm, ZoneEntry := RecordsZoneDialog()

	ReportsByVehicleType := widget.NewButton("Reporte por tipo de vehiculo", func() {

		VehicleTypeEntry.SetSelected("")

		dialog.ShowCustomConfirm("Reporte por tipo de vehiculo", "Generar", "Cancelar", RecordsVehiclesTypeForm, func(b bool) {
			if b {
				result, err := controller.GetVehiclesBySpotType(VehicleTypeEntry.Selected)
				if err != nil {
					dialog.ShowInformation("Error", err.Error(), window)
					return
				}
				message := fmt.Sprintf("En las ultimas 24 horas, %d vehiculos de tipo %s, ingresaron a el parqueadero.", len(result), VehicleTypeEntry.Selected)
				dialog.ShowInformation("Informacion", message, window)
			}
			RecordsVehiclesTypeForm.Hide()
		}, window)

		window.CenterOnScreen()
		RecordsVehiclesTypeForm.Show()
	})

	ReportsByZone := widget.NewButton("Reporte por zona", func() {

		ZoneEntry.SetSelected("")

		dialog.ShowCustomConfirm("Reporte por zona", "Generar", "Cancelar", RecordsZoneForm, func(b bool) {
			if b {
				result, err := controller.GetVehiclesByZoneType(ZoneEntry.Selected)
				if err != nil {
					dialog.ShowInformation("Error", err.Error(), window)
					return
				}
				message := fmt.Sprintf("En las ultimas 24 horas, %d vehiculos en la zona %s ingresaron a el parqueadero.", len(result), ZoneEntry.Selected)
				dialog.ShowInformation("Informacion", message, window)
			}
			RecordsZoneForm.Hide()
		}, window)

		window.CenterOnScreen()
		RecordsZoneForm.Show()
	})

	ReportsProvidersCollection := widget.NewButton("Recoleccion de dinero", func() {

		toPay, err := controller.GetVehiclesProvidersToPay()
		if err != nil {
			dialog.ShowInformation("Error", err.Error(), window)
			return
		}
		dialog.ShowInformation("Informacion", fmt.Sprintf("En total se han recolectado: $%d pesos.", len(toPay)*5000), window)
	})

	ExitButton := widget.NewButton("Salir", func() {
		window.Close()
	})

	return ReportsByVehicleType, ReportsByZone, ReportsProvidersCollection, ExitButton
}
