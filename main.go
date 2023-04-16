package main

import (
	"fmt"
	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/dialog"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
	"github.com/fzbian/parking/models"
	"github.com/fzbian/parking/utils"
)

func main() {
	app := app.New()

	window := app.NewWindow("Parqueadero Colegio Jose Max Leon")

	window.Resize(fyne.Size{
		Width:  800,
		Height: 600,
	})

	AddVehicleForm, AddVehiclePlateEntry, VehicleTypeEntry := AddVehicleDialog()

	AddVehicleButton := widget.NewButton("Agregar vehiculo", func() {
		dialog.ShowCustomConfirm("Agregar vehiculo", "Agregar", "Cancelar", AddVehicleForm, func(b bool) {
			if b {
				fmt.Println("Placa: ", AddVehiclePlateEntry.Text)
				fmt.Println("Tipo de vehiculo: ", VehicleTypeEntry.Selected)
			}
			AddVehicleForm.Hide()
		}, window)
	})

	ExitVehicleForm, ExitVehiclePlateEntry := ExitVehicleDialog()

	ExitVehicleButton := widget.NewButton("Sacar vehiculo", func() {
		dialog.ShowCustomConfirm("Sacar vehiculo", "Agregar", "Cancelar", ExitVehicleForm, func(b bool) {
			if b {
				fmt.Println("Placa: ", ExitVehiclePlateEntry.Text)
				fmt.Println("Tipo de vehiculo: ", VehicleTypeEntry.Selected)
			}
			ExitVehicleForm.Hide()
		}, window)
	})

	ButtonContainer := fyne.NewContainerWithLayout(layout.NewVBoxLayout(), AddVehicleButton, ExitVehicleButton)

	// Getting data

	var spots []models.Spot
	utils.Db.Find(&spots)

	// Datos de ejemplo
	var data [][]string

	for _, u := range spots {
		var text string
		if u.InUse {
			text = "Ocupado"
		} else {
			text = "Libre"
		}
		data = append(data, []string{fmt.Sprint(u.ID), u.Type, u.Zone, text})
	}

	table := widget.NewTable(
		func() (int, int) {
			return len(data), len(data[0])
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("")
		},
		func(cell widget.TableCellID, cellView fyne.CanvasObject) {
			cellView.(*widget.Label).SetText(data[cell.Row][cell.Col])
		},
	)

	table.SetColumnWidth(0, 50)
	table.SetColumnWidth(1, 90)
	table.SetColumnWidth(2, 50)
	table.SetColumnWidth(3, 50)

	RightContainer := fyne.NewContainerWithLayout(layout.NewGridWrapLayout(fyne.NewSize(1180, 690)), table)

	mainContainer := fyne.NewContainerWithLayout(layout.NewHBoxLayout(),
		RightContainer,
	)

	HorizontalContainer := fyne.NewContainerWithLayout(layout.NewHBoxLayout(), ButtonContainer, mainContainer)

	window.SetContent(HorizontalContainer)
	window.ShowAndRun()
}

func AddVehicleDialog() (fyneForm *widget.Form, vehiclePlateEntry *widget.Entry, vehicleTypeEntry *widget.RadioGroup) {
	// Add vehicle
	VehiclePlateEntry := widget.NewEntry()
	VehicleTypeEntry := widget.NewRadioGroup([]string{"NORMAL", "VIP", "DISCAPACITADO", "EMERGENCIA", "PROVEEDOR"}, nil)

	AddVehicleFormItems := []*widget.FormItem{
		widget.NewFormItem("Placa", VehiclePlateEntry),
		widget.NewFormItem("Tipo de vehiculo", VehicleTypeEntry),
	}

	AddVehicleForm := &widget.Form{
		Items: AddVehicleFormItems,
	}
	return AddVehicleForm, VehiclePlateEntry, VehicleTypeEntry
}

func ExitVehicleDialog() (fyneForm *widget.Form, vehiclePlateExit *widget.Entry) {
	// Exit vehicle
	VehiclePlateExit := widget.NewEntry()

	ExitVehicleFormItems := []*widget.FormItem{
		widget.NewFormItem("Placa", VehiclePlateExit),
	}

	ExitVehicleForm := &widget.Form{
		Items: ExitVehicleFormItems,
	}
	return ExitVehicleForm, VehiclePlateExit
}
